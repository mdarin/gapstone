FROM ruby:3.4-slim-bookworm

# ARG BINDINGS=capstone/bindings/python/capstone
# ARG SPECS=capstone/tests

WORKDIR /build

# COPY genconst genconst
# COPY genspec genspec

# RUN set -eux; \
#     \
#     apk add --no-cache \
#     bash
#     dos2unix \
#     curl \
#     gpg \
#     gpg-agent
RUN apt-get update && apt-get upgrade -y && apt-get install -y \
    file \
    dos2unix \
    curl \
    gpg \
    gpg-agent


# Install Go
WORKDIR /tmp

RUN curl https://go.dev/dl/go1.23.8.linux-arm64.tar.gz -Lo ./go.linux-arm64.tar.gz \
    && curl https://go.dev/dl/go1.23.8.linux-arm64.tar.gz.asc -Lo ./go.linux-arm64.tar.gz.asc \
    && curl https://dl.google.com/dl/linux/linux_signing_key.pub  -Lo linux_signing_key.pub \
    && gpg --import linux_signing_key.pub && gpg --verify ./go.linux-arm64.tar.gz.asc ./go.linux-arm64.tar.gz \
    && tar -C /usr/local -xzf go.linux-arm64.tar.gz 
ENV PATH="$PATH:/usr/local/go/bin"
RUN go version

# WORKDIR /build
# # make sure thet we have unix line ending
# RUN dos2unix ./gen*; \
#     chmod +x ./gen*

# generating files with constants
# RUN ./genconst capstone/bindings/python/capstone; ./genspec capstone/tests

CMD [ "echo", "See usage at README" ]
