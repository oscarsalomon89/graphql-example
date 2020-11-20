package types

import (
	"fmt"
	"main/internal/models"
	"main/internal/repository"

	"github.com/graphql-go/graphql"
)

var EmployeeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Employee",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
		},
		"idRol": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"city": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"date_created": &graphql.Field{
			Type: graphql.DateTime,
		},
		"date_updated": &graphql.Field{
			Type: graphql.DateTime,
		},
		"rol": &graphql.Field{
			Type: graphql.NewNonNull(RoleType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				db, err := repository.GetDB()
				if err != nil {
					fmt.Println("error connet DB: " + err.Error())
				}

				employee := params.Source.(models.Employee)
				roleRepository := repository.NewRoleRepository(db)

				rol, err := roleRepository.GetRole(employee.IdRol)

				if err != nil {
					return nil, err
				}
				return rol, nil
			},
		},
	},
})

var RoleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Role",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
		},
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"description": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		/*"employees": &graphql.Field{
			Type: graphql.NewList(EmployeeType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				db, err := repository.GetDB()
				if err != nil {
					fmt.Println("error connet DB: " + err.Error())
				}

				employeeRepository := repository.NewEmployeeRepository(db)
				rol := params.Source.(models.Role)

				employees, err := employeeRepository.GetEmployeesByRol(rol.ID)

				if err != nil {
					return nil, err
				}
				return employees, nil
			},
		},*/
	},
})
