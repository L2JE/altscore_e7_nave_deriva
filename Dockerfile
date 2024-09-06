FROM golang:alpine3.20 AS compiler_stg

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY ./src/go.mod ./src/go.sum ./
RUN go mod download && go mod verify

COPY ./src ./
RUN CGO_ENABLED=0 GOOS=linux go build -v -o ./altscore_e7_nave_deriva

FROM alpine:3.20
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=compiler_stg /usr/src/app/altscore_e7_nave_deriva .

CMD ["./altscore_e7_nave_deriva"]