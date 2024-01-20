package model

type Role string

const (
	RoleAdmin    Role = "admin"
	RolePro      Role = "pro"
	RoleStandard Role = "standard"
)

var Roles = map[Role]int{
	RoleAdmin:    1,
	RolePro:      2,
	RoleStandard: 3,
}
