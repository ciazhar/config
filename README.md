# Config

>Multi Stage JSON Config File Manager for Go

## Running Project
```
go run main.go <yourJsonConfigFileName>
```

## How To Use
```golang
	conf :=config.Load()
```

## Get Field Value
```golang
conf.Get("port").Float()
conf.Get("database").Get("name").String()
```