package main

import (
	_ "embed"
	"hostmonitorsystem/sqlite"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ilyakaznacheev/cleanenv"
	_ "modernc.org/sqlite"
)

var config Config
var gin_r *gin.Engine

type HostRow struct {
	Hostname string `json:"hostname"`
	IP       string `json:"ip"`
	CPUMax   string `json:"cpu_max"`
	MemMax   string `json:"mem_max"`
	Dept     string `json:"dept"`
	App      string `json:"app"`
	Count    string `json:"count"`
}

// QueryResult 是 /query 的响应体
type QueryResult struct {
	OK    bool      `json:"ok"`
	Error string    `json:"error,omitempty"`
	Total int       `json:"total"`
	Rows  []HostRow `json:"rows"`
}

func main() {
	initConfig()
	err := sqlite.InitDB(config.host_db_name)
	if err != nil {
		println("初始化数据库失败")
		panic(err)
	}
	defer sqlite.CloseDB()
	initGin()

	log.Printf("服务已启动: http://%s", config.host+":"+config.port+"/")
	if err := gin_r.Run(config.host + ":" + config.port); err != nil {
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
