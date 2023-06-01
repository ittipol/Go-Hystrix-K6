# Hystrix Circuit Breaker

## Go Packages

- hystrix [https://pkg.go.dev/github.com/afex/hystrix-go/hystrix](https://pkg.go.dev/github.com/afex/hystrix-go/hystrix)
- fiber [https://pkg.go.dev/github.com/gofiber/fiber/v2](https://pkg.go.dev/github.com/gofiber/fiber/v2)

``` bash
# Install hystrix package
go get github.com/afex/hystrix-go/hystrix

# Install fiber package
go get github.com/gofiber/fiber/v2
```

## Hystrix dashboard
- Docker https://hub.docker.com/r/mlabouardy/hystrix-dashboard

## Start server and application
``` bash
docker compose up -d --build
```

## Dashboard setup
1. Open http://localhost:9002/hystrix
2. Input host (Metric) "http://host.docker.internal:8081"
3. Input Delay
4. Input Title
5. Click "Monitor Stream" button

## Test
``` bash
# run K6 script
docker compose run --rm k6 run /scripts/test.js
```