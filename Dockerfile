FROM golang:1.24-alpine as build

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o webCloudApp ./cmd/server

FROM scratch
WORKDIR /app
COPY --from=build /app/webCloudApp .
EXPOSE 8000
ENTRYPOINT [ "./webCloudApp" ]