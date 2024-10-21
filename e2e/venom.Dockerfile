FROM ovhcom/venom:latest

RUN apk --no-cache add bash
RUN apk --no-cache add curl
RUN apk --no-cache add mysql-client

WORKDIR /workdir

COPY seeds.sql seeds.sql

RUN curl -s https://raw.githubusercontent.com/vishnubob/wait-for-it/refs/heads/master/wait-for-it.sh > wait-for-it.sh
RUN ["chmod", "+x", "wait-for-it.sh"]

ENTRYPOINT []
