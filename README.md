# Autogestion blockchain

Autogestion blockchain is a Hyperledger Fabric based blockchain network. The goal is to build a scalable solution to keep record of students grades.


## Requirements

- Install the Hyperledger Fabric prerequisites. Read more on [Hyperledger prerequisites](https://hyperledger-fabric.readthedocs.io/en/latest/prereqs.html). This includes:
   - Git
   - cURL
   - Docker
   - Docker-compose
   - Go 


## Installation

1. Clone the repository
   
   ```bash
    git clone https://github.com/luis-nvrr/autogestion-hyperledger-fabric && cd autogestion-hyperledger-fabric
   ```

2. Install the Hyperledger Fabric binaries & docker images

    ```bash
    bash install.sh
    ```

## Instructions to create a grade

1. Bring up the blockchain network with a channel named mychannel

    ```bash
    cd test-network
    ./network.sh up createChannel -c mychannel -ca
    ```
2. Deploy the smartcontract to the channel

    ```bash
    ./network.sh deployCC -ccn basic -ccp ../autogestion/chaincode-go -ccl go -c mychannel
    ```

3. Create a test teacher identity in the network

    ```bash
    cd ../autogestion/users-application
    go run enrollUser.go org1 test
    ```

4. Bring up the MongoDB database
   
   ```bash
   cd ../mongo-database
   docker-compose up -d
   ```

5. Create a user in the database for the test user

    ```bash
    cd ../authentication-service
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

6. Log in the test user to generate a JWT token

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

7. Bring up the grades service
   
   ```bash
   cd ../grades-service
   go run server.go
   ```

8. Create a grade. Replace {$TOKEN} with the JWT token obtained in the last step.

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
    
    
## Instructions to stop the project

1. Go to the root directory of the project. Shut down the blockchain.

    ```bash
    cd test-network
    ./network.sh down
    ```

2. Stop the database container.

    ```bash
    docker-compose down
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
![Architecture diagram](https://github.com/luis-nvrr/autogestion-hyperledger-fabric/blob/main/architecture-diagram.png)


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
