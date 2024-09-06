package models

// 操作数据库-表 gin-user

// 与 user 表一一映射，大写供外部操作
type User struct { // 默认表名是 `users`
	Id       int
	Username string
	Age      int
	Email    string
	AddTime  int // add_time 改为驼峰式 AddTime
}

// 配置数据库表名， User 默认操作的是 `users` ，改为操作 `user`
func (User) TableName() string {
	return "user"
}
