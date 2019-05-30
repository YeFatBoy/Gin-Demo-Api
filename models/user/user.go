/*
@Time : 2019/5/29 10:53
@Author : SuperShuYe
@File : user.go
@Software: GoLand
*/
package user

type User struct {
	Id        int64     `gorm:"column:id" json:"id"`
	Uid       int64     `gorm:"column:uid" json:"uid"`
	UserName  string    `gorm:"column:username" json:"username"`
	Img       string    `gorm:"column:img" json:"img"`
	//DeletedAt time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func (User) TableName() string {
	return "jmf_users"
}
