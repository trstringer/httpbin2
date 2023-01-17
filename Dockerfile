FROM golang:1.19 AS builder
COPY . /var/app
WORKDIR /var/app
RUN CGO_ENABLED=0 go build -o httpbin2 .

FROM gcr.io/distroless/static
LABEL org.opencontainers.image.source https://github.com/trstringer/httpbin2
COPY --from=builder /var/app/httpbin2 /httpbin2
CMD ["/httpbin2"]
