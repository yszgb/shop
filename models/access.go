package models

type Access struct {
	Id          int
	ModuleName  string // 模块名称
	ActionName  string // 操作名称
	Type        int    // 节点类型 :  1、表示模块    2、表示菜单     3、操作
	Url         string // 路由跳转地址
	ModuleId    int    // 用于自关联，为 0 表示一级模块，一级模块有多个二级模块
	Sort        int
	Description string
	Status      int
	AddTime     int
	AccessItem  []Access `gorm:"foreignKey:ModuleId;references:Id"` // Access 类型的切片，自己和自己的 ModuleId 一对多关联
	Checked     bool     `gorm:"-"`                                 // 忽略本字段
}

func (Access) TableName() string {
	return "access"
}
