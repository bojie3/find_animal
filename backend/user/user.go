package user

import (
	"database/sql"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/bojie/animal/backend/db"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var Users []*User

type User struct {
	ID       uint   `json:"id"` 
	Username string `json:"username" binding:"required,min=5,max=20"`
	Password string `json:"password" binding:"required,min=5,max=20"`
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := new(User)
		if err := c.Bind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		password, err := hashPassword(user.Password)
		if err != nil {
			panic(err)
		}
		user.Password = password
		
		_ , err = db.DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, user.Password)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username taken"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message":"signup successful"})
	}
}

func signin() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, foundUser := new(User), new(User)
		if err := c.Bind(user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		row := db.DB.QueryRow("SELECT id,username,password from users WHERE (username = $1)", user.Username)
		
		if err := row.Scan(&foundUser.ID, &foundUser.Username, &foundUser.Password); err != nil {
			if err  == sql.ErrNoRows {
				c.JSON(http.StatusBadRequest, gin.H{"message": "password/username incorrect"})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"message":err.Error()})
			return
		}

		correct_password := checkPasswordHash(user.Password, foundUser.Password)
		
		if !correct_password {
			c.JSON(http.StatusBadRequest, gin.H{"message": "password/username incorrect"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "signin successful"})
	}
}
