FROM ubuntu:16.04

WORKDIR /opt/gluten-master

ADD ./backend /opt/gluten-master/master
ADD migrations /opt/gluten-master/migrations
ADD ./master-config.yml /opt/gluten-master/master-config.yml

EXPOSE 8888 8889

ENTRYPOINT ["/opt/gluten-master/master", "-r", "8888", "-w", "8889"]