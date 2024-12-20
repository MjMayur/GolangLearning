package main

import (
	// "database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

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
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE,PATCH, OPTIONS")

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
	router.Use(CORSMiddleware())
	router.POST("/user/create/", CreateUserHandler)
	router.GET("/user/list/", ListUser)
	router.DELETE("/user/delete/:id", DeleteUser)
	router.GET("/user/get/:id", GetDetails)
	router.PATCH("/user/update/:id", UpdateUser)
	router.Run(":8000")
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

type responseStruct struct {
	Code    int    `json:"code"`
	Massage string `json:"massage"`
}

type createUserRequest struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type listRes struct {
	Id      int    `db:"id"`
	Name    string `db:"name"`
	Email   string `db:"email"`
	Phone   string `db:"phone"`
	Message string `db:"message"`
}

type tempResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Message string `json:"message"`
}

func validationEmail(email string) error {
	// Regular expression for validating an email

	var emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)

	if !re.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

func ListUser(c *gin.Context) {

	tempModels := []listRes{}

	query := "SELECT * FROM users"
	err = db.Select(&tempModels, query)
	if err != nil {
		log.Println("Email validation error:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	response := []tempResponse{}

	for _, tempModel := range tempModels {
		response = append(response, tempResponse{
			ID:      tempModel.Id,
			Name:    tempModel.Name,
			Phone:   tempModel.Phone,
			Email:   tempModel.Email,
			Message: tempModel.Message,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"Users":   response,
		"code":    200,
		"message": "Success recorder added successfully.",
	})
}

func CreateUserHandler(c *gin.Context) {

	var request createUserRequest

	err = c.ShouldBindJSON(&request)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"FAILED": err})
		return
	}
	// Validate email format
	if err := validationEmail(request.Email); err != nil {
		log.Println("Email validation error:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return

	}
	query := "INSERT INTO users (name, email, phone, message) VALUES (?, ?, ?, ?)"
	result, err := db.Exec(query, request.Name, request.Email, request.Phone, request.Message)
	if err != nil {
		res := responseStruct{
			Code:    400,
			Massage: "Internal Server Error.",
		}
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"Query Exection Error": res})
		return
	}

	userID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"Last Insert ID getting Error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"SUCCESS": userID,
		"code":    200,
		"message": "Success recorder added successfully.",
	})
}

func DeleteUser(c *gin.Context) {
	userIDstr := c.Param("id")
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		log.Println("Email validation error:", err)
		c.JSON(http.StatusForbidden, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	query := "DELETE FROM users WHERE id=?"
	_, err = db.Exec(query, userID)
	if err != nil {
		res := responseStruct{
			Code:    400,
			Massage: "Internal Server Error.",
		}
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	res := responseStruct{
		Code:    200,
		Massage: "Record Deleted successfully",
	}
	c.JSON(http.StatusOK, res)
}

func GetDetails(c *gin.Context) {
	userIDstr := c.Param("id")
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		log.Println("Email validation error:", err)
		c.JSON(http.StatusForbidden, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}
	tempModels := listRes{}

	query := "SELECT * FROM users WHERE id=?"
	err = db.Get(&tempModels, query, userID)
	if err != nil {
		res := responseStruct{
			Code:    400,
			Massage: "Internal Server Error.",
		}
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"Query Exection Error": res})
		return
	}

	response := tempResponse{
		ID:      tempModels.Id,
		Name:    tempModels.Name,
		Phone:   tempModels.Phone,
		Email:   tempModels.Email,
		Message: tempModels.Message,
	}

	c.JSON(http.StatusOK, gin.H{
		"User":    response,
		"code":    200,
		"message": "Success recorder added successfully.",
	})
}

func UpdateUser(c *gin.Context) {
	userIDstr := c.Param("id")
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		log.Println("Email validation error:", err)
		c.JSON(http.StatusForbidden, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	var request createUserRequest
	err = c.ShouldBindJSON(&request)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"FAILED": err})
		return
	}
	// Validate email format
	if err := validationEmail(request.Email); err != nil {
		log.Println("Email validation error:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return

	}
	query := "UPDATE users SET name=?, phone=?, email=?,message=? WHERE id=?"
	_, err = db.Exec(query, request.Name, request.Phone, request.Email, request.Message, userID)
	if err != nil {
		res := responseStruct{
			Code:    400,
			Massage: "Internal Server Error.",
		}
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"Query Exection Error": res})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Success recorder Updated successfully.",
	})
}
