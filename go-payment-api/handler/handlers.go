package handlers

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
    "github.com/google/uuid"
    "go-payment-api/models"
)

var db *gorm.DB

func SetDb(d *gorm.DB) {
    db = d
}

func RegisterUser(c *gin.Context) {
    var user models.User

    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := db.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}

func CreatePayment(c *gin.Context) {
    var payment models.Payment

    if err := c.BindJSON(&payment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    payment.TransactionID = uuid.New().String()
    if err := db.Create(&payment).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, payment)
}

func ListPayments(c *gin.Context) {
    var payments []models.Payment
    db.Find(&payments)
    c.JSON(http.StatusOK, payments)
}
