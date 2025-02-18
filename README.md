# ingestion-go
Repository includes authentication, message ingestion, real-time stock updates. 

Steps to Setup & produce updates locally

# Setup env variables in .env file 
cp .env-example .env

# install the dependecies 
make setup 

# start PostgreSQL and Kakfka 
make db-up       # Start PostgreSQL
make kafka-up    # Start Kafka & Zookeeper

# make the migrations up
make migrate-up 

# test the application 
make test 

# run the application 
make run 

# register a user - authentication 
POST localhost:8080/api/v1/auth/register
Body
{
  "email": "test@example.com",
  "password": "password123"
}

# login with the registered user and get the access token
POST localhost:8080/api/v1/auth/login
Body
{
  "email": "test@example.com",
}
Response 
{
    "access_token": "xxxx",
    "type": "Bearer",
    "expires_at": 3600000000000
}

# from now each request should have access token in header 
Authorization: Bearer xxxx

# produce message for stock updates 
POST localhost:8080/api/v1/send-message/stocks
Body
{
  "id":   230,
  "name": "NasInsurance",
  "price": 62
}
Header
Authorization: Bearer xxxx

Response : Status 201 

# Notes
	•	All API requests require authentication (except register/login).
	•	Kafka messages are consumed and persisted into PostgreSQL.
	•	Modify .env before running the application.
	•	Ensure consumer services are running locally to process Kafka messages.
