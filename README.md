## Project
Golang todo application, a simple one.

## Architecture
Code architecture using Clean Architecture by uncle bob, which highlighted couple thigns
1. Entity
2. Usecase
3. Repository

The goals is to abstract the whole implementation of one another, which mean if i want to change my DB for example, into MYSQL it will be relatively easy without breaking the existing code.

## How to run
First of all, please fill your db user and password into `env/config.json`

```bash
   go mod tidy
   go run cmd/server/main.go
```
