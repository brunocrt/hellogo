FROM golang as builder

WORKDIR /C/Users/boto/Sources/test/istio/

COPY hello.go .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /C/Users/boto/Sources/test/istio/app .
ENV PORT 8080
CMD [ "./app" ]