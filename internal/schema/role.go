package schema

import (
	"fmt"
	"main/internal/models"
	"main/internal/repository"
	"main/internal/types"

	"github.com/graphql-go/graphql"
)

type RoleHandler interface {
	GetAll() *graphql.Field
	Save() *graphql.Field
	DeleteRol() *graphql.Field
}

type roleHandler struct {
	roleRepository repository.RoleRepository
}

func NewRoleHandler(roleRepository repository.RoleRepository) RoleHandler {
	return &roleHandler{
		roleRepository: roleRepository,
	}
}

func (h *roleHandler) GetAll() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(types.RoleType),
		Description: "Get all roles",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			fmt.Println("obtiene roles")
			char, err := h.roleRepository.GetRoles()

			if err != nil {
				return nil, err
			}
			return char, nil
		},
	}
}

func (h *roleHandler) Save() *graphql.Field {
	return &graphql.Field{
		Type:        types.RoleType,
		Description: "create role",
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Description: "new name of the human",
				Type:        graphql.NewNonNull(graphql.String),
			},
			"description": &graphql.ArgumentConfig{
				Description: "new home planet of the human",
				Type:        graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			fmt.Println("alta de rol")
			name, _ := params.Args["name"].(string)
			city, _ := params.Args["description"].(string)

			rol := models.Role{
				Name:        name,
				Description: city,
			}

			err := h.roleRepository.AddRole(&rol)

			if err != nil {
				// Sí hay algun error al guardar los datos se devolvera un error 500
				fmt.Println(err)
				return nil, err
			}
			return rol, nil
		},
	}
}

func (h *roleHandler) DeleteRol() *graphql.Field {
	return &graphql.Field{
		Type:        types.RoleType,
		Description: "delete rol",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "id to delete",
				Type:        graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id, _ := params.Args["id"]
			emp, err := h.roleRepository.DeleteRole(id.(int))

			if err != nil {
				// Sí hay algun error al guardar los datos se devolvera un error 500
				fmt.Println(err)
				return nil, err
			}
			return emp, nil
		},
	}
}
