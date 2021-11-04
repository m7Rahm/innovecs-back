
FROM golang

WORKDIR /go/src/app
COPY ./ .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build .
CMD ["app"]
EXPOSE 4000