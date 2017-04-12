package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"encoding/json"
	"io/ioutil"
)

var db *gorm.DB
var err error

//Holiday описываем структура праздника
type Holiday struct {
	Date time.Time `json:"date"`
	Name string    `json:"name"`
	Description string `json:"description"`
}

var holidays []Holiday

func main() {
	db, err = gorm.Open("sqlite3", "./holidays.db")

	if err != nil {
		fmt.Println("DB ERROR", err)
	}
	defer db.Close()

	bytes, err := ioutil.ReadFile("holidays.json")
	if err != nil {
		fmt.Println(err)
	}

	if json.Unmarshal(bytes, &holidays) != nil {
		fmt.Println(err)
	}

	r := gin.Default()
	r.GET("/", GetHolidays)
	r.GET("/holidays/", GetHolidays)
	r.GET("/holidays/:date", GetHoliday)

	r.Run(":8080")
}

func GetHoliday(c *gin.Context) {
	date := c.Params.ByName("date")
	t, _ := time.Parse("2006-01-02", date)
	for _, h := range holidays {
		if h.Date == t{
			c.JSON(200, h)
		}
	}
	c.AbortWithStatus(404)
	fmt.Println(err)
}

func GetHolidays(c *gin.Context) {
	c.JSON(200, holidays)
}
