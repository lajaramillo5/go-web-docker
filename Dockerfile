FROM golang:alpine

WORKDIR /build

COPY go.mod .
#COPY go.sum .
RUN go mod download


COPY . .

RUN go build -o server .

WORKDIR /dist

RUN cp /build/server .

EXPOSE 8083

CMD [ "/dist/server" ]