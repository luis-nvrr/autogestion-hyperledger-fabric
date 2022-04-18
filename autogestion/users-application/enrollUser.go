package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

const (
	org1MSP        = "Org1MSP"
	org3MSP        = "Org3MSP"
	org1HostName   = "org1.example.com"
	org3HostName   = "org3.example.com"
	org1User       = "User1@org1.example.com"
	org3User       = "User3@org3.example.com"
)

var now = time.Now()

func main() {
	log.Println("============ enrollUser application-golang starts ============")

	org := os.Args[1]
	userId := os.Args[2]

	switch org {
	case "org1":
		err := addUserToWallet(userId, org1MSP, org1HostName, org1User, "wallet")
		if err != nil {
			log.Println(err)
		}
	case "org3":
		err := addUserToWallet(userId, org3MSP, org3HostName, org3User, "wallet")
		if err != nil {
			log.Println(err)
		}
	default:
		log.Println("invalid organization")
	}
}

func addUserToWallet(user, orgMSP, orgHostName, orgUser, walletName string) error {
	fmt.Printf("\n--> Register and enrolling new user")

	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	if err != nil {
		log.Fatalf("Error setting DISCOVERY_AS_LOCALHOST environemnt variable: %v", err)
	}

	wallet, err := gateway.NewFileSystemWallet(walletName)
	if err != nil {
		log.Fatalf("Failed to create wallet: %v", err)
	}

	if wallet.Exists(user) {
		return fmt.Errorf("user %s already exists", user)
	}

	credPath := filepath.Join(
		"..",
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		orgHostName,
		"users",
		orgUser,
		"msp",
	)

	certPath := filepath.Join(credPath, "signcerts", "cert.pem")
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore")
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return fmt.Errorf("keystore folder should have contain one file")
	}

	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity(orgMSP, string(cert), string(key))

	return wallet.Put(user, identity)
}
