apig new -u wantedly apig-sample
# fill models
apig gen
go get ./...
go build -o bin/server
