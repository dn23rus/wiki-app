#!/bin/sh

cp configs/main.json.dist configs/main.json

CONF_FILE=configs/main.json

APP_LISTEN_ADDRESS="${APP_LISTEN_ADDRESS:-0.0.0.0}:${APP_LISTEN_PORT:-8001}";
DB_HOST="${DB_HOST:-app-db}";
DB_NAME="${DB_NAME:-wiki}";
DB_USER="${DB_USER:-user}";
DB_PASSWORD="${DB_PASSWORD:-password}";

sed -i 's#${APP_LISTEN_ADDRESS}#'${APP_LISTEN_ADDRESS}'#g' "$CONF_FILE";
sed -i 's#${DB_HOST}#'${DB_HOST}'#g' "$CONF_FILE";
sed -i 's#${DB_NAME}#'${DB_NAME}'#g' "$CONF_FILE";
sed -i 's#${DB_USER}#'${DB_USER}'#g' "$CONF_FILE";
sed -i 's#${DB_PASSWORD}#'${DB_PASSWORD}'#g' "$CONF_FILE";

exec "$@"
