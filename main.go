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
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")            // Add "Authorization" header to allowed headers
	config.AllowHeaders = append(config.AllowHeaders, "Content-Type")             // Add "Content-Type" header to allowed headers
	config.AllowHeaders = append(config.AllowHeaders, "Accept")                   // Add "Accept" header to allowed headers
	config.AllowHeaders = append(config.AllowHeaders, "Origin")                   // Add "Origin" header to allowed headers
	config.AllowHeaders = append(config.AllowHeaders, "Upgrade")                  // Add "Origin" header to allowed headers
	config.AllowHeaders = append(config.AllowHeaders, "Sec-WebSocket-Key")        // Add "Origin" header to allowed headers
	config.AllowHeaders = append(config.AllowHeaders, "Sec-WebSocket-Version")    // Add "Origin" header to allowed headers
	config.AllowHeaders = append(config.AllowHeaders, "Sec-WebSocket-Extensions") // Add "Origin" header to allowed headers
	config.AllowHeaders = append(config.AllowHeaders, "Sec-WebSocket-Protocol")   // Add "Origin" header to allowed headers
	config.AllowHeaders = append(config.AllowHeaders, "Sec-WebSocket-Origin")     // Add "Origin" header to allowed headers
	config.AllowHeaders = append(config.AllowHeaders, "Sec-WebSocket-Location")   // Add "Origin" header to allowed headers
	config.AllowHeaders = append(config.AllowHeaders, "Sec-WebSocket-Accept")     // Add "Origin" header to allowed headers
	config.AllowHeaders = append(config.AllowHeaders, "Sec-WebSocket-Protocol")   // Add "Origin" header to allowed headers
	config.AllowHeaders = append(config.AllowHeaders, "Connection")               // Add "Origin" header to allowed headers

	config.AllowWebSockets = true

	// config.AllowHeaders = true

	// Route definition
	router.Use(cors.New(config))
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome to the gateway",
		})
	})
	router.Any("/users/*path", handleService1)
	router.Any("/stores/*path", handleService2)
	router.Any("/products/*path", handleService3)
	router.Any("/sales/*path", handleService4)
	router.Any("/online_orders/*path", handleService5)
	router.Any("/enquiry/*path", handleService6)
	router.Any("/ads/*path", handleService7)
	router.Any("/search/*path", handleService8)
	router.Any("/notifications/*path", handleService9)
	router.Any("/store-customer/*path", handleService10)

	router.Any("/ws/*path", handleWebSocket)
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
	reverseProxy("http://online-order-service:8080", c)
}
func handleService6(c *gin.Context) {
	// Forward request to Service 2
	// reverseProxy("http://localhost:8086", c)
	reverseProxy("http://enquiry-service:8080", c)
}

func handleService7(c *gin.Context) {
	// Forward request to Service 2
	// reverseProxy("http://localhost:8086", c)
	reverseProxy("http://ad-service:8080", c)
}

func handleService8(c *gin.Context) {
	// Forward request to Service 2
	// reverseProxy("http://localhost:8086", c)
	reverseProxy("http://search-service:8080", c)
}

func handleService9(c *gin.Context) {
	// Forward request to Service 2
	// reverseProxy("http://localhost:8086", c)
	reverseProxy("http://notification-service:8080", c)
}
func handleService10(c *gin.Context) {
	// Forward request to Service 2
	// reverseProxy("http://localhost:8086", c)
	reverseProxy("http://store-customer-service:8080", c)
}

func handleWebSocket(c *gin.Context) {
	// Forward request to Service 2
	// reverseProxy("http://localhost:8085", c)
	reverseProxy("http://websocket-service:8080", c)
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
