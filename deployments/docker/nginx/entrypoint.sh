#!/bin/sh

CONF_FILE=/etc/nginx/conf.d/default.conf

APP_PROXY_HOST="${APP_PROXY_HOST:-wiki-app}";
APP_PROXY_PORT="${APP_PROXY_PORT:-8001}";

APP_PROXY_PASS="http://$APP_PROXY_HOST:$APP_PROXY_PORT";
WEB_ROOT="${WEB_ROOT:-/web/app}";

sed -i 's#${APP_PROXY_PASS}#'${APP_PROXY_PASS}'#g' "$CONF_FILE";
sed -i 's#${WEB_ROOT}#'${WEB_ROOT}'#g' "$CONF_FILE";

exec "$@"
