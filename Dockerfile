# Start by building the application.
FROM golang:1.18 as build

WORKDIR /go/src/dnf-proxy-cache
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/dnf-proxy-cache

# Now copy it into our base image.
FROM gcr.io/distroless/static-debian11
COPY --from=build /go/bin/dnf-proxy-cache /
CMD ["/dnf-proxy-cache"]