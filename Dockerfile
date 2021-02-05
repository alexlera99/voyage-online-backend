FROM golang

COPY . .

RUN go get -t "github.com/gorilla/mux"


RUN go get -t "github.com/rs/cors"


EXPOSE 4443

CMD go run main.go