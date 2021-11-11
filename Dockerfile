FROM hub.deepin.com/library/golang:bullseye AS builder
ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /src
COPY . .
RUN set -eux; \
    go mod tidy; \
    go build -o main github.com/ismdeep/nginx-proxy

FROM hub.deepin.com/library/nginx:latest
RUN mkdir -p /etc/nginx/tcp.d
COPY --from=builder /src/main /usr/bin/nginx-proxy-generator
COPY ./nginx.conf /etc/nginx/nginx.conf
COPY ./start.bash /start.bash
CMD ["bash", "/start.bash"]
