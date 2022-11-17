# How to run

## Clone this repository
```
use this following command
git clone https://github.com/capstone-project-alterra-9/be-loyalty-point.git
```

## Run apps using docker 
```
docker-compose -f docker-compose.yml up
```

## Standard running apps
```
npm run main.go
```

## Run unit testing
```
go test -v -coverprofile=cover.out && go tool cover -html=cover.out
```