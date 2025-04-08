package main

import (
	"log"

	"github.com/ajz007/learngo-fakestoreapi/internal/handler"
	"github.com/gin-gonic/gin"
)

/*
*

	gin.Context actually represents the context of HTTP request.
	Think of *gin.Context as a powerful wrapper around:

	- The HTTP request (http.Request)
	- The HTTP response (http.ResponseWriter)
	- Params, query strings, headers
	- Middleware state, errors, etc.

*
*/
func main() {
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		//“Set the HTTP response code to 200, and write this JSON object as the response body.”
		// gin.H is --> type H map[string]interface{} which basically means a map with key as string and value as object of any type
		//internally this translates to following:
		/*
			- c.Writer.WriteHeader(200) // sets HTTP status
			- jsonBytes, _ := json.Marshal(data) // marshals the map to JSON
			- c.Writer.Write(jsonBytes) // writes to the HTTP response stream

			--> c.Writer is a wrapper over http.ResponseWriter, which is how Go handles writing HTTP responses.
		*/

		//
		c.JSON(200, gin.H{
			"status": "up!",
		})
	})

	router.GET("/products", handler.NewProductHandler().GetProducts)

	log.Println("Starting the server on : 8080...")

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start the server on 8080 %v", err)
	}
	router.Run()

}
