中文简体 | [English](README_en-US.md)

[![Build Status](https://github.com/axetroy/FileMass/workflows/ci/badge.svg)](https://github.com/axetroy/FileMass/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/axetroy/FileMass)](https://goreportcard.com/report/github.com/axetroy/FileMass)
![Latest Version](https://img.shields.io/github/v/release/axetroy/FileMass.svg)
[![996.icu](https://img.shields.io/badge/link-996.icu-red.svg)](https://996.icu)
![Repo Size](https://img.shields.io/github/repo-size/axetroy/FileMass.svg)

## FileMass

FileMass 是一个强大的文件生成工具，用于生成任意层级的目录结构和多种类型的文件（如 txt、png、markdown、doc 等）。它支持自定义文件大小、文件类型和目录层次，非常适合用于测试、模拟数据生成或文件系统压力测试。

### 功能特性

- **多文件类型支持**：生成 txt、png、markdown、doc 等多种文件类型。
- **自定义文件大小**：支持指定文件的最小和最大大小（以 KB 为单位）。
- **任意目录层次**：生成任意层级的目录结构，满足复杂文件系统的需求。
- **并发控制**：通过信号量限制并发任务数量，避免资源耗尽。
- **进度跟踪**：实时显示文件生成进度，方便监控任务执行情况。

### 安装

1. Shell (Mac/Linux)

```bash
curl -fsSL https://github.com/release-lab/install/raw/v1/install.sh | bash -s -- -r=axetroy/FileMass
```

2. PowerShell (Windows):

```bash
$r="axetroy/FileMass";iwr https://github.com/release-lab/install/raw/v1/install.ps1 -useb | iex
```

3. [Github release page](https://github.com/axetroy/FileMass/releases) (全平台支持)

> 下载可执行文件，并且把它加入到`$PATH` 环境变量中

4. 使用 [Golang](https://golang.org) 从源码中构建并安装 (全平台支持)

```bash
go install github.com/axetroy/FileMass/cmd/FileMass
```

5. 通过 npm 安装

```sh
npm install @axetroy/FileMass -g
```

### 开源许可

The [Anti-996 License](LICENSE_zh-CN)
