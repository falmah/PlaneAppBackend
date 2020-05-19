FROM base_image

COPY ./backend/operatorService/app/handler/* ./app/handler/
COPY ./backend/operatorService/app/app.go ./app/
COPY ./backend/operatorService/driver/* ./driver/operatorDriver/
COPY ./backend/operatorService/main.go ./
COPY ./backend/model/* ./model/

RUN /usr/local/go/bin/go build main.go

CMD ["./main"]
