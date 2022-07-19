package handlers

import (
	"ecomm/web-service-gin/interfaces"
	"ecomm/web-service-gin/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

type OrderHandler struct {
	IOrder interfaces.IOrder
	// IProduct interfaces.IProduct
	// m        *messaging.ProduceMessage
	//IMessage messaging.IMessage
}

// GetAllCustomers responds with the list of all customers as JSON.
func (oh *OrderHandler) GetAllOrders() func(c *gin.Context) {
	return func(c *gin.Context) {
		var orders []models.Order
		orders, err := oh.IOrder.GetAll()
		glog.Info("orders fetched :", orders)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "error in fetching orders",
			})
			glog.Errorln(err)
			c.Abort()
			return
		}
		if orders == nil {
			c.JSON(http.StatusNoContent, nil)
			glog.Info("No content")
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, orders)
		glog.Info("Orders successfully fetched:", orders)
		c.Abort()
	}
}

func (oh *OrderHandler) GetOrderByID() func(*gin.Context) {
	return func(c *gin.Context) {
		if oh == nil || oh.IOrder == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in the handler",
			})
			glog.Errorln("OrderHandler or IOrder is nil")
			c.Abort()
			return
		}

		id, ok := c.Params.Get("id")
		glog.Info("order id fetched:", id)

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
				"message": "error in id param",
			})
			glog.Errorln("id cannot be empty")
			c.Abort()
			return
		}
		order, err := oh.IOrder.Get(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "error in fetching customer",
			})
			glog.Errorln(err)
			c.Abort()
			return
		}
		if order == nil {
			c.JSON(http.StatusNoContent, nil)
			glog.Info("No content")
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, &order)
		glog.Info("Order successfully fetched:", &order)
		c.Abort()
	}
}

func (oh *OrderHandler) PlaceOrder() func(*gin.Context) {
	return func(c *gin.Context) {
		if oh == nil || oh.IOrder == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in the handler",
			})
			glog.Errorln("OrderHandler or IOrder is nil")
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

		order := &models.Order{}
		err = json.Unmarshal(buf, order)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in body json format",
			})
			glog.Errorln(err)
			c.Abort()
			return
		}
		err = order.Validate()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
			glog.Errorln(err)
			c.Abort()
			return
		}
		// if err := oh.IOrder.IfExists(order.); err != nil {
		// 	err = nil
		// 	if err != nil {
		// 		c.JSON(http.StatusBadRequest, models.Response{Status: "fail", Message: err.Error()})
		// 		glog.Errorln(err)
		// 		c.Abort()
		// 		return
		// 	}
		// }
		// customer.LastModified = time.Now().UTC().String()
		// ProductDetails

		// id, ok := c.Params.Get("prodid")
		// glog.Info("prod id fetched:", id)
		// if !ok {
		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"status":  "fail",
		// 		"message": "id parameter not found",
		// 	})
		// 	glog.Errorln("id parameter not found")
		// 	c.Abort()
		// 	return
		// }

		// if id == "" {
		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"status":  "fail",
		// 		"message": "error in id param",
		// 	})
		// 	glog.Errorln("id cannot be empty")
		// 	c.Abort()
		// 	return
		// }
		// product, err := oh.IProduct.Get(id)
		// glog.Info("prod details fetched:", product)

		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"status":  "fail",
		// 		"message": "error in fetching product",
		// 	})
		// 	glog.Errorln(err)
		// 	c.Abort()
		// 	return
		// }
		// if product == nil {
		// 	c.JSON(http.StatusNoContent, nil)
		// 	glog.Info("No content")
		// 	c.Abort()
		// 	return
		// }
		// c.JSON(http.StatusOK, *product)
		// glog.Info("Product successfully fetched:", *product)
		// // order.Amount = append(order.Amount, product.price)

		id, err := oh.IOrder.Create(order)
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
		glog.Info("Order successfully Placed:", id)
		c.Abort()
	}
}

func (oh *OrderHandler) CancelOrder() func(*gin.Context) {
	return func(c *gin.Context) {
		if oh == nil || oh.IOrder == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in the handler",
			})
			glog.Errorln("OrderHandler or IOrder is nil")
			c.Abort()
			return
		}

		id, ok := c.Params.Get("id")

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
		result, err := oh.IOrder.Delete(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "error in cancelling order",
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
		glog.Info("Order successfully cancelled:", result)
		c.Abort()
	}
}
