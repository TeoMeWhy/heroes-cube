FROM golang:1.24.2

WORKDIR /app

COPY . .

RUN make build

CMD [ "make", "run" ]