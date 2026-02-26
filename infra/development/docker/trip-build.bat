set CGO_ENABLED=0
set GOOS=linux
set GOARCH=arm64
go build -o build/trip-service ./services/trip-service/cmd