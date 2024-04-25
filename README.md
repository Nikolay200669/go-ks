## go-ks

A simple Golang application
HTTP server that listens on port 8989.
The server calculate factorial of a number and return the result.

## How to run the application
1. Clone the repository
```bash
git clone https://github.com/Nikolay200669/go-ks.git && \ 
cd go-ks
```

2. Run the following command to start the server:
```bash
go run main.go
```

```bash
curl -X POST http://localhost:8989/calculate \
  -H 'Content-Type: application/json' \
  -d '{"a": 5, "b": 3}'
```

### Or
3. Use integration test:
```bash
go test -v
```
output should be:
```txt
=== RUN   TestCalcHandler
=== RUN   TestCalcHandler/Run_test_OK
=== RUN   TestCalcHandler/Run_test_middleware
--- PASS: TestCalcHandler (0.00s)
    --- PASS: TestCalcHandler/Run_test_OK (0.00s)
    --- PASS: TestCalcHandler/Run_test_middleware (0.00s)
PASS
ok      github.com/Nikolay200669/go-ks  0.335s

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