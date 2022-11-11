FROM golang:1.18-alpine as builder
LABEL maintainer="Kien Nguyen-Tuan <kiennt2609@gmail.com>"
ENV APP=$GOPATH/src/goload
COPY . $APP
WORKDIR $APP
RUN go build -ldflags "-s -w" -o /bin/goload *.go

FROM alpine:3.15
LABEL maintainer="Kien Nguyen-Tuan <kiennt2609@gmail.com>"
COPY --from=builder /bin/goload /bin/goload
USER nobody
ENTRYPOINT [ "/bin/goload" ]
