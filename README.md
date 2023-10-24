# includejni

## 安装

```bash
go install github.com/ClarkGuan/includejni@latest
```

## 使用

```bash
includejni go build
```

自动追加 JNI 头文件路径到环境变量 CGO_CFLAGS 以及 CGO_CXXFLAGS 中。
