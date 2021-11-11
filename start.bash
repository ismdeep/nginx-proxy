#!/usr/bin/env bash
set -e
CONFIG_FILE=/config.yaml UPSTREAM_OUTPUT_FILE=/etc/nginx/tcp.d/default.conf /usr/bin/nginx-proxy-generator
nginx -g "daemon off;"
