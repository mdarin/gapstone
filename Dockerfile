FROM ruby:3.3-slim-bookworm

# ARG BINDINGS=capstone/bindings/python/capstone
# ARG SPECS=capstone/tests

WORKDIR /build

COPY genconst genconst
COPY genspec genspec

# RUN set -eux; \
#     \
#     apk add --no-cache \
#     dos2unix \
#     curl \
#     gpg \
#     gpg-agent
RUN apt-get update && apt-get upgrade -y && apt-get install -y \
    dos2unix \
    curl \
    gpg \
    gpg-agent


# Install Go
WORKDIR /tmp

RUN curl https://go.dev/dl/go1.23.8.linux-amd64.tar.gz -Lo ./go.linux-amd64.tar.gz \
    && curl https://go.dev/dl/go1.23.8.linux-amd64.tar.gz.asc -Lo ./go.linux-amd64.tar.gz.asc \
    && curl https://dl.google.com/dl/linux/linux_signing_key.pub  -Lo linux_signing_key.pub \
    && gpg --import linux_signing_key.pub && gpg --verify ./go.linux-amd64.tar.gz.asc ./go.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go.linux-amd64.tar.gz 
ENV PATH="$PATH:/usr/local/go/bin"
RUN go version

WORKDIR /build
# make sure thet we have unix line ending
RUN dos2unix ./gen*; \
    chmod +x ./gen*

# generating files with constants
# RUN ./genconst capstone/bindings/python/capstone; ./genspec capstone/tests

CMD [ "ls" ]
