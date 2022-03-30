package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

var jwtKey = []byte("my_secret_key")

type User struct {
	Username     string
	Password     string
	Organization string
}

var (
	router = gin.Default()
	db     *mongo.Database
)

func ConnDB() {
	uri := "mongodb://root:rootpassword@localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	db = client.Database("users")
	fmt.Println("Successfuly connected to the database.")
}

func Register(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, "couldn't serialize body")
		return
	}

	hashedPassword, _ := HashPassword(user.Password)

	usersC := db.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)

	_, err := usersC.InsertOne(ctx, bson.M{
		"username":     user.Username,
		"password":     hashedPassword,
		"organization": user.Organization,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "can't insert to DB")
		return
	}

	c.JSON(http.StatusCreated, user)
}

type Claims struct {
	Username     string `json:"username"`
	Organization string `json:"organization"`
	jwt.StandardClaims
}

func Login(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, "couldn't serialize body")
		return
	}

	var userResult User
	usersC := db.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	err := usersC.FindOne(ctx, bson.M{"username": user.Username}).Decode(&userResult)
	if err != nil {
		c.JSON(http.StatusBadRequest, "user not found")
		return
	}

	if !CheckPasswordHash(user.Password, userResult.Password) {
		c.JSON(http.StatusForbidden, "bad login")
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username:     user.Username,
		Organization: user.Organization,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error with jwt")
		return
	}

	c.SetCookie("token", tokenString, expirationTime.Minute(), "", "", false, false)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	router.POST("/login", Login)
	router.POST("/users", Register)
	ConnDB()
	router.Run(":8080")
}
