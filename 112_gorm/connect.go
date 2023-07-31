package main

type User struct {
	Id       uint   `gorm:"primarykey"`
	Name     string `gorm:"type:varchar(20);not null"`
	Articles []Article
}
type Article struct {
	ID        uint   `gorm:"primarykey"`
	Title     string `gorm:"type:varchar(20);not null"`
	UserRefer uint
	User      User `gorm:"foreignKey:UserRefer"` // 外键关联
}

func main() {
	DB.AutoMigrate(&User{}, &Article{})
}
