#!/bin/sh
# Wait for Spicedb to be ready
while ! nc -z spicedb 50051; do sleep 1; done
# Load the schema using the zed CLI or other appropriate method
zed schema write -e "spicedb:50051" -f /app/schema.relations
