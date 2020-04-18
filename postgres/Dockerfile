FROM ubuntu:16.04

#default docker0 adapter address
ARG DOCKER_HOST=172.17.0.1

RUN apt-get update && apt-get install -y wget
RUN wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add -
RUN echo "deb http://apt.postgresql.org/pub/repos/apt/ xenial-pgdg main" | tee  /etc/apt/sources.list.d/pgdg.list

RUN apt-get update && apt-get install -y --allow-unauthenticated postgresql-12 postgresql-client-12 postgresql-contrib-12

USER postgres

COPY . .

RUN /etc/init.d/postgresql start  && \
    psql --command "CREATE USER docker WITH SUPERUSER PASSWORD 'docker';" && \
    createdb -O docker app_db && \
    psql -d app_db --file /scripts/types.sql && \
    psql -d app_db --file /scripts/tables.sql && \
    psql -d app_db --file /scripts/csvImport.sql

RUN echo "host  all all 0.0.0.0/0   md5" >> /etc/postgresql/12/main/pg_hba.conf
RUN echo "hostnossl  all all  0.0.0.0/0   md5" >> /etc/postgresql/12/main/pg_hba.conf
RUN echo "host  all docker  ${DOCKER_HOST}/32   md5" >> /etc/postgresql/12/main/pg_hba.conf
RUN echo "hostnossl  all docker  ${DOCKER_HOST}/32   md5" >> /etc/postgresql/12/main/pg_hba.conf
RUN echo "listen_addresses='*'" >> /etc/postgresql/12/main/postgresql.conf

EXPOSE 5432

VOLUME  ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]

CMD ["/usr/lib/postgresql/12/bin/postgres", "-D", "/var/lib/postgresql/12/main", "-c", "config_file=/etc/postgresql/12/main/postgresql.conf"]
