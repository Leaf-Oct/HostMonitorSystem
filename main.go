package main

import (
	"hostmonitorsystem/sqlite"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ilyakaznacheev/cleanenv"
	_ "modernc.org/sqlite"
)

var config Config
var gin_r *gin.Engine

func main() {
	initConfig()
	err := sqlite.InitDB(config.Host_db_name)
	if err != nil {
		println("初始化数据库失败")
		panic(err)
	}
	defer sqlite.CloseDB()
	initGin()

	log.Println("服务已启动: http://" + config.Host + ":" + config.Port + "/")
	if err := gin_r.Run(config.Host + ":" + config.Port); err != nil {
		println("服务启动失败")
		panic(err.Error())
	}
}

func initConfig() {
	if err := cleanenv.ReadConfig("config.yaml", &config); err != nil {
		panic("加载配置文件失败 " + err.Error())
	}
}
func initGin() {
	gin.SetMode(gin.DebugMode)
	gin_r = gin.New()
	gin_r.Use(gin.Logger(), gin.Recovery())
	gin_r.LoadHTMLGlob("assets/*")
	gin_r.GET("/", indexHandler)
	gin_r.GET("/exception", exceptionHandler)
	gin_r.GET("/about", aboutHandler)
	gin_r.GET("/query/:table", queryHandler)

}
