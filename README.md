# impare (NOT ~~impact~~ ~~淫趴~~)

(WIP)

A file verification and repair tool, written in Rust and Go

## Build

Please install the latest `rust` and `go >= 1.22` first.

### Linux

```shell
cargo b -r
mv impare/target/release/impare impare
```

### Macos

```shell
cargo b -r
mv impare/target/release/impare impare
```

### Windows

```shell
CGO_ENABLED=1 cargo b -r
mv impare/target/release/impare.exe impare.exe
```

### Note

- Cross-compile to `x86_64-unknown-linux-musl` is not supported!
