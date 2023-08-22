FROM golang:1.21 as build
COPY main.go ./main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -a -installsuffix cgo -ldflags="-w -s" \
    -o /bin/my-web-app ./main.go

FROM scratch
COPY --from=build /bin/my-web-app /bin/my-web-app
CMD ["/bin/my-web-app"]
