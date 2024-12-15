package main

import (
	// "database/sql"
	"fmt"
	"log"
	"net/http"

	// "strconv"
	// "time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	DRIVER    = "mysql"
	MYSQLPORT = "3306"
	HOST      = "localhost"
	USER      = "root"
	PASSWORD  = "password"
	DBNAME    = "project"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func main() {
	ConnectDatabase()

	router := gin.Default()

	// Apply custom CORS middleware
	router.Use(CORSMiddleware())

	router.POST("/user/create/", CreateUserHandler)
	router.Run(":3000")
}

var db *sqlx.DB
var err error

func ConnectDatabase() {
	// connectionString := "USER:PASSWORD@tcp(HOST:MYSQLPORT)/DBNAME"
	connectionStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", USER, PASSWORD, HOST, MYSQLPORT, DBNAME)
	db, err = sqlx.Open(DRIVER, connectionStr)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func CreateUserHandler(c *gin.Context) {

	type createUserRequest struct {
		Name    string `json:"name" binding:"required"`
		Email   string `json:"email" binding:"required"`
		Phone   string `json:"phone" binding:"required"`
		Message string `json:"message" binding:"required"`
	}

	var request createUserRequest
	err = c.ShouldBindJSON(&request)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"FAILED": err})
		return
	}

	query := "INSERT INTO users (name, email, phone, message) VALUES (?, ?, ?, ?)"
	result, err := db.Exec(query, request.Name, request.Email, request.Phone, request.Message)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"Query Exection Error": err})
		return
	}

	userID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"Last Insert ID getting Error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"SUCCESS": userID})
}

// func GetUserDetails(c *gin.Context) {

// 	userIDstr := c.Param("id")
// 	userID, err := strconv.Atoi(userIDstr)
// 	if err != nil {
// 		log.Fatal(err)
// 		c.JSON(http.StatusBadRequest, gin.H{"FAILED": err})
// 	}

// 	type userModel struct {
// 		ID        int          `db:"id"`
// 		Name      string       `db:"name"`
// 		Age       int          `db:"age"`
// 		Email     string       `db:"email"`
// 		CreatedAt sql.NullTime `db:"created_at"`
// 	}

// 	tempModel := userModel{}

// 	query := "SELECT * FROM users WHERE id=?"
// 	err = db.Get(&tempModel, query, userID)
// 	if err != nil {
// 		log.Fatal(err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"Query Exection Error": err})
// 		return
// 	}

// 	type tempResponse struct {
// 		Name     string `json:"name"`
// 		Age      int    `json:"age"`
// 		Email    string `json:"email"`
// 		CreateAt string `json:"createdAt"`
// 	}

// 	response := tempResponse{
// 		Name:     tempModel.Name,
// 		Age:      tempModel.Age,
// 		Email:    tempModel.Email,
// 		CreateAt: tempModel.CreatedAt.Time.Format("2006-01-02 15:04:05"),
// 	}

// 	c.JSON(http.StatusOK, gin.H{"SUCCESS": response})
// }

// func GetUserListing(c *gin.Context) {

// 	type userModel struct {
// 		ID       int       `db:"id"`
// 		Name     string    `db:"name"`
// 		Age      int       `db:"age"`
// 		Email    string    `db:"email"`
// 		CreateAt time.Time `db:"created_at"`
// 	}

// 	tempModels := []userModel{}

// 	query := "SELECT * FROM users"
// 	err = db.Select(&tempModels, query)
// 	if err != nil {
// 		log.Fatal(err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"Query Exection Error": err})
// 	}

// 	type tempResponse struct {
// 		Name     string `json:"name"`
// 		Age      int    `json:"age"`
// 		Email    string `json:"email"`
// 		CreateAt string `json:"created_at"`
// 	}

// 	response := []tempResponse{}

// 	for _, tempModel := range tempModels {
// 		response = append(response, tempResponse{
// 			Name:     tempModel.Name,
// 			Age:      tempModel.Age,
// 			Email:    tempModel.Email,
// 			CreateAt: tempModel.CreateAt.Format("2006-01-02 15:04:05"),
// 		})
// 	}

// 	c.JSON(http.StatusOK, gin.H{"SUCCESS": response})
// }

// func UpdateUserDetails(c *gin.Context) {
// 	userIDstr := c.Param("id")
// 	userID, err := strconv.Atoi(userIDstr)
// 	if err != nil {
// 		log.Fatal(err)
// 		c.JSON(http.StatusBadRequest, gin.H{"FAILED": err})
// 	}

// 	type updateUserRequest struct {
// 		Name  string `json:"name" binding:"required"`
// 		Age   int    `json:"age" binding:"required"`
// 		Email string `json:"email" binding:"required"`
// 	}

// 	var request updateUserRequest
// 	err = c.ShouldBindJSON(&request)
// 	if err != nil {
// 		log.Fatal(err)
// 		c.JSON(http.StatusBadRequest, gin.H{"FAILED": err})
// 	}

// 	query := "UPDATE users SET name=?, age=?, email=? WHERE id=?"
// 	_, err = db.Exec(query, request.Name, request.Age, request.Email, userID)
// 	if err != nil {
// 		log.Fatal(err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"Query Exection Error": err})
// 	}

// 	c.JSON(http.StatusOK, gin.H{"SUCCESS": "user data update successfully"})
// }

// func DeleteUser(c *gin.Context) {
// 	userIDstr := c.Param("id")
// 	userID, err := strconv.Atoi(userIDstr)
// 	if err != nil {
// 		log.Fatal(err)
// 		c.JSON(http.StatusBadRequest, gin.H{"FAILED": err})
// 	}

// 	query := "DELETE FROM users WHERE id=?"
// 	_, err = db.Exec(query, userID)
// 	if err != nil {
// 		log.Fatal(err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"Query Exection Error": err})
// 	}

// 	c.JSON(http.StatusOK, gin.H{"SUCCESS": "user deleted"})
// }
