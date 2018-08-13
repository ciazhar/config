# Config

>Multi Stage JSON Config File Manager for Go

## Running Project
```
go run main.go <yourJsonConfigFileName>
```

## How To Use
```golang
config.Load()
```

## Get Field Value
```golang
config.GetString("yourString")
config.GetInt("yourInteger")
config.GetBool("yourBool")
```