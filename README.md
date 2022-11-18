# Yaag
> Upload Sw**ag**ger to **Ya**pi.

English | [简体中文](README_zh.md)

## Installation

### go install

```bash
go install github.com/leslieleung/yaag@latest
```

## Quick Start

### Create a `yaag.yaml` under your project

You can either copy `example/yaag.yaml` to your project root or 
simply use `yaag -init` to create one.

### Fill in the `yaag.yaml`

```yaml
docDir: "" # the directory where you want to store the swagger.json
mergeMode: "" # the mergeMode of Yapi, see http://yapi.smart-xwork.cn/doc/openapi.html
swagGeneral: "" # the general API annotations, default to main.go
yapiUrl: "" # url to yapi
yapiToken: "" # token to yapi
```

### Install swag and write annotations

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

For more details, see [swaggo/swag](https://github.com/swaggo/swag) .

### Run `yaag` in your project

## Usage

```bash
Usage of yaag:
  -docDir string
        path to swagger.json
  -init
        init
  -mergeMode string
        merge mode
  -upload
        upload only, no swag generate
```