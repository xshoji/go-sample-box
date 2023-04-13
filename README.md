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


## cannot resolve symbol golang

1. `rm -rf .idea`
2. Open subdirectory in IntelliJ
3. Build -> Build Project
 
