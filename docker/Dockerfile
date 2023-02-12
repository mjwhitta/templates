FROM alpine:latest

# Use bash b/c it's better
RUN apk --no-cache --update add bash && \
    rm -fr /tmp/* /var/{cache/apk,tmp}/*
SHELL ["/bin/bash", "-c"]

# 1. Install dependencies
# 2. Clean up unnecessary files and packages
RUN set -o pipefail && \
    ( \
        apk --no-cache --update upgrade && \
        apk --no-cache --update add \
            curl \
            grep \
            less \
            shadow \
            sudo \
            TODO \
    ) && ( \
        rm -fr /tmp/* /var/{cache/apk,tmp}/* \
    )

WORKDIR /TODO

# Add scripts
ADD root/dockerentry root/TODO /
RUN chmod 755 /dockerentry /TODO

# Set entrypoint
ENTRYPOINT ["/dockerentry"]
