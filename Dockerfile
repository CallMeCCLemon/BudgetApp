FROM --platform=linux/amd64 golang:1.22
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /budgeting-app

EXPOSE 8080

CMD ["/budgeting-app"]
