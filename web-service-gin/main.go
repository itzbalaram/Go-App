package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	_ "github.com/lib/pq"

	"ecomm/web-service-gin/database"
	h "ecomm/web-service-gin/handlers"
)

var (
	DBCONNECTION string = "host=localhost user=postgres password=postgres dbname=MyDB port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	PORT         string = ":9091"
)

func usage() {
	//fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARNING|FATAL] -log_dir=[string]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	flag.Usage = usage
	// NOTE: This next line is key you have to call flag.Parse() for the command line
	// options or "flags" that are defined in the glog module to be picked up.
	flag.Parse()
}

func main() {

	if os.Getenv("PORT") != "" {
		PORT = os.Getenv("PORT")
	}

	if os.Getenv("DBCONNECTION") != "" {
		DBCONNECTION = os.Getenv("DBCONNECTION")
	}

	db, err := database.GetConnection(DBCONNECTION)

	if err != nil {
		glog.Fatal("Database Error:", err)
	}

	pdb := &database.ProductDB{Client: db}
	//cdb := &filedb.FileDB{}
	productHandler := &h.ProductHandler{IProduct: pdb}

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/product/products", productHandler.GetAllProducts())
	router.GET("/product/get/:productid", productHandler.GetProductByID())
	router.POST("/product/create", productHandler.CreateProduct())
	router.DELETE("/product/delete/:productid", productHandler.DeleteContact())

	router.Run("localhost:9090")
	// router.Run(PORT)

}
