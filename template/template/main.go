package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/swag/example/celler/controller"
	"golang-microservices-gRPC/src/services/user/controller"
	"golang-microservices-gRPC/src/services/user/docs"
	"log"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/
func main() {
	// programatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:9080"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.New()
	var v2Grouping = r.Group("v2")
	v2Grouping.GET("/testapi/get-string-by-int/:some_id", controller.GetStringByInt)
	v2Grouping.GET("//testapi/get-struct-array-by-string/:some_id", controller.GetStructArrayByString)
	v2Grouping.POST("/testapi/upload", controller.Upload)
	// use ginSwagger middleware to serve the API docs
	v2Grouping.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run("localhost:9080")
	if err != nil {
		log.Print("Can't start server at port :%s", "9080")
	}
}

//...
