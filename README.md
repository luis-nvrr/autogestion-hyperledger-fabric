# Autogestion blockchain

Autogestion blockchain is a Hyperledger Fabric based blockchain network. The goal is to build a scalable solution to keep record of students grades.


## Requirements

- Install the Hyperledger Fabric prerequisites. Read more on [Hyperledger prerequisites](https://hyperledger-fabric.readthedocs.io/en/latest/prereqs.html). This includes:
   - Git
   - cURL
   - Docker
   - Docker-compose
   - Go 

- Install the Hyperledger Fabric binaries. Read more on [Hyperledger binaries](https://hyperledger-fabric.readthedocs.io/en/latest/install.html)

    ```bash
    curl -sSL https://bit.ly/2ysbOFE | bash -s
    ```

## Instructions to create a grade

1. Clone the repository
   
   ```bash
    git clone https://github.com/luis-nvrr/autogestion-hyperledger-fabric && cd autogestion-hyperledger-fabric
   ```

2. Bring up the blockchain network with a channel named mychannel

    ```bash
    cd test-network
    ./network.sh up createChannel -c mychannel -ca
    ```
3. Deploy the smartcontract to the channel

    ```bash
    ./network.sh deployCC -ccn basic -ccp ../autogestion/chaincode-go -ccl go -c mychannel
    ```

4. Create a test teacher identity in the network

    ```bash
    cd ../autogestion/enroll-user
    go run enrollUser.go org1 test
    ```

5. Bring up the MongoDB database
   
   ```bash
   cd ../mongo-database
   docker-compose up -d
   ```

6. Create a user in the database for the test user

    ```bash
    cd ../user-service
    go run server.go
    ```
    ```bash
    curl --request POST \
    --url http://localhost:8080/api/users \
    --header 'Content-Type: application/json' \
    --data '{
	    "username": "test",
	    "password": "test",
	    "organization": "org1"
        }'
    ```

7. Log in the test user to generate a JWT token

    ```bash
    curl --request POST \
    --url http://localhost:8080/api/users/auth \
    --header 'Content-Type: application/json' \
    --data '{
	    "username": "test",
	    "password": "test"
        }'
    ```

    The response should look something like this: 
    ```bash
    {
	    "username": "test",
	    "organization": "org1",
	    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QiLCJvcmdhbml6YXRpb24iOiJvcmcxIiwiZXhwIjoxNjQ4NzgxNTI0fQ.rIKFP4a-jJCtq0Tx5wEDJ9ZRK8g2a3wNPmUT1kKWENI"
    }
    ```

8. Bring up the grades service
   
   ```bash
   cd ../grades-service
   go run server.go
   ```

9. Create a grade. Replace {$TOKEN} with the JWT token obtained in the last step.

    ```bash
    curl --request POST \
    --url http://localhost:8081/api/grades \
    --header 'Authorization: {$TOKEN}' \
    --header 'Content-Type: application/json' \
    --data '{
	    "grade": 8,
	    "date": "2020-01-02",
	    "student": { 
		    "id": 79581,
		    "name": "luis",
		    "lastName": "navarro",
		    "year": 2020
	        },
	    "instance": "parcial",
	    "observations": "nada"
    }'
    ```


## Architecture

The network consists of:
- 3 Peer nodes
- 1 Orderer node
- 4 Certification authorities
- 1 Users service
- 1 Users database
- 1 Grades service
- 1 Web application

The architecture-diagram.jpg file shows a rough diagram of the blockchain network.


## Smartcontract

You can take look at the smartcontract at:

```bash
cd ./autogestion/blockchain-go/chaincode/
cat smartcontract.go
```

## Some gotchas

- org1 is intended to represent teachers. This is why we used this organization for the test user in the example.
- org2 represents the faculty authorities
- org3 represents students. That's why in the smartcontract the validation on which users can submit grades is done against org3 users.
