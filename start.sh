#!/bin/sh

set -e 
echo "Start migration..." 
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up 
echo "Migration completed" 
exec "$@"