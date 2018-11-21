# Purchase System - 購物資料庫系統
## 編譯
```
go build -o (檔名) *.go
```

在 Linux 環境下，您能透過使用不同的 `GOOS` 和
`GOARCH` 變數來編譯不同系統的版本：

```
# 此例將編譯 Windows (AMD64) 的版本
GOOS=windows GOARCH=amd64 go build -o (檔名) *.go

# 此例將編譯 Windows (x86) 的版本
GOOS=windows GOARCH=386 go build -o (檔名) *.go

# 此例將編譯 Linux (x86) 的版本
GOOS=linux GOARCH=386 go build -o (檔名) *.go

# 此例將編譯 Darwin (macOS, AMD64) 的版本
GOOS=darwin GOARCH=amd64 go build -o (檔名) *.go
```

## 使用
打開編譯好的程式後輸入 usage 得知使用方法。

## 作者
- pan93412, 2018.
