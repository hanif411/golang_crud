FROM golang:1.26rc2-alpine3.23

ENV KEY_CRYPTO="rahasiasecretrahasiasecret123456"

WORKDIR /app

COPY . .

RUN go build -o crud

EXPOSE 5000

CMD ./crud