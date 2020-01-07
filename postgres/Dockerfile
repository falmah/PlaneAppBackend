FROM ubuntu:16.04

RUN apt-get update && apt-get install -y wget
RUN wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add -
RUN echo "deb http://apt.postgresql.org/pub/repos/apt/ xenial-pgdg main" | tee  /etc/apt/sources.list.d/pgdg.list

RUN apt-get update && apt-get install -y --allow-unauthenticated postgresql-12 postgresql-client-12 postgresql-contrib-12

USER postgres

COPY . .

RUN /etc/init.d/postgresql start  && \
    psql --command "CREATE USER docker WITH SUPERUSER PASSWORD 'docker';" && \
    createdb -O docker test_db && \
    psql -d test_db --file /scripts/types.sql && \
    psql -d test_db --file /scripts/tables.sql && \
    psql -d test_db --file /scripts/csvImport.sql

EXPOSE 5432

VOLUME  ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]

CMD ["/usr/lib/postgresql/12/bin/postgres", "-D", "/var/lib/postgresql/12/main", "-c", "config_file=/etc/postgresql/12/main/postgresql.conf"]
