ARG IMAGE=golang:1.12.6-alpine3.9
ARG WORKDIR=/app

FROM ${IMAGE} as builder
LABEL author.email=yaphetsglhf@gmail.com

RUN mkdir -p $WORKDIR
WORKDIR $WORKDIR

RUN echo "http://mirrors.aliyun.com/alpine/v3.8/main/" > /etc/apk/repositories && \
    echo "http://mirrors.aliyun.com/alpine/v3.8/community" >>  /etc/apk/repositories

ARG GOPROXY=https://goproxy.cn
ENV GOPROXY=$GOPROXY
ENV GO111MODULE=on

COPY go.mod go.sum $WORKDIR/
RUN go mod download
COPY . $WORKDIR/

RUN go build -o server


FROM ${IMAGE}
LABEL author.email=yaphetsglhf@gmail.com

RUN apk update && \
    apk add --no-cache tzdata \
                       gcc build-base curl

ENV LC_ALL=en_CA.UTF-8
ENV LANG=en_CA.UTF-8
ENV LANGUAGE=en_CA.UTF-8
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN mkdir -p $WORKDIR
WORKDIR $WORKDIR

COPY --from=builder $WORKDIR/server $WORKDIR/.
CMD ["/server"]

ARG APP_NAME=ddd
ENV PORT=8080
EXPOSE $PORT
