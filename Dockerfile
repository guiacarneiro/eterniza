FROM ubuntu

RUN mkdir -p /home/eterniza && \
  mkdir -p /home/eterniza/timezone && \
  mkdir -p /home/eterniza/config && \
  mkdir -p /home/eterniza/log

COPY ./out/eterniza /home/eterniza/
COPY ./arquivos/start.sh /home/

RUN ln -nsf /home/eterniza/timezone/localtime /etc/ && \
  chmod 777 /home/start.sh

VOLUME /home/eterniza/config
VOLUME /home/eterniza/log
VOLUME /home/eterniza/timezone
VOLUME /home/eterniza/arquivos

COPY ./arquivos/Sao_Paulo /home/eterniza/timezone/localtime

EXPOSE 3000

ENTRYPOINT [ "/bin/sh", "/home/start.sh"]

