// Example using Gin framework
package main

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = append(config.AllowHeaders, "Authorization") // Add "Authorization" header to allowed headers

	// Route definition
	router.Use(cors.New(config))
	router.Any("/users/*path", handleService1)
	router.Any("/stores/*path", handleService2)
	router.Any("/products/*path", handleService3)
	router.Any("/sales/*path", handleService4)
	router.Any("/online_orders/*path", handleService5)
	router.Any("/enquiry/*path", handleService6)

	// Add more routes for other microservices

	router.Run(":8000") // Run server on port 8080
}

// Handler function for Service 1
func handleService1(c *gin.Context) {
	// Forward request to Service 1
	reverseProxy("http://user-service:8080", c)
	// reverseProxy("http://localhost:8080", c)
}

// Handler function for Service 2
func handleService2(c *gin.Context) {
	// Forward request to Service 2
	reverseProxy("http://store-service:8080", c)
	// store-service
}
func handleService3(c *gin.Context) {
	// Forward request to Service 2
	// reverseProxy("http://localhost:8082", c)
	reverseProxy("http://product-service:8080", c)
}
func handleService4(c *gin.Context) {
	// Forward request to Service 2
	// reverseProxy("http://localhost:8083", c)
	reverseProxy("http://sale-service:8080", c)
}
func handleService5(c *gin.Context) {
	// Forward request to Service 2
	// reverseProxy("http://localhost:8084", c)
	reverseProxy("http://onlineorder-service:8080", c)
}
func handleService6(c *gin.Context) {
	// Forward request to Service 2
	// reverseProxy("http://localhost:8086", c)
	reverseProxy("http://enquiry-service:8080", c)
}

// Reverse proxy function
func reverseProxy(target string, c *gin.Context) {
	// Parse target URL
	url, _ := url.Parse(target)

	// Create reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update request URL to match target URL
	c.Request.URL.Path = c.Param("path")

	// Serve HTTP response
	proxy.ServeHTTP(c.Writer, c.Request)
}
