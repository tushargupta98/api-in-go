#!/bin/bash
# WARNING: Don't excute this script explicitly. Use corresponding `make` target to execute this script. 

###----- Create New Schema: Start ------###
set -euo pipefail

export PGPASSWORD="${PG_PASSWORD}"

echo "Adding schema ${PG_SCHEMA} to the database..."

psql -h ${PG_URL} -p ${PG_PORT} -U ${PG_USERNAME} -d ${PG_DATABASE} -c "CREATE SCHEMA ${PG_SCHEMA};"

echo "Schema added successfully!"
###----- Create New Schema: End ------###