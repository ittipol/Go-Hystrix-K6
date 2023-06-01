FROM golang:1.19  as build

WORKDIR /go-app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o ./app

# =================================================
FROM golang:1.19  as production

WORKDIR /go-app

COPY --from=build /go-app/app ./

EXPOSE 8080 8081

ENTRYPOINT ["./app"]