FROM golang:1.16-alpine

# appディレクトリの作成
RUN mkdir /go/src/app
# ワーキングディレクトリの設定
WORKDIR /go/src/app
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/app

RUN apk upgrade --update && \
    apk --no-cache add git

#RUN go get github.com/codegangsta/gin

#RUN go get -u github.com/cosmtrek/air && \
#    go build -o /go/bin/air github.com/cosmtrek/air

#CMD ["air", "-c", ".air.toml"]

RUN go mod download golang.org/x/text

CMD ["go","run","main.go"]