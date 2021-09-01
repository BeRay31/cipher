FROM golang as build
WORKDIR /app
COPY . .
RUN go get
RUN go build -o app main.go

FROM alpine
WORKDIR /
COPY --from=build /app /
ENTRYPOINT [ "/app" ]

