FROM golang:1.21

WORKDIR /app
COPY . .
RUN go mod tidy && go build -buildvcs=false -o main . 
EXPOSE 8080
CMD [ "./main" ]