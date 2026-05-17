FROM alpine:latest
WORKDIR /app

COPY main .

RUN mkdir /app/database
EXPOSE 8080
CMD ["./main"]