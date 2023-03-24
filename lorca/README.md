## About

lorca's sample application.

## Build and run

```
# Initialize command
go mod init "github.com/xshoji/go-sample-box/lorca"
go mod tidy
```

```
// generage code
go generate ./...

// build
go build -o /tmp/app .

```

## References

> gobuffalo/packr： The simple and easy way to embed static files into Go binaries.  
> https://github.com/gobuffalo/packr  

> kazu634/packr-example  
> https://github.com/kazu634/packr-example  
