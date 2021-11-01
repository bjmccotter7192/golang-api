# GOLANG API

## This is a simple example of a Go API using:
1. Docker
2. Docker-compose
3. Built in go packages

## How to run using Docker-compose (Easiest)
1. Clone repository
2. ```cd golang-api```
3. ```docker-compose up -d``` //This will allow the container to run in the background
4. Open http://localhost:8081/ in your browser and enjoy

## How to run using Docker
1. Clone repository
2. ```cd golang-api```
3. ```docker build -t golang-api .```
4. ```docker run -p 8081:8081 -it golang-api```
5. Open http://localhost:8081/ in your browser and enjoy

## How to run locally
1. Make sure Go is installed and added to your PATH
2. Clone Repository
3. ```cd golang-api```
4. ```go run main.go```

## Available Routes
- / (Just returns a home string)
- /albums (Returns a JSON list of albums statically defined inside of main.go)