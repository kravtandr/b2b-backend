#!/bin/bash

set -e

host="postgres"
port="5432"
cmd="$@"

>&2 echo "!!!!!!!! Check postrgres for available !!!!!!!!"

until curl http://"$host":"$port"; do
  >&2 echo "postrgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "postrgres is up - executing command"

exec $cmd