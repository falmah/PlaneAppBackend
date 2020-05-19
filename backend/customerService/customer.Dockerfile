FROM base_image

COPY ./backend/customerService/app/handler/* ./app/handler/
COPY ./backend/customerService/app/app.go ./app/
COPY ./backend/customerService/driver/* ./driver/customerDriver/
COPY ./backend/customerService/main.go ./
COPY ./backend/model/* ./model/

RUN /usr/local/go/bin/go build main.go

CMD ["./main"]
