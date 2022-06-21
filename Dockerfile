FROM alpine:latest

# Install bash b/c it's better
RUN apk --no-cache --update add bash && \
    rm -fr /tmp/* /var/{cache/apk,tmp}/*
SHELL ["/bin/bash", "-c"]

# Add scripts
ADD dockerentry TODO /

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

# Set entrypoint
WORKDIR /TODO
ENTRYPOINT ["/dockerentry"]
