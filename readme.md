# Go REST Api

## Initialize go project.

```go
go mod init github.com/sanjivpaul/studentapi
```

## Folder Structure

```go
cmd                  [cmd dir]
  |---studentapi     [app name]
      |--- main.go   [main entry file]

go.mod
```

## Run Go App

```bash
go run cmd/studentapi/main.go
```

## Start the server

```bash
go run cmd/studentapi/main.go -config config/local.yaml
```

Note: If we dont pass the config flag then we get error `-config config/local.yaml`
