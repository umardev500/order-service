FROM golang as dev

WORKDIR /app

COPY . .

EXPOSE 5011

CMD air
