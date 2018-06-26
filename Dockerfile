FROM golang:1.8

# Add the source code:
WORKDIR /app/
ADD public/ ./public
ADD templates/ ./templates
ADD ./*.go .

# Build it:
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o davidjreed.net .

# Executable container
FROM alpine

WORKDIR /app/

COPY --from=0 /app/ .

CMD ["/app/davidjreed.net"]

EXPOSE 8080
