GOOS=linux GOARCH=amd64 go build -o time-test
docker build -t jst-time-test1 -f Dockerfile1 .
docker build -t jst-time-test2 -f Dockerfile2 .
docker build -t jst-time-test3 -f Dockerfile3 .
