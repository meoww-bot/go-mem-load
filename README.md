# go-mem-load


build
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -o bin/mem-load-amd64 .
```

run without params (default 50% memory load)
```
chmod +x bin/mem-load-amd64
./bin/mem-load-amd64 
```

run without params (default 50% memory load)
```
chmod +x bin/mem-load-amd64
./bin/mem-load-amd64 
```


run with params 80% load
```
chmod +x bin/mem-load-amd64
./bin/mem-load-amd64 0.8
```

