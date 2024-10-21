FROM golang:1.22-bookworm

RUN apt-get update && apt-get install -y make

RUN curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | bash
RUN apt-get install migrate=4.18.1

WORKDIR /app

COPY . /app

RUN curl -s https://raw.githubusercontent.com/vishnubob/wait-for-it/refs/heads/master/wait-for-it.sh > wait-for-it.sh
RUN ["chmod", "+x", "wait-for-it.sh"]
RUN go generate ./...
RUN go build -o /app/bin /app/cmd/.

CMD ["/app/bin"]
