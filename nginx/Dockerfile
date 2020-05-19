FROM ubuntu:16.04

RUN apt-get update && apt install -y curl gnupg2 ca-certificates lsb-release apt-utils gettext-base
RUN echo "deb http://nginx.org/packages/ubuntu `lsb_release -cs` nginx" | tee /etc/apt/sources.list.d/nginx.list
RUN curl -fsSL https://nginx.org/keys/nginx_signing.key | apt-key add -
RUN apt-key fingerprint ABF5BD827BD9BF62
RUN apt-get update && apt-get install -y nginx

COPY ./default.conf /etc/nginx/conf.d/default.template

RUN envsubst < /etc/nginx/conf.d/default.template > /etc/nginx/conf.d/default.conf 

EXPOSE 80/tcp

WORKDIR /usr/share/nginx/html/
CMD ["/bin/sh", "-c", "exec nginx -g 'daemon off;';"]

