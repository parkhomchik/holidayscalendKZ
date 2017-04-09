package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

//Holiday описываем структура праздника
type Holiday struct {
	Date time.Time `json:"date"`
	Name string    `json:"name"`
	Description string `json:"description"`
}

func main() {
	db, err = gorm.Open("sqlite3", "./holidays.db")

	if err != nil {
		fmt.Println("DB ERROR", err)
	}
	defer db.Close()

	//db.AutoMigrate(&Holiday{})

	r := gin.Default()
	r.GET("/holidays/", GetHolidays)
	r.GET("/holidays/:date", GetHoliday)

	r.Run(":8080")
}

func GetHoliday(c *gin.Context) {
	date := c.Params.ByName("date")
	var holiday Holiday
	if err := db.Where("date = ?", date).First(&holiday).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, holiday)
	}
}

func GetHolidays(c *gin.Context) {
	var holiday []Holiday
	if err := db.Find(&holiday).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, holiday)
	}

}
