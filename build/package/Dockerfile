FROM golang:1.21-alpine AS build 

WORKDIR /app
COPY ./ . 

RUN apk --no-cache add tzdata ca-certificates git
RUN CGO_ENABLED=0 GOOS=linux go build -C cmd/cli/ -o /app_bin

FROM scratch
COPY --from=build /app_bin /app_bin
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo

EXPOSE 8080
ENTRYPOINT ["/app_bin"]