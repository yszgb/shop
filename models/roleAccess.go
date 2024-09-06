package models

type RoleAccess struct {
	AccessId int
	RoleId   int
}

// 使结构体和数据库中的表对应
func (RoleAccess) TableName() string {
	return "role_access"
}
