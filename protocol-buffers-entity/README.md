## About

Sample of entity generation using protocol buffer.

## Requirements

```
brew install protobuf
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
```

## Build and run

```
// generage code
./generate_model.sh -m entity

// run
go run cmd/main/main.go
{
  "id": 1,
  "name": "a",
  "nickname": "a",
  "age": 16,
  "birth": {
    "datetime": "2020-04-23T17:36:22.421530Z",
    "weight": 12,
    "hospital": "hugeHospital"
  },
  "addresss": {
    "country": "JP",
    "zipCode": 1,
    "state": "test",
    "city": "city"
  }
}
```

## References

> Protocol BuffersをJSONに変換する - c-bata web  
> https://nwpct1.hatenablog.com/entry/convert-protobuf-to-json

