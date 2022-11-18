# Yaag
> 上传 Sw**ag**ger 到 **Ya**pi.

[English](README.md) | 简体中文

## 安装

### go install

```bash
go install github.com/leslieleung/yaag@latest
```

## 快速开始

### 在项目目录下创建一个 `yaag.yaml`

可以将 `example/yaag.yaml` 复制到项目根目录，或者使用 `yaag -init` 来创建一个。

### 填写 `yaag.yaml`

```yaml
docDir: "" # swagger.json 生成目录
mergeMode: "" # Yapi合并模式， 见 http://yapi.smart-xwork.cn/doc/openapi.html
swagGeneral: "" # API 基础信息注解的目录, 默认为 main.go
yapiUrl: "" # Yapi 的 url
yapiToken: "" # Yapi 项目的 token
```

### 安装 swag 并写注解

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

更多详情请参考 [swaggo/swag](https://github.com/swaggo/swag) 。

### 运行 `yaag`

## 使用方式

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