package handlers

import (
	"ecomm/web-service-gin/interfaces"
	"ecomm/web-service-gin/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

type ProductHandler struct {
	IProduct interfaces.IProduct
	// m        *messaging.ProduceMessage
	//IMessage messaging.IMessage
}

// getProducts responds with the list of all products as JSON.
func (ph *ProductHandler) GetAllProducts() func(c *gin.Context) {
	return func(c *gin.Context) {
		var products []models.Product
		products, err := ph.IProduct.GetAll()
		glog.Info("products fetched :", products)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "error in fetching contact",
			})
			glog.Errorln(err)
			c.Abort()
			return
		}
		if products == nil {
			c.JSON(http.StatusNoContent, nil)
			glog.Info("No content")
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, products)
		glog.Info("Product successfully fetched:", products)
		c.Abort()
	}
}

func (ph *ProductHandler) GetProductByID() func(*gin.Context) {
	return func(c *gin.Context) {
		if ph == nil || ph.IProduct == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in the handler",
			})
			glog.Errorln("ProductHandler or IProduct is nil")
			c.Abort()
			return
		}

		productid, ok := c.Params.Get("productid")
		glog.Info("productid fetched:", productid)

		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "id parameter not found",
			})
			glog.Errorln("id parameter not found")
			c.Abort()
			return
		}

		if productid == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "error in id param",
			})
			glog.Errorln("id cannot be empty")
			c.Abort()
			return
		}
		product, err := ph.IProduct.Get(productid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "error in fetching product",
			})
			glog.Errorln(err)
			c.Abort()
			return
		}
		if product == nil {
			c.JSON(http.StatusNoContent, nil)
			glog.Info("No content")
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, *product)
		glog.Info("Product successfully fetched:", *product)
		c.Abort()
	}
}

func (ph *ProductHandler) CreateProduct() func(*gin.Context) {
	return func(c *gin.Context) {
		if ph == nil || ph.IProduct == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in the handler",
			})
			glog.Errorln("ProductHandler or IProduct is nil")
			c.Abort()
			return
		}

		buf, err := ioutil.ReadAll(c.Request.Body)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in the body",
			})
			glog.Errorln(err)
			c.Abort()
			return
		}

		product := &models.Product{}
		err = json.Unmarshal(buf, product)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in body json format",
			})
			glog.Errorln(err)
			c.Abort()
			return
		}
		err = product.Validate()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
			glog.Errorln(err)
			c.Abort()
			return
		}
		if err := ph.IProduct.IfExists(product.Name); err != nil {
			err = nil
			if err != nil {
				c.JSON(http.StatusBadRequest, models.Response{Status: "fail", Message: err.Error()})
				glog.Errorln(err)
				c.Abort()
				return
			}
		}
		product.Status = "active"
		product.LastModified = time.Now().UTC().String()
		id, err := ph.IProduct.Create(product)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error to store in database",
			})
			glog.Errorln(err)
			c.Abort()
			return
		}
		// produce
		// ph.m = &messaging.ProduceMessage{Brokers: "localhost:29092", Topic: "CONTACT_CREATION", Key: []byte(string(contact.ID)), Data: buf}
		// err = ch.m.Produce(context.Background())
		glog.Errorln(err)
		c.JSON(http.StatusCreated, gin.H{
			"status":  "success",
			"message": id,
		})
		glog.Info("Contact successfully created:", id)
		c.Abort()
	}
}

func (ph *ProductHandler) DeleteContact() func(*gin.Context) {
	return func(c *gin.Context) {
		if ph == nil || ph.IProduct == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in the handler",
			})
			glog.Errorln("ProductHandler or IProduct is nil")
			c.Abort()
			return
		}

		id, ok := c.Params.Get("productid")

		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "id parameter not found",
			})
			glog.Errorln("id parameter not found")
			c.Abort()
			return
		}

		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "id cannot be empty",
			})
			glog.Errorln("id cannot be empty")
			c.Abort()
			return
		}
		result, err := ph.IProduct.Delete(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "error in deleting Product",
			})
			glog.Errorln(err)
			c.Abort()
			return
		}
		if result.(int64) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "no record to delete with the given id:" + id,
			})
			glog.Info("no record to delete with the given id:", id)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": fmt.Sprint(result, " record deleted"),
		})
		glog.Info("Product successfully deleted:", result)
		c.Abort()
	}
}
