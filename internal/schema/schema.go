package schema

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

type SchemaBuilder interface {
	Build() (graphql.Schema, error)
}

type schemaBuilder struct {
	employeeHandler EmployeeHandler
	roleHandler     RoleHandler
}

func NewSchemaBuilder(employeeHandler EmployeeHandler,
	roleHandler RoleHandler) SchemaBuilder {
	return &schemaBuilder{
		employeeHandler: employeeHandler,
		roleHandler:     roleHandler}
}

func (hdr *schemaBuilder) Build() (graphql.Schema, error) {
	fmt.Println("Inicializando esquemas....")

	fieldsQueries := graphql.Fields{}
	fieldsMutations := graphql.Fields{}

	fieldsQueries["employees"] = hdr.employeeHandler.GetAll()
	fieldsQueries["employee"] = hdr.employeeHandler.Get()
	fieldsQueries["roles"] = hdr.roleHandler.GetAll()

	fieldsMutations["addEmployee"] = hdr.employeeHandler.Save()
	fieldsMutations["addRole"] = hdr.roleHandler.Save()
	fieldsMutations["deleteEmployee"] = hdr.employeeHandler.Delete()
	fieldsMutations["deleteRole"] = hdr.roleHandler.DeleteRol()

	return graphql.NewSchema(
		graphql.SchemaConfig{
			Query: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   "Query",
					Fields: fieldsQueries,
				},
			),
			Mutation: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   "Mutation",
					Fields: fieldsMutations,
				},
			),
		},
	)
}
