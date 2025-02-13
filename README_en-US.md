[中文简体](README.md) | English

[![Build Status](https://github.com/axetroy/FileMass/workflows/ci/badge.svg)](https://github.com/axetroy/FileMass/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/axetroy/FileMass)](https://goreportcard.com/report/github.com/axetroy/FileMass)
![Latest Version](https://img.shields.io/github/v/release/axetroy/FileMass.svg)
[![996.icu](https://img.shields.io/badge/link-996.icu-red.svg)](https://996.icu)
![Repo Size](https://img.shields.io/github/repo-size/axetroy/FileMass.svg)

## FileMass

FileMass is a powerful file generation tool used to create directory structures and various types of files (such as txt, png, markdown, doc, etc.) at any level. It supports custom file sizes, file types, and directory hierarchies, making it ideal for testing, mock data generation, or file system stress testing.

### Features

- **Multiple File Types**: Generate various file types such as txt, png, markdown, doc, etc.
- **Custom File Sizes**: Specify the minimum and maximum file sizes (in KB).
- **Arbitrary Directory Hierarchies**: Generate directory structures at any level to meet complex file system requirements.
- **Concurrency Control**: Limit the number of concurrent tasks using semaphores to avoid resource exhaustion.
- **Progress Tracking**: Display real-time file generation progress for easy task monitoring.

### Installation

1. Shell (Mac/Linux)

```bash
curl -fsSL https://github.com/release-lab/install/raw/v1/install.sh | bash -s -- -r=axetroy/FileMass
```

2. PowerShell (Windows):

```bash
$r="axetroy/FileMass";iwr https://github.com/release-lab/install/raw/v1/install.ps1 -useb | iex
```

3. [Github release page](https://github.com/axetroy/FileMass/releases) (All platforms supported)

> Download the executable file and add it to the `$PATH` environment variable.

4. Build and install from source using [Golang](https://golang.org) (All platforms supported)

```bash
go install github.com/axetroy/FileMass/cmd/FileMass
```

5. Install via npm

```sh
npm install @axetroy/FileMass -g
```

### Usage

```bash
$ filemass --help
FileMass - a powerful file generation tool used to create directory structures and various types of files

USAGE:
  filemass [OPTIONS]

OPTIONS:
  --help                Print help information
  --version             Print version information
  --concurrence         Set the number of concurrent tasks, default is 1
  --output              Set the output directory
  --depth               Set the depth of the directory, default is 1
  --min-size            Set the minimum size(in KB) of the file, default is 1KB
  --max-size            Set the maximum size(in KB) of the file, default is 1024KB
  --dirs                Set the number of directories, default is 1
  --files               Set the number of files in the directory, default is 1
  --clean               Clean the output directory before generating files

SOURCE CODE:
  https://github.com/axetroy/FileMass
```

### Open Source License

The [Anti-996 License](LICENSE)
