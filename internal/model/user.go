package model

const (
	UserInactive uint = iota
	UserActive
)

type User struct {
	*BaseModel
	Username       string `json:"username"`
	PasswordHashed string `json:"-" db:"password_hashed"` // json中不显示password
	Nickname       string `json:"nickname"`
	Status         uint   `json:"status"`
	RoleId         int64  `json:"role_id" db:"role_id"`
	RoleName       string `json:"role_name" db:"role_name"`
}
