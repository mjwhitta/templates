FROM alpine:latest

# Install bash b/c it's better
RUN apk update && apk add bash
SHELL ["/bin/bash", "-c"]

# Add scripts
ADD dockerentry TODO /

# 1. Install dependencies
# 2. Clean up unnecessary files and packages
RUN set -o pipefail && \
    ( \
        apk upgrade && \
        apk add \
            curl \
            grep \
            less \
            shadow \
            sudo \
            TODO \
    ) && ( \
        rm -rf /tmp/* /var/{cache/apk,tmp}/* \
    )

# Set entrypoint
WORKDIR /TODO
ENTRYPOINT ["/dockerentry"]
