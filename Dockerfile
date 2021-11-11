FROM hub.deepin.com/library/golang:bullseye AS builder
WORKDIR /src
COPY . .
RUN go build -mod vendor -o main github.com/ismdeep/nginx-proxy

FROM hub.deepin.com/library/nginx:latest
RUN mkdir -p /etc/nginx/tcp.d
ENV CONFIG_FILE /config.yml
ENV UPSTREAM_OUTPUT_FILE /etc/nginx/tcp.d/default.conf
COPY --from=builder /src/main /main
COPY ./nginx.conf /etc/nginx/nginx.conf
RUN echo '#!/usr/bin/env bash\n/main && nginx -g "daemon off;"' > /start.bash
CMD ["bash", "/start.bash"]