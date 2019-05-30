/*
@Time : 2019/5/30 15:49
@Author : SuperShuYe
@File : category.go
@Software: GoLand
*/
package category

type Category struct {
	Id    int     `gorm:"column:id" json:"id"`
	Name  string  `gorm:"column:name" json:"name"`
	Topic []Topic `gorm:"ForeignKey:CategoryTopic" json:"topic"`
}

type Topic struct {
	Id            int    `gorm:"column:id" json:"id"`
	CategoryTopic int    `gorm:"column:topic_category_id" json:"topic_category_id"`
	Title         string `gorm:"column:title" json:"title"`
	LogoUrl       string `gorm:"column:logo_url" json:"logo_url"`
}

func (Category) TableName() string {
	return "jmf_topic_category"
}

func (Topic) TableName() string {
	return "jmf_topic"
}