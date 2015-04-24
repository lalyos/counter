go build
GOOS=linux go build -o counter-linux .

docker build -t lalyos/counter .
docker build -f Dockerfile.cockroach -t lalyos/cockroach .
