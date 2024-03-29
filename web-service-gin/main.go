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
	cdb := &database.CustomerDB{Client: db}
	odb := &database.OrderDB{Client: db}
	//cdb := &filedb.FileDB{}
	productHandler := &h.ProductHandler{IProduct: pdb}
	customerHandler := &h.CustomerHandler{ICustomer: cdb}
	orderHandler := &h.OrderHandler{IOrder: odb}

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Order APIs
	router.GET("/order/orders", orderHandler.GetAllOrders())
	router.GET("/order/get/:id", orderHandler.GetOrderByID())
	router.POST("/order/create", orderHandler.PlaceOrder())
	router.DELETE("/order/delete/:id", orderHandler.CancelOrder())

	// Product APIs
	router.GET("/product/products", productHandler.GetAllProducts())
	router.GET("/product/get/:id", productHandler.GetProductByID())
	router.POST("/product/create", productHandler.CreateProduct())
	router.DELETE("/product/delete/:id", productHandler.DeleteProduct())

	// Customer APIs
	router.GET("/customer/customers", customerHandler.GetAllCustomers())
	router.GET("/customer/get/:id", customerHandler.GetCustomerByID())
	router.POST("/customer/create", customerHandler.CreateCustomer())
	router.DELETE("/customer/delete/:id", customerHandler.DeleteCustomer())

	router.Run("localhost:9090")
	// router.Run(PORT)
}
