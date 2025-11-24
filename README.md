# go-sample-box

Go's sample and snippet code.

```
# Initialize command
go mod init "github.com/xshoji/go-sample-box/${PWD##*/}"; go mod tidy

# Add module
go get golang.org/x/sync/errgroup

# Update go.mod
rm -rf go.*; go mod init "github.com/xshoji/go-sample-box/${PWD##*/}"; go mod tidy; go run main.go
```


## build

### Common command

```
go build -ldflags="-s -w" -trimpath -o /tmp/$(basename "$PWD") main.go
```

### Cross compile

```
APP=/tmp/app; go build -ldflags="-s -w"  -trimpath -o ${APP} main.go; chmod +x ${APP}
# APP=/tmp/tfr; GOOS=linux GOARCH=amd64   go build -ldflags="-s -w" -trimpath -o ${APP} main.go; chmod +x ${APP} # linux
# APP=/tmp/tfr; GOOS=darwin GOARCH=amd64  go build -ldflags="-s -w"  -trimpath -o ${APP} main.go; chmod +x ${APP} # macOS
# APP=/tmp/tfr; GOOS=windows GOARCH=amd64 go build -ldflags="-s -w"  -trimpath -o ${APP} main.go; chmod +x ${APP} # windows
```


## cannot resolve symbol golang

1. `rm -rf .idea`
2. Open subdirectory in IntelliJ
3. Build -> Build Project
 
