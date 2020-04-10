ARG IMAGE=golang:1.13-alpine

FROM ${IMAGE} as builder
ARG WORKDIR=/app
MAINTAINER author.email=yaphetsglhf@gmail.com

RUN mkdir -p ${WORKDIR}
WORKDIR ${WORKDIR}

RUN echo "http://mirrors.aliyun.com/alpine/v3.8/main/" > /etc/apk/repositories && \
    echo "http://mirrors.aliyun.com/alpine/v3.8/community" >>  /etc/apk/repositories

ARG GOPROXY=https://goproxy.cn
ENV GOPROXY=$GOPROXY
ENV GO111MODULE=on

COPY go.mod go.sum ${WORKDIR}/
RUN go mod download
COPY . ${WORKDIR}/

RUN go build -o server

FROM ${IMAGE}
ARG WORKDIR=/app
MAINTAINER author.email=yaphetsglhf@gmail.com

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk update && \
    apk add --no-cache tzdata  gcc build-base curl

ENV LC_ALL=en_CA.UTF-8
ENV LANG=en_CA.UTF-8
ENV LANGUAGE=en_CA.UTF-8
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN mkdir -p ${WORKDIR}
WORKDIR ${WORKDIR}

COPY --from=builder ${WORKDIR}/server ${WORKDIR}/.
CMD ["/server"]

ARG APP_NAME=ddd
ENV PORT=8080
ENV GRPC_PORT=9090

EXPOSE $PORT
EXPOSE $GRPC_PORT
