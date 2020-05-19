FROM golang:1.13
COPY ./QA-app-halocom /usr/src/QA-app-halocom
WORKDIR /usr/src/QA-app-halocom
RUN go build -o app .
RUN go get github.com/codegangsta/gin
# for dev
CMD ["gin","--appPort","8080","--all", "-i"]
# for production
# CMD ["./app"] 
