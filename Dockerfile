FROM golang:latest as build

WORKDIR /build
COPY ./ ./
RUN CGO_ENABLED=0 go build -o ./health

FROM scratch
COPY --from=build /build/health /health
ENTRYPOINT ["/health"]