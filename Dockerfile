FROM golang:1.22.4 AS builder
WORKDIR /srv/shop-app
COPY . .
RUN CGO_ENABLED=0 go build -gcflags="all=-N -l" -o shop .

FROM golang:1.22.4
WORKDIR /srv/shop-app
COPY  --from=builder /srv/shop-app/config.json .
COPY  --from=builder /srv/shop-app/shop .

RUN CGO_ENABLED=0 go install -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv@latest

EXPOSE 8080

CMD [ "/go/bin/dlv", "--listen=:8080", "--headless=true", "--log=true", "--accept-multiclient", "--api-version=2", "exec", "/srv/shop-app/shop" ]

# CMD [ "./shop" ]