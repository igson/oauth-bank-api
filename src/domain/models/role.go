package models

import (
	"fmt"
	"strings"
)

type RolePermissions struct {
	rolePermissions map[string][]string
}

func (p RolePermissions) IsAuthorizedFor(role string, routeName string) bool {
	perms := p.rolePermissions[role]
	for _, r := range perms {
		fmt.Printf("Rota %s com permissão de %s \n", r, role)
		if r == strings.TrimSpace(routeName) {
			return true
		}
	}
	return false
}

func GetRolePermissions() RolePermissions {
	return RolePermissions{map[string][]string{
		"admin": {"ListarClientes", "BuscarCliente", "BuscarContaDeCliente", "RealizarTransacao"},
		"user":  {"BuscarCliente", "RealizarTransacao", "ListarClientes"},
	}}
}
