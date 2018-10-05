# First iteration, proof of concept image for bowline activate usage

FROM alpine

RUN echo "#!/usr/bin/env sh" > /usr/bin/bowline && \
    echo "echo \"export BOWLINE_ACTIVATED=true\"" >> /usr/bin/bowline && \
    chmod +x /usr/bin/bowline

CMD ["/usr/bin/bowline"]
