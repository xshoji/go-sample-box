## About

packr's web app sample.

## Build and run

```
// generage code
go generate ./...

// replace package name
sed -i '' "s/.*packrd\"/import _ \"github.com\/xshoji\/go-sample-box\/webserver-packr\/packrd\"/g" main-packr.go

// build
go build -o /tmp/webapp .

// run
/tmp/./webapp
```

## References

> gobuffalo/packr： The simple and easy way to embed static files into Go binaries.  
> https://github.com/gobuffalo/packr  

> kazu634/packr-example  
> https://github.com/kazu634/packr-example  
