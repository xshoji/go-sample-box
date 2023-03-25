# go-sample-box

Go's sample and snippet code.

```
# Initialize command
go mod init "github.com/xshoji/go-sample-box/${PWD##*/}"; go mod tidy

# Update go.mod
rm -rf go.*; go mod init "github.com/xshoji/go-sample-box/${PWD##*/}"; go mod tidy; go run main.go
```
