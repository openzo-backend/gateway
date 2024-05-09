// Example using Gin framework
package main

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Route definition

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
	// reverseProxy("http://localhost:8080", c)
	reverseProxy("http://localhost:8080", c)
}

// Handler function for Service 2
func handleService2(c *gin.Context) {
	// Forward request to Service 2
	reverseProxy("http://localhost:8081", c)
	// store-service
}
func handleService3(c *gin.Context) {
	// Forward request to Service 2
	reverseProxy("http://localhost:8082", c)
}
func handleService4(c *gin.Context) {
	// Forward request to Service 2
	reverseProxy("http://localhost:8083", c)
}
func handleService5(c *gin.Context) {
	// Forward request to Service 2
	reverseProxy("http://localhost:8084", c)
}
func handleService6(c *gin.Context) {
	// Forward request to Service 2
	reverseProxy("http://localhost:8086", c)
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
