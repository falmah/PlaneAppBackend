FROM base_image

COPY ./app/handler/* ./app/handler/
COPY ./app/app.go ./app/
COPY ./driver/* ./driver/customerDriver/
COPY ./main.go ./

RUN /usr/local/go/bin/go build main.go

CMD ["./main"]
