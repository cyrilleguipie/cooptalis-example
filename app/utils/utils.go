package utils

type UserRole string

const (
	ADMIN     UserRole = "admin"
	MEMBER    UserRole = "member"
	ANONYMOUS UserRole = "anonymous"
)

type MenuType string

const (
	LIST    MenuType = "list"
	DETAILS MenuType = "details"
)

type CollaboratorStatus string

const (
	IMMIGRATION CollaboratorStatus = "immigration"
	RELOCATION  CollaboratorStatus = "relocation"
)

type JsonResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
