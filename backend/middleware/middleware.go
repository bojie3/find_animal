package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Authenticate(c *gin.Context) {
	//get the cookie
	tokenString, _ := c.Cookie("Auth")

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Unauthorized, cookie not found"})
		c.Abort()
		return
	} 

	// Parse takes the token string and a function for looking up the key. 
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error in parsing the token"})
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//check expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(http.StatusUnauthorized, gin.H{"error":"token expired, please login again"})
			c.Abort()
			return
		}

		//attach to req
		c.Set("username", claims["username"])
		c.Set("id", claims["sub"])

		//continue
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"invalid token"})
		c.Abort()
		return
	}
}


func CORSmiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Auth")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT,PATCH")
	
			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(204)
				return
			}
	
			c.Next()
		}
}