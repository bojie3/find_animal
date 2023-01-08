package user

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/bojie/animal/backend/db"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	_ "github.com/lib/pq"
)

var Users []*User

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username" binding:"required,min=5,max=20"`
	Password string `json:"password" binding:"required,min=5,max=20"`
}

type Tokens struct {
	Auth    string `json:auth`
	Refresh string `json:refresh`
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func register() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := new(User)
		if err := c.Bind(&user); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "invalid input, unable to parse",
				"error output" : err.Error(),
			})
			return
		}
		password, err := hashPassword(user.Password)
		if err != nil {
			panic(err)
		}
		user.Password = password

		_, err = db.DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, user.Password)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Username taken",
				"error output" : err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "registration successful"})
	}
}

func login() gin.HandlerFunc {
	return func(c *gin.Context) {
		//get the input and compare with values in database
		user, foundUser := new(User), new(User)
		if err := c.Bind(user); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "invalid input, unable to parse",
				"error output" : err.Error(),
			})
			return
		}

		row := db.DB.QueryRow("SELECT id,username,password from users WHERE (username = $1)", user.Username)

		if err := row.Scan(&foundUser.ID, &foundUser.Username, &foundUser.Password); err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "password/username incorrect",
					"error output": err.Error(),
				})
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "error occured",
				"error output": err.Error(),
			})
			return
		}

		correct_password := checkPasswordHash(user.Password, foundUser.Password)

		if !correct_password {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "password/username incorrect"})
			return
		}

		//generate jwt token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub":   user.ID,
			"username": user.Username,
			"exp":  time.Now().Local().Add(time.Minute * 15).Unix(),
		})

		tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "failed to create jwt token",
				"error output": err.Error(),
			})
			panic(err)
		}	

		refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub":   user.ID,
			"exp":  time.Now().Local().Add(time.Hour * 12).Unix(),
		})

		refreshString, err := refresh.SignedString([]byte(os.Getenv("SECRET_KEY")))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "failed to create jwt refresh token",
				"error output": err.Error(),
			})
			panic(err)
		}

		// c.SetSameSite(http.SameSiteLaxMode)
		// c.SetCookie("Auth", tokenString, 3600 * 24 * 30, "/", "localhost:3000", false, true)

		c.JSON(http.StatusOK, gin.H{
			"message": "login successful",
			"auth":   tokenString,
			"refresh": refreshString,
		})
	}
}

func refresh() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokens := new(Tokens)
		if err := c.Bind(tokens); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "invalid tokens, unable to parse",
				"error output" : err.Error(),
			})
		}

		refresh := tokens.Refresh
		auth := tokens.Auth
		newAuth := false
		newRefresh := false
		username := ""
		id := 0.0

		// checks if the auth token is valid and if new one is needed
		// Parse takes the token string and a function for looking up the key. 
		authToken, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error with auth token",
				"error output" : err.Error(),
			})
			c.Abort()
			return
		}

		if claims, ok := authToken.Claims.(jwt.MapClaims); ok && authToken.Valid {
			//check expiration
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				c.JSON(http.StatusUnauthorized, gin.H{"error":"token expired, please login again"})
				c.Abort()
				return
			}
			
			//if going to expire soon, then produce new auth token
			if float64(time.Now().Add(time.Minute * 5).Unix()) > claims["exp"].(float64) {
				newAuth = true
			}

			username = claims["username"].(string)
			id = claims["sub"].(float64)
			
	
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error":"invalid token"})
			c.Abort()
			return
		}

		// checks if refresh token is valid and if new one is needed
		// Parse takes the token string and a function for looking up the key. 
		refreshToken, err := jwt.Parse(refresh, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error with auth token",
				"error output" : err.Error(),
			})
			c.Abort()
			return
		}

		if claims, ok := refreshToken.Claims.(jwt.MapClaims); ok && refreshToken.Valid {
			//check expiration
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				c.JSON(http.StatusUnauthorized, gin.H{"error":"token expired, please login again"})
				c.Abort()
				return
			}
			
			//if going to expire soon, then produce new auth token
			if float64(time.Now().Add(time.Hour * 1).Unix()) > claims["exp"].(float64) {
				newRefresh = true
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error":"invalid token"})
			c.Abort()
			return
		}

		resp := new(Tokens)
		if (newAuth) {
			//generate auth token
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"sub":   id,
				"username": username,
				"exp":  time.Now().Local().Add(time.Minute * 15).Unix(),
			})

			tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "failed to create jwt token",
					"error output": err.Error(),
				})
				c.Abort()
				return
			}	
			resp.Auth = tokenString
		}

		if (newRefresh) {
			//generate refresh token
			id, _ := c.Get("sub")
			refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"sub":   id,
				"exp":  time.Now().Local().Add(time.Hour * 12).Unix(),
			})
	
			refreshString, err := refresh.SignedString([]byte(os.Getenv("SECRET_KEY")))
	
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "failed to create jwt refresh token",
					"error output": err.Error(),
				})
				c.Abort()
				return
			}

			resp.Refresh = refreshString
		}

		c.JSON(http.StatusOK, gin.H{
			"Auth": resp.Auth,
			"Refresh": resp.Refresh,
		})
	}
}


func validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Get("sub")
		username, _ := c.Get("username")

		c.JSON(http.StatusOK, gin.H{
			"message":"I am logged in",
			"id":id,
			"username":username,
		})
	}
}