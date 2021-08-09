# work on current directory

> Windows

## compile and run

1. `go build -o ./5_1.exe ./5_1.go`
2. `go build -o ./fetch.exe ../fetch/fetch.go`
3. `./fetch.exe https://golang.google.cn | ./5_1.exe`

## run at once

```cmd
go run ..\fetch\fetch.go https://baidu.com | go run .\5_1.go 
```
