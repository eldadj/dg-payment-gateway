##
## Build
##
FROM golang:1.17.8-alpine3.15 as builder
COPY go.mod go.sum /go/src/github.com/eldadj/dgpg/
WORKDIR /go/src/github.com/eldadj/dgpg
RUN go mod download
COPY . /go/src/github.com/eldadj/dgpg
COPY .env_docker /go/src/github.com/eldadj/dgpg/.env
RUN go build -o build/dgpg github.com/eldadj/dgpg
#RUN go build -o ./dgpg
#EXPOSE 8080
#CMD [ "build/dgpg" ]
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/dgpg github.com/eldadj/dgpg

##
## Deploy
##
FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
WORKDIR /app/
COPY --from=builder /go/src/github.com/eldadj/dgpg/build/dgpg .
COPY --from=builder /go/src/github.com/eldadj/dgpg/conf/app.conf ./conf/app.conf
COPY --from=builder /go/src/github.com/eldadj/dgpg/.env .
RUN ls -la
EXPOSE 8080 8080
ENTRYPOINT ["/app/dgpg"]
