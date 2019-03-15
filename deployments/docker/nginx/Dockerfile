FROM nginx:alpine

LABEL Description="nginx proxy"

RUN apk update

ENV WEB_ROOT=/app/web
ENV APP_PROXY_PASS=http://wiki-app:8001

RUN rm /etc/nginx/conf.d/default.conf
ADD default.conf /etc/nginx/conf.d/

RUN set -x ; \
  addgroup -g 82 -S www-data ; \
  adduser -u 82 -D -S -G www-data www-data && exit 0 ; exit 1

ADD nginx.conf /etc/nginx/nginx.conf
ADD entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

WORKDIR /app

ENTRYPOINT ["/entrypoint.sh"]
CMD ["nginx", "-g", "daemon off;"]
