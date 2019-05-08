FROM kiyoto/fluentd:0.10.56-2.1.1
RUN mkdir /etc/fluent
RUN ["apt-get", "update"]
RUN ["apt-get", "install", "--yes", "make", "libcurl4-gnutls-dev"]
ADD fluent.conf /etc/fluent/
RUN ["/usr/local/bin/gem", "install", "fluent-plugin-record-reformer", "--no-rdoc", "--no-ri"]
RUN ["/usr/local/bin/gem", "install", "fluent-plugin-elasticsearch", "--no-rdoc", "--no-ri"]
ENTRYPOINT ["/usr/local/bin/fluentd", "-c", "/etc/fluent/fluent.conf"]
