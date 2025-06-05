package main
import ( 
	"go-payment-api/models"
    "go-payment-api/handler"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func main() {

	db ,err := gorm.open(sqlite.open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{}, &models.Payment{})
	handlers.setDb(db)

	r := gin.Default()

	r.POST("/register", handlers.RegisterUser)
    r.POST("/pay", handlers.CreatePayment)
    r.GET("/payments", handlers.ListPayments)

	r.Run(":8080")
}