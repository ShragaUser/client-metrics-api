#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git make
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make build

#final stage
FROM scratch
COPY --from=builder /go/src/app/client-metrics /client-metrics
EXPOSE 9091
ENTRYPOINT ["/client-metrics"]