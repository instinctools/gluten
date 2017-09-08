FROM ubuntu:16.04

WORKDIR opt/gluten-slave

ADD ./slave /opt/gluten-slave/slave
ADD migrations /opt/gluten-slave/migrations
ADD ./slave-config.yml /opt/gluten-slave/slave-config.yml

EXPOSE 7000

ENTRYPOINT ["/opt/gluten-slave/slave"]