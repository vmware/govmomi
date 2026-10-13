# © Broadcom. All Rights Reserved.
# The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
# SPDX-License-Identifier: Apache-2.0

ARG GO_VERSION=1.25.0
FROM golang:${GO_VERSION} AS build

WORKDIR /src
COPY . .
WORKDIR /src/vcsim
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/vcsim . \
    && mkdir -p /out/tmp \
    && chmod 1777 /out/tmp

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build --chown=10001:10001 /out/tmp /tmp
COPY --from=build --chown=10001:10001 /out/vcsim /vcsim

USER 10001:10001
EXPOSE 8989
ENTRYPOINT ["/vcsim"]
CMD ["-l", "0.0.0.0:8989"]
