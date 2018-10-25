FROM alpine:latest

RUN apk upgrade && apk add bash
SHELL ["/bin/bash", "-c"]

RUN set -o pipefail && \
    apk add TODO && \
    TODO

CMD ["/bin/bash"]
