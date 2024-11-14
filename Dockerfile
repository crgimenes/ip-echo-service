# Stage 1 - Build
FROM golang:latest as build
WORKDIR /build

LABEL maintainer="crg@crg.eti.br"
LABEL version="0.0.1"
LABEL description="IP Echo Service"

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux \
    go build -a -installsuffix cgo -ldflags '-s -w' -o ip-echo-service

# Stage 2 - Deploy
FROM scratch
COPY --from=build /build/ip-echo-service /ip-echo-service
EXPOSE 8001
ENTRYPOINT ["/ip-echo-service"]

