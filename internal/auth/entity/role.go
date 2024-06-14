package entity

type Role struct {
	ID          string
	Name        string
	Permissions []string
}

func NewRole(id, name string, permissions []string) *Role {
	return &Role{
		ID:          id,
		Name:        name,
		Permissions: permissions,
	}
}
