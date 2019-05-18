
## Memo

```
// Create a self-signed Certificate
openssl genrsa -des3 -out server.key 1024
openssl req -new -key server.key -out server.csr
openssl rsa -in server.key.org -out server.key
```

```
// run
go run main.go -tls
```

## References


> Go言語とHTTP2 ｜ SOTA  
> https://deeeet.com/writing/2015/11/19/go-http2/

> オレだよオレオレ認証局で証明書つくる - Qiita  
> https://qiita.com/ll_kuma_ll/items/13c962a6a74874af39c6

> How to create a self-signed Certificate  
> https://www.akadia.com/services/ssh_test_certificate.html
