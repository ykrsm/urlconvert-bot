# Slack URL converter

sup_id=100 to http..../100

# Compile for linux
```
dep ensure
env GOOS=linux GOARCH=386 go build -o main
```
# Start docker
```
docker-compose build
docker-compose up
```

# Feeling lazy? (one liner)
```
env GOOS=linux GOARCH=386 go build -o main && docker-compose up --build
```