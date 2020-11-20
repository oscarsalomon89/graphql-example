package schema

import (
	"fmt"
	"main/internal/models"
	"main/internal/repository"
	"main/internal/types"

	"github.com/graphql-go/graphql"
)

type EmployeeHandler interface {
	Get() *graphql.Field
	GetAll() *graphql.Field
	Save() *graphql.Field
	Delete() *graphql.Field
}

type employeeHandler struct {
	employeeRepository repository.EmployeeRepository
}

func NewEmployeeHandler(employeeRepository repository.EmployeeRepository) EmployeeHandler {
	return &employeeHandler{
		employeeRepository: employeeRepository,
	}
}

func (h *employeeHandler) GetAll() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(types.EmployeeType),
		Description: "Get all employees",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			fmt.Println("obtiene empleados")
			char, err := h.employeeRepository.GetEmployees()

			if err != nil {
				return nil, err
			}
			return char, nil
		},
	}
}

func (h *employeeHandler) Get() *graphql.Field {
	return &graphql.Field{
		Type: types.EmployeeType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "id of the employee",
				Type:        graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			fmt.Println("obtiene empleados por ID")
			id, ok := p.Args["id"].(int)
			if !ok {
				return nil, fmt.Errorf("cound not find id Args")
			}

			char, err := h.employeeRepository.GetEmployee(id)
			if err != nil {
				return nil, err
			}
			return char, nil
		},
	}
}

func (h *employeeHandler) Save() *graphql.Field {
	/*employeeInputType := graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "EmployeeInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String)},
			"city": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String)},
			"idRol": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int)},
		},
	})*/

	return &graphql.Field{
		Type:        types.EmployeeType,
		Description: "create employee character",
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"city": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"idRol": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			name, _ := params.Args["name"].(string)
			city, _ := params.Args["city"].(string)
			idRol, _ := params.Args["idRol"].(int)

			emp := models.Employee{
				Name:  name,
				City:  city,
				IdRol: idRol,
			}

			err := h.employeeRepository.AddEmployee(&emp)

			if err != nil {
				// Sí hay algun error al guardar los datos se devolvera un error 500
				fmt.Println(err)
				return nil, err
			}
			return emp, nil
		},
	}
}

func (h *employeeHandler) Delete() *graphql.Field {
	return &graphql.Field{
		Type:        types.EmployeeType,
		Description: "delete employee",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "id to delete",
				Type:        graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id, _ := params.Args["id"].(int)

			emp, err := h.employeeRepository.DeleteEmployee(id)

			if err != nil {
				// Sí hay algun error al guardar los datos se devolvera un error 500
				fmt.Println(err)
				return nil, err
			}
			return emp, nil
		},
	}
}
