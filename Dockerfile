FROM golang:latest as build
WORKDIR /build

COPY main.go .
RUN go build -o health

FROM scratch
COPY --from=build /build/main.go .
ENTRYPOINT ["/health"]