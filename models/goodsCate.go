package models

// 操作商品分类数据库

type GoodsCate struct {
	Id             int
	Title          string
	CateImg        string
	Link           string
	Template       string // 加载的模板
	Pid            int    // 上级分类的 id
	SubTitle       string
	Keywords       string
	Description    string
	Sort           int
	Status         int
	AddTime        int
	GoodsCateItems []GoodsCate `gorm:"foreignKey:pid;references:Id"` // 与 pid 自关联
}

func (GoodsCate) TableName() string {
	return "goods_cate"
}
