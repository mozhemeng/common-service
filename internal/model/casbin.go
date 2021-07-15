package model

type CasbinPolicy struct {
	PType    string `json:"-" db:"p_type"`
	RoleName string `json:"role_name" db:"v0"`
	Path     string `json:"path" db:"v1"`
	Method   string `json:"method" db:"v2"`
	V3       string `json:"-" db:"v3"`
	V4       string `json:"-" db:"v4"`
	V5       string `json:"-" db:"v5"`
}
