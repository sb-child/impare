# impare (NOT ~~impact~~ ~~淫趴~~)

(WIP)

A file verification and repair tool, written in Rust and Go

## How to build?

Please install the latest `rust` and `go >= 1.22` first.

### Linux

Install C/C++ toolchain first.

```shell
cargo b -r
mv target/release/impare impare
```

### Macos

Install C/C++ toolchain first.

```shell
cargo b -r
mv target/release/impare impare
```

### Windows

Install `MinGW`:

- Download <https://github.com/niXman/mingw-builds-binaries/releases/download/13.2.0-rt_v11-rev1/x86_64-13.2.0-release-win32-seh-msvcrt-rt_v11-rev1.7z>
- Decompress to a directory
- Add `X:\path\to\directory\mingw64\bin` to your `Path` variable

Add `x86_64-pc-windows-gnu` target:

- `rustup target add x86_64-pc-windows-gnu`

```shell
# git-bash
cargo b -r --target=x86_64-pc-windows-gnu
mv target/x86_64-pc-windows-gnu/release/impare.exe impare.exe
```

### Note

- Cross-compile to `*-musl` is not supported!
