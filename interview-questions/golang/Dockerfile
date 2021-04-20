FROM golang:1.16.3-alpine3.13 as builder

RUN mkdir /app
WORKDIR /app

ADD go.mod /app
#ADD go.sum .

RUN go mod download

ADD . /app

# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /go/bin/schellar .
RUN go build -o /bin/app


FROM alpine:3.13.4

COPY --from=builder /bin/app /bin/
ADD startup.sh /

CMD ["sh","startup.sh"]⏎

