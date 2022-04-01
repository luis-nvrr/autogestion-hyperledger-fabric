package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

var jwtKey = []byte("my_secret_key")

var (
	router = gin.Default()
)

type CreateStudentRequest struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Year     int    `json:"year"`
}

type CreateGradeRequest struct {
	Grade        int                   `json:"grade"`
	Date         string                `json:"date"`
	Student      *CreateStudentRequest `json:"student"`
	Instance     string                `json:"instance"`
	Observations string                `json:"observations"`
}

type Claims struct {
	Username     string `json:"username"`
	Organization string `json:"organization"`
	jwt.StandardClaims
}

func getClaims(tokenString string) (*Claims, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !tkn.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

func getWallet() (*gateway.Wallet, error) {
	walletPath := filepath.Join("..", "enroll-user", "wallet")
	wallet, err := gateway.NewFileSystemWallet(walletPath)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func connectToGatway(claims *Claims, wallet *gateway.Wallet) (*gateway.Gateway, error) {
	ccpPath := filepath.Join(
		"..",
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		fmt.Sprintf("%s.example.com", claims.Organization),
		fmt.Sprintf("connection-%s.yaml", claims.Organization),
	)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, claims.Username),
	)
	if err != nil {
		return nil, err
	}

	return gw, nil
}

func submitTransaction(c *gin.Context, query string, arg ...string) ([]byte, error) {
	if err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true"); err != nil {
		c.JSON(http.StatusInternalServerError, "error setting env variable")
		log.Printf("Error setting DISCOVERY_AS_LOCALHOST environemnt variable: %v", err)
		return nil, err
	}

	tokenString := c.Request.Header.Get("Authorization")
	claims, err := getClaims(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "error reading jwt")
		log.Printf("failed to get jwt claims: %v", err)
		return nil, err
	}

	wallet, err := getWallet()
	if err != nil {
		c.JSON(http.StatusUnauthorized, "wallet not found")
		log.Printf("failed to create wallet: %v", err)
		return nil, err
	}
	if !wallet.Exists(claims.Username) {
		c.JSON(http.StatusUnauthorized, "user not found")
		log.Printf("failed to find user certificate")
		return nil, err
	}

	gw, err := connectToGatway(claims, wallet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to connect to gateway")
		log.Printf("failed to connect to gateway: %v", err)
	}
	defer gw.Close()

	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to get channel")
		log.Printf("Failed to get network: %v", err)
		return nil, err
	}

	log.Printf("--> Submit Transaction: %s\n", query)
	contract := network.GetContract("basic")
	result, err := contract.SubmitTransaction(query, arg...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error submiting transaction")
		log.Printf("Error submiting: %v", err)
		return nil, err
	}
	return result, nil
}

func CreateGrade(c *gin.Context) {
	var request CreateGradeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, "couldn't serialize body")
		log.Printf("failed to bind json: %v", err)
		return
	}
	query := "CreateGrade"
	result, err := submitTransaction(c, query,
		fmt.Sprintf("%v", request.Grade),
		request.Date,
		fmt.Sprintf("%d", request.Student.Id),
		request.Student.Name,
		request.Student.LastName,
		fmt.Sprintf("%d", request.Student.Year),
		request.Instance,
		request.Observations,
	)

	if err != nil {
		return
	}

	c.JSON(http.StatusCreated, result)
}

func main() {
	router.Use(cors.Default())
	router.POST("/api/grades", CreateGrade)
	router.Run(":8081")
}
