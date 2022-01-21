FROM golang:1.17.6

WORKDIR /app
EXPOSE 8080
ENV VARIABLE="par def"

COPY . .

RUN go build -o web

CMD ./web
