FROM golang:alpine

ENV PORT=8000

WORKDIR /go/src/app

COPY . .

RUN apk update && apk add --no-cache git && apk add --no-cache make

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE $PORT

RUN make build

WORKDIR /go/src/app/bin


CMD ["app"]