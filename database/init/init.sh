#!/bin/bash

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
		CREATE USER mark_bykhovets WITH PASSWORD 'softree-group';
		CREATE DATABASE receipts;
		GRANT ALL PRIVILEGES ON DATABASE receipts TO mark_bykhovets;
EOSQL

pg_restore -U mark_bykhovets -d receipts -1 backup.sql
