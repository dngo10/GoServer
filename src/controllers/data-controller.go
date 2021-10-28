package controllers

import (
	data "goapp2/src/models"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DataController struct{
	Db *gorm.DB
}

func (d *DataController) Init(){
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
		  SlowThreshold:              time.Second,   // Slow SQL threshold
		  LogLevel:                   logger.Silent, // Log level
		  IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
		  Colorful:                  false,          // Disable color
		},
	  )

	os := runtime.GOOS
	var path string

	if os == "linux"{
		path = "/home/GoServer/data.db"
	}else if os == "windows"{
		path = "data.db"
	}

	db , err := gorm.Open(sqlite.Open(path), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil{
		panic("failed to connect database")
	}else{
		d.Db = db
		d.Db.AutoMigrate(&data.Data{})
	}
}

func (d *DataController) GetData(c *gin.Context){
	uuId := c.Param("uuid")
	var dataGotten data.Data
	result := d.Db.Where(&data.Data{Uuid: uuId}).First(&dataGotten)

	if result.Error != nil{
		c.IndentedJSON(http.StatusNotFound, data.SendErr("Couldn't find data"))
		return
	}
	c.IndentedJSON(http.StatusCreated, dataGotten)
}

func (d *DataController) AddData(c *gin.Context){
	var newData data.Data

	if err := c.BindJSON(&newData); err != nil{
		c.IndentedJSON(http.StatusBadGateway, data.SendErr("couldn't bind new data"))
		return
	}
	result := d.Db.Create(&newData)
	if result.Error != nil{
		c.IndentedJSON(http.StatusExpectationFailed, data.SendErr("couldn't add new row"))
		return
	}
	c.IndentedJSON(http.StatusOK, data.SendErr("success"))
}

func (d *DataController) EditData(c *gin.Context){
	var editedData data.Data
 
	if err := c.BindJSON(&editedData); err != nil{
		c.IndentedJSON(http.StatusBadRequest, data.SendErr("cound't bind new data"))
		return
	}

	result := d.Db.Model(&data.Data{}).Where(&data.Data{Uuid: editedData.Uuid}).Updates(data.Data{Data: editedData.Data})
	if result.Error != nil{
		c.IndentedJSON(http.StatusBadRequest, data.SendErr("couldn't update new data"))
	}
	c.IndentedJSON(http.StatusOK, data.SendErr("success"))
}

func (d *DataController) DeleteData(c *gin.Context){
	uuId := c.Param("uuid")

	result := d.Db.Where(&data.Data{Uuid: uuId}).Delete(&data.Data{})

	if result.Error != nil{
		c.IndentedJSON(http.StatusConflict, data.SendErr("couldn't delete data"))
		return
	}
	c.IndentedJSON(http.StatusOK, data.SendErr("success"))
}