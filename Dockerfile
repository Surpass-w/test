FROM golang:1.18.10

LABEL maintainer = "wangpeng@moresec.cn"

WORKDIR /app
COPY . .

RUN go mod tidy && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build --trimpath --ldflags "-w -s" -o test main.go

EXPOSE 8000

CMD ["./test"]
