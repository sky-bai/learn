package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"learn/112_gorm/dal"
)

var DB *gorm.DB

func init() {
	userName := "root"
	password := "12345678"
	host := "127.0.0.1"
	port := "3306"
	Dbname := "gorm"
	var mylogger logger.Interface
	mylogger = logger.Default.LogMode(logger.Info)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", userName, password, host, port, Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mylogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connect success", db)
	DB = db
	dal.SetDefault(DB)
}

type Student struct {
	Id   int
	Name string
	Age  int
}

func main() {
	//DB.AutoMigrate(&Student{})

	g := gen.NewGenerator(gen.Config{
		OutPath:      "./112_gorm/dal",
		ModelPkgPath: "./112_gorm/model",
		Mode:         gen.WithDefaultQuery,
	})

	g.UseDB(DB)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()

	dal.Student.WithContext().Where(dal.Student.ID.Eq(1))

}
