# tldf

[![Go](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go&logoColor=white)](https://go.dev)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Platform](https://img.shields.io/badge/Platform-macOS%20%7C%20Linux-lightgrey)]()

Fuzzy finder for [tldr](https://github.com/tldr-pages/tldr) pages. Pipes `tldr --list` into `fzf` with a live `bat`-rendered preview, prints the selected page to stdout.

One binary. No config. No flags.

## Install

```sh
go install github.com/dutch-casa/tldrfzf@latest
```

Or build from source:

```sh
git clone https://github.com/dutch-casa/tldrfzf.git
cd tldrfzf
go build -o tldf .
```

## Dependencies

These must be on your `$PATH`:

- [tldr](https://github.com/tldr-pages/tldr) -- the page database
- [fzf](https://github.com/junegunn/fzf) -- fuzzy selection
- [bat](https://github.com/sharkdp/bat) -- syntax-highlighted preview

## Usage

```sh
tldf
```

That's it. You get a fuzzy search over every tldr page. Arrow keys to browse, enter to select, `esc` or `ctrl-c` to quit. The selected page prints directly to your terminal.

## How it works

Three shell commands composed into a pipeline:

1. `tldr --list` produces the full page index
2. `fzf` presents it with a live preview (`tldr <page> | bat`)
3. `tldr <selected>` prints the final result

No temp files, no caching, no background processes. The binary is a thin coordinator. If any dependency is missing, it fails immediately with a clear error.

## Design

<!-- Invariant: all three binaries (tldr, fzf, bat) must exist on $PATH.
   Enforced at first use -- exec.Command fails fast on missing binary.
   No fallback behavior. Absence is an error, not a degraded mode. -->

<!-- fzf exit codes 1 (no match) and 130 (user cancel) are not errors.
   They represent intentional user decisions. Treating them as failures
   would break the contract: the user chose to quit. -->

The program has exactly two outcomes: print a tldr page, or exit cleanly. There is no partial state. Each external command either succeeds fully or the process terminates with a diagnostic on stderr.

User cancellation (`esc`, `ctrl-c`, no match) exits 0. That's not a failure -- it's the user saying "never mind."

## License

MIT
