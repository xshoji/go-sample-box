## About

packr's web app sample.

## Build and run

```
# Initialize command
go mod init "github.com/xshoji/go-sample-box/packr"
go mod tidy
```


```
// generage code
go generate ./...

// replace package name
sed -i '' "s/.*packrd\"/import _ \"github.com\/xshoji\/go-sample-box\/packr\/packrd\"/g" main-packr.go

// run
go run .


// build and run
go build -o /tmp/webapp .
/tmp/./webapp
```

## References

> gobuffalo/packr： The simple and easy way to embed static files into Go binaries.  
> https://github.com/gobuffalo/packr  

> kazu634/packr-example  
> https://github.com/kazu634/packr-example  
