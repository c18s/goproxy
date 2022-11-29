FROM golang:1.19-bullseye as build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-X main.cartPath=/etc/proxy/" -o proxy
RUN ./proxy keygen -C proxy

FROM scratch
ENV TZ Asia/Bangkok
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/proxy /usr/bin/
COPY --from=build /app/proxy.crt /app/proxy.key /etc/proxy/
ENTRYPOINT ["proxy"]
CMD ["--version"]
