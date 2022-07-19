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

type CustomerHandler struct {
	ICustomer interfaces.ICustomer
	// m        *messaging.ProduceMessage
	//IMessage messaging.IMessage
}

// GetAllCustomers responds with the list of all customers as JSON.
func (ch *CustomerHandler) GetAllCustomers() func(c *gin.Context) {
	return func(c *gin.Context) {
		var customers []models.Customer
		customers, err := ch.ICustomer.GetAll()
		glog.Info("customers fetched :", customers)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "error in fetching customer",
			})
			glog.Errorln(err)
			c.Abort()
			return
		}
		if customers == nil {
			c.JSON(http.StatusNoContent, nil)
			glog.Info("No content")
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, customers)
		glog.Info("Product successfully fetched:", customers)
		c.Abort()
	}
}

func (ch *CustomerHandler) GetCustomerByID() func(*gin.Context) {
	return func(c *gin.Context) {
		if ch == nil || ch.ICustomer == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in the handler",
			})
			glog.Errorln("CustomerHandler or ICustomer is nil")
			c.Abort()
			return
		}

		id, ok := c.Params.Get("id")
		glog.Info("customer id fetched:", id)

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
		customer, err := ch.ICustomer.Get(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "error in fetching customer",
			})
			glog.Errorln(err)
			c.Abort()
			return
		}
		if customer == nil {
			c.JSON(http.StatusNoContent, nil)
			glog.Info("No content")
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, *customer)
		glog.Info("Customer successfully fetched:", *customer)
		c.Abort()
	}
}

func (ch *CustomerHandler) CreateCustomer() func(*gin.Context) {
	return func(c *gin.Context) {
		if ch == nil || ch.ICustomer == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in the handler",
			})
			glog.Errorln("CustomerHandler or ICustomer is nil")
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

		customer := &models.Customer{}
		err = json.Unmarshal(buf, customer)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in body json format",
			})
			glog.Errorln(err)
			c.Abort()
			return
		}
		err = customer.Validate()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
			glog.Errorln(err)
			c.Abort()
			return
		}
		if err := ch.ICustomer.IfExists(customer.Mobile); err != nil {
			err = nil
			if err != nil {
				c.JSON(http.StatusBadRequest, models.Response{Status: "fail", Message: err.Error()})
				glog.Errorln(err)
				c.Abort()
				return
			}
		}
		// customer.LastModified = time.Now().UTC().String()
		id, err := ch.ICustomer.Create(customer)
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
		glog.Info("Customer successfully created:", id)
		c.Abort()
	}
}

func (ch *CustomerHandler) DeleteCustomer() func(*gin.Context) {
	return func(c *gin.Context) {
		if ch == nil || ch.ICustomer == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in the handler",
			})
			glog.Errorln("CustomerHandler or ICustomer is nil")
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
		result, err := ch.ICustomer.Delete(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "error in deleting customer",
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
		glog.Info("Customer successfully deleted:", result)
		c.Abort()
	}
}
