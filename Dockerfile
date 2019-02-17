FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY build/OpenLibraryServer /OpenLibraryServer
COPY wishList.sqlite3 /

COPY swaggerui /swaggerui
COPY swagger.yaml /swaggerui/swagger.yaml
ENTRYPOINT ["/OpenLibraryServer"]