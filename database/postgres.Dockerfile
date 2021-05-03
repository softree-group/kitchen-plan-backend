FROM postgres:alpine

ADD database/init/backup.sql /
ADD database/init/init.sh /docker-entrypoint-initdb.d/
RUN chmod +x /docker-entrypoint-initdb.d/init.sh

EXPOSE 5432
