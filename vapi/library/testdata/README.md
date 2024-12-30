Please note, the only purpose of the `.ovf`, `.vmdk` and `.iso` files in this directory is to instrument the content library simulator. The files themselves are not valid. They are the product of running the following command to produce `1MiB` files:

```shell
dd if=/dev/zero of=ttylinux-pc_i486-16.1-disk1.vmdk bs=1024 count=0 seek=1024
dd if=/dev/zero of=ttylinux-pc_i486-16.1.iso bs=1024 count=0 seek=1024
```

The `.mf` manifest was created with:

```shell
sha1sum --tag *.ovf *.vmdk >ttylinux-pc_i486-16.1.mf
```

The `.ova` file was constructed with:

```shell
tar cf ttylinux-pc_i486-16.1.ova *.ovf *.mf *.vmdk
```
