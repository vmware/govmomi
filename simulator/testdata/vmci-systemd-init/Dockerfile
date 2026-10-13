# vmci-systemd-init — minimal debian+systemd image for govmomi VMCI tests.
#
# Purpose: TestVMCI_SystemdInit_GuestInfoRoundTrip runs this image with
# RUN.vmci=true and RUN.privileged=true to verify that a systemd-managed
# one-shot service can write a guestinfo key via the GuestRPC Unix socket
# (Component A, $VMX_RPC_SOCK=/run/vmware/rpc.sock).
#
# Design choices:
#   - Debian bookworm (slim) base: systemd packages are well-tested there.
#   - Only packages required to boot systemd cleanly are installed; no
#     containerd, kubelet, or other heavy services.
#   - CMD ["/sbin/init"] runs systemd as PID 1. The STOPSIGNAL matches
#     systemd's clean shutdown signal.
#   - /usr/bin/vmware-rpctool is NOT included: vcsim injects it at container
#     start when RUN.vmci=true is set (govmomi/simulator/container.go
#     toolbox auto-injection).
#
# Build:
#   docker build -t docker.io/library/vmci-systemd-init:latest \
#     ./simulator/testdata/vmci-systemd-init/
#
# The image is intentionally small (~120 MB); do not add k8s-related packages.

FROM debian:bookworm-slim

# Install systemd and the bare minimum for a clean boot.
# systemd-sysv provides /sbin/init → /lib/systemd/systemd symlink.
# dbus is required for systemd's internal IPC.
RUN apt-get update && apt-get install -y --no-install-recommends \
        systemd \
        systemd-sysv \
        dbus \
        procps \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Mask units that are not meaningful or harmful in a container environment.
# systemd-networkd-wait-online and similar units would cause boot delays.
RUN systemctl mask \
        getty@tty1.service \
        getty.target \
        console-getty.service \
        systemd-networkd-wait-online.service \
        systemd-resolved.service \
        apt-daily.service \
        apt-daily-upgrade.service \
        apt-daily.timer \
        apt-daily-upgrade.timer \
        e2scrub_reap.service \
    && true

# Bake the test service into the image so it is always available without
# requiring a bind-mount injection at test time.  The govmomi vmci simulation
# injects /usr/bin/vmware-rpctool as a bind mount when RUN.vmci=true, so the
# binary does not need to be present in the image itself.
COPY vmci-test-writer.service /etc/systemd/system/vmci-test-writer.service
RUN systemctl enable vmci-test-writer.service

# /run is NOT declared as a tmpfs in the OCI spec; the host bind-mount for
# /run/vmware/rpc.sock (Component A) must survive container start.
# Declaring /run as a VOLUME would cause Docker to initialise it as a tmpfs
# and shadow the bind mount — do NOT add it here.

STOPSIGNAL SIGRTMIN+3
CMD ["/sbin/init"]
