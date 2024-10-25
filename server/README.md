###Run server with docker

run commands:
docker build -t backend-go:local .
docker run -p 8080:8080 backend-go:local

or:
docker-compose up --build