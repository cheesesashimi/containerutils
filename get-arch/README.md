# get-arch

This utility is mostly useful for multi-arch container builds. What it does is it maps the current system architecture to a vanity name of the users' choosing.

## Problem Statement

Sometimes, it is necessary to download a pre-built binary from GitHub Releases
or another location. These binaries are often built for a specific operating
system / architecture combination such as linux/amd64, linux/arm64, or
darmin/arm64. Unfortunately, there is no standardization around whether one
should use x86-64 or aarch64 in their path names. This makes it very difficult
to rely upon `$(uname -m)` to correctly identify the underlying host's
architecture.

## Solution

`get-arch` solves this by allowing one to specify the correct architecture
vanity name for each architecture one is attempting to download a release for:

```dockerfile
RUN curl -Lo ripgrep "https://github.com/BurntSushi/ripgrep/releases/download/14.1.0/ripgrep-14.1.0-$(get-arch --x86_64 --arch64)-unknown-linux-musl.tar.gz"
```

In this example, if the container is being built on an amd64 host, it will output `x86_64`. If built on an arm64 host, it will output `aarch64`.

## Usage

get-arch includes defaults for `amd64`, `x86-64`, `x86_64`, `aarch64` and
`arm64`. It expects _exactly_ one flag for each architecture. Multiple flags
for an architecture or a missing flag will result in an error.

If none of those options applies to your particular scenario, you can
provide a custom value by using the `--custom-amd64 "<custom-amd64>"` and
`--custom-arm64 "<custom-arm64>"` flags, respectively.

```console
./get-arch --help
NAME:
   get-arch - Select the correct arch for your download.

USAGE:
   get-arch [global options] command [command options]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help

   Options for amd64:

   --amd64               use to select amd64 (default: false)
   --custom-amd64 value  use to provide a custom architecture for amd64
   --x86-64              use to select x86-64 (default: false)
   --x86_64              use to select x86_64 (default: false)

   Options for arm64:

   --aarch64             use to select aarch64 (default: false)
   --arm64               use to select arm64 (default: false)
   --custom-arm64 value  use to provide a custom architecture for arm64
```
