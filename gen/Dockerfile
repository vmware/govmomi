## The version of Go from which this image is based. It must satisfy the go
## directive in ../go.mod. It also builds goimports below: goimports built
## with Go 1.27 indents lists in doc comments differently, which changes the
## comments of every generated file.
ARG GO_VERSION=1.25.14

## The digest pins the base image, so that the image does not change when the
## tag moves. Update it together with GO_VERSION.
ARG GO_IMAGE_DIGEST=sha256:2c4c60ef415fbfa5e90300722293bef36c5e63fae17570ce18f580af933dbd73

## Docker image used as base of this image. Debian trixie provides Ruby 3.3,
## which the nokogiri version in Gemfile.lock requires.
FROM --platform=${BUILDPLATFORM} golang:${GO_VERSION}-trixie@${GO_IMAGE_DIGEST}


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
## Install Ruby, Bundler and the build
## tools needed by native gems
## --------------------------------------

## Debian packages come from a fixed snapshot of the archive, and the versions
## are pinned, so that the image does not change when Debian publishes updates.
## Update the snapshot and the versions together.
ARG DEBIAN_SNAPSHOT=20261105T000000Z
RUN rm -f /etc/apt/sources.list.d/debian.sources && \
    for s in "debian trixie" "debian trixie-updates" "debian-security trixie-security"; do \
      set -- $s; \
      echo "deb [check-valid-until=no] http://snapshot.debian.org/archive/$1/${DEBIAN_SNAPSHOT} $2 main"; \
    done > /etc/apt/sources.list && \
    apt-get update && \
    apt-get install -y --no-install-recommends \
      build-essential=12.12 \
      curl=8.14.1-2+deb13u7 \
      pkg-config=1.8.1-4 \
      ruby=1:3.3+b1 \
      ruby-dev=1:3.3+b1 \
      ruby3.3=3.3.8-2 \
      ruby3.3-dev=3.3.8-2 \
      zlib1g-dev=1:1.3.dfsg+really1.3.1-1+b1 && \
    rm -rf /var/lib/apt/lists/*

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
## Install goimports to format gen'd files
## --------------------------------------

## A fixed version keeps the formatting of generated files reproducible.
ARG GOIMPORTS_VERSION=v0.42.0
RUN go install golang.org/x/tools/cmd/goimports@${GOIMPORTS_VERSION}
