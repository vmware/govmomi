## The version of Go from which this image is based.
ARG GO_VERSION=1.23.2

## Docker image used as base of this image.
FROM --platform=${BUILDPLATFORM} golang:${GO_VERSION}


## --------------------------------------
## Multi-platform support
## --------------------------------------

ARG TARGETOS
ARG TARGETARCH


## --------------------------------------
## Environment variables
## --------------------------------------

ENV GOOS=${TARGETOS}
ENV GOARCH=${TARGETARCH}


## --------------------------------------
## Update the apt cache & essentials
## --------------------------------------
RUN apt-get update && \
    apt-get install -y build-essential curl


## --------------------------------------
## Install the version of openssl
## required by ruby 2.x, which is
## required to generate the types
## --------------------------------------
RUN mkdir -p /opt/src /opt/lib && \
    curl -sSL https://www.openssl.org/source/openssl-1.1.1g.tar.gz | \
    tar -C /opt/src -xz && \
    cd /opt/src/openssl-1.1.1g && \
    ./config --prefix=/opt/lib/openssl-1.1.1g \
             --openssldir=/opt/lib/openssl-1.1.1g && \
    make && \
    make install && \
    rm -fr /opt/lib/openssl-1.1.1g/certs && \
    ln -s /etc/ssl/certs /opt/lib/openssl-1.1.1g/certs


## --------------------------------------
## Install Ruby & Bundler
## --------------------------------------

ENV PATH="/root/.rbenv/shims:${PATH}" \
    RUBY_CONFIGURE_OPTS="--with-openssl-dir=/opt/lib/openssl-1.1.1g"

RUN apt-get install -y rbenv && \
    rbenv install 3.1.0 && \
    rbenv rehash && \
    rbenv global 3.1.0

RUN gem install bundler -v 2.6.6


## --------------------------------------
## Configure the working directory
## --------------------------------------

WORKDIR /govmomi/gen


## --------------------------------------
## Cache the gen program dependencies
## --------------------------------------

COPY Gemfile Gemfile.lock .
RUN bundle install


## --------------------------------------
## Install goimports to forma gen'd files
## --------------------------------------

RUN go install golang.org/x/tools/cmd/goimports@latest
