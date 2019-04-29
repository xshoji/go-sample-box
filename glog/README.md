## About

glog

## Build and run

```
// basic
dep ensure
go run main.go

// level = INFO
// -stderrthreshold : 出力するログレベルを設定する
go run main.go -stderrthreshold=INFO

// level = WARNING
go run main.go -stderrthreshold=WARNING

// level = ERROR
go run main.go -stderrthreshold=ERROR

// Fatal error occurred!
go run main.go -stderrthreshold=ERROR -fatal
```

## References

> go言語におけるロギングについて ｜ さにあらず  
> https://blog.satotaichi.info/logging-frameworks-for-go/
