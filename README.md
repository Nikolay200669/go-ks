## go-ks

A simple Golang application
HTTP server that listens on port 8989.
The server calculate factorial of a number and return the result.

## How to run the application
1. Clone the repository
```bash
git clone ... \ 
cd go-ks
```

1. Run the following command to start the server:
```bash
go run main.go
```

```bash
curl -X POST http://localhost:8989/calculate 
  -H 'Content-Type: application/json' \
  -d '{"a": 5, "b": 3}' && \ 
echo "END"
```

### Or
1. Use integration test:
```bash
go test -v
```
output should be:
```txt
=== RUN   TestCalcHandler
=== RUN   TestCalcHandler/TestCalcHandler
--- PASS: TestCalcHandler (0.00s)
    --- PASS: TestCalcHandler/TestCalcHandler (0.00s)
PASS
ok      github.com/Nikolay200669/go-ks  0.343s
```

### Or

1. Run Docker:
```bash
docker build -t go-ks . && \
docker run go-ks
```

### TODO:
- [ ] Add unit tests
- [ ] Refactor for use better practices (structs, interfaces, etc...)
- [ ] Logging
- [ ] Use environment variables
- [ ] Refactor struct response
- [ ] Graceful shutdown