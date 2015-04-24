FROM scratch

ADD counter-linux /counter
ENV PORT 9090
EXPOSE 9090

ENTRYPOINT ["/counter"]

