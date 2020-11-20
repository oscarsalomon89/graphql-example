package main

import (
	"fmt"
	"main/internal/controllers"
	"main/internal/repository"
	"main/internal/schema"
	"main/internal/services"
	"net/http"

	"github.com/friendsofgo/graphiql"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	/*employeeController := newEmployeeController()
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc

	myRouter.HandleFunc("/employee", employeeController.Get).Methods("GET")
	myRouter.HandleFunc("/employees", employeeController.GetAll).Methods("GET")
	myRouter.HandleFunc("/employee", employeeController.Save).Methods("POST")

	fmt.Println("server is started at: http://localhost:8080/")
	http.ListenAndServe(":8080", myRouter)*/

	schema := newGraphQLSchema()

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	// static file server to serve Graphiql in-browser editor
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		panic(err)
	}

	http.Handle("/graphiql", graphiqlHandler)
	// graphql api  server
	http.Handle("/graphql", h)

	fmt.Println("graphql api server is started at: http://localhost:8080/graphql")
	http.ListenAndServe(":8080", nil)
}

func newEmployeeController() controllers.EmployeeController {
	db, err := repository.GetDB()
	if err != nil {
		fmt.Println("Errorrrr")
	}

	employeeRepository := repository.NewEmployeeRepository(db)
	employeeService := services.NewEmployeeService(employeeRepository)
	employeeController := controllers.EmployeeController{
		EmployeService: employeeService,
	}

	return employeeController
}

func newGraphQLSchema() graphql.Schema {
	db, err := repository.GetDB()
	if err != nil {
		fmt.Println("error connet DB: " + err.Error())
	}

	employeeRepository := repository.NewEmployeeRepository(db)
	employeeHandler := schema.NewEmployeeHandler(employeeRepository)

	roleRepository := repository.NewRoleRepository(db)
	roleHandler := schema.NewRoleHandler(roleRepository)

	schemaCreate := schema.NewSchemaBuilder(employeeHandler, roleHandler)

	schema, err := schemaCreate.Build()

	if err != nil {
		fmt.Println("error creating graphQL Schema: " + err.Error())
	}

	return schema
}
