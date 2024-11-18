<p align="center">
   <span style="color: #2ba9f1;font-size: 24px;font-weight: bold;">Go Webview</span>
</p>

<p align="center" style="font-size: 24px;">
    <strong>
        是Go基于 LCL & Webview2, Webkit2 构建桌面应用的框架
    </strong>
</p>

---
![go-version](https://img.shields.io/github/go-mod/go-version/energye/wv?logo=git&logoColor=green)
[![github](https://img.shields.io/github/last-commit/energye/wv/main.svg?logo=github&logoColor=green&label=commit)](https://github.com/energye/wv)
[![release](https://img.shields.io/github/v/release/energye/wv?logo=git&logoColor=green)](https://github.com/energye/wv/releases)
![repo](https://img.shields.io/github/repo-size/energye/wv.svg?logo=github&logoColor=green&label=repo-size)
[![Go Report](https://goreportcard.com/badge/github.com/energye/wv)](https://goreportcard.com/report/github.com/energye/wv)
[![Go Reference](https://pkg.go.dev/badge/github.com/energye/wv)](https://pkg.go.dev/github.com/energye/wv)
[![license](https://img.shields.io/github/license/energye/wv.svg?logo=git&logoColor=red)](http://www.apache.org/licenses/LICENSE-2.0)
---

### 项目简介

> [Go Webview2(wv)](https://github.com/energye/wv) 
> 是 Go 基于
> [LCL](https://www.lazarus-ide.org/)、
> [Windows-Webview2](https://learn.microsoft.com/en-us/microsoft-edge/webview2)
> [Linux-Webkit2](https://www.webkitgtk.org)
> [MacOS-Webkit2](https://developer.apple.com/documentation/webkit)
> 开发的框架
>
>> LCL - 基础库, 图形用户界面(GUI)组件库, 提供了非常丰富的系统原生控件
>> 
>> Windows-Webview2 - 浏览器组件库 [WebView4Delphi](https://github.com/salvadordf/WebView4Delphi), 在LCL基础上封装的Webview2库
>> 
>> Linux-Webkit2 - 浏览器组件库 Webkit2GTK, 在LCL基础上封装的Webkit2库
>> 
>> MacOS-Webkit2 - 浏览器组件库 Webkit2Cocoa, 在LCL基础上封装的Webkit2库
>> 
>> 使用 Go 和 Web 端技术 ( HTML + CSS + JavaScript ) 构建支持Windows平台桌面应用。
>>
>> 将web内容无缝集成到应用程序中，并自定义内容交互以满足应用程序的需求。
> 
> 构建&使用
> 
>> LCL 单独使用, 开发原生图形用户界面(GUI)应用. 轻量级, 丰富的系统原生控件
>
>> LCL + Webview 混合使用, 开发原生图形用户界面(GUI)和浏览器应用. 轻量级, 全量webview2 API, 轻量级 Webkit2 API



### 特点

> - 依赖 `microsoft-edge` 运行时环境
> - 具有完整的 Webview2 API 轻量级 Webkit2 API 和 LCL 系统原生控件
> - 开发环境简单, 编译速度快, 仅需Go和Webview2所需的动态链接库
> - 前端技术: 支持主流前端框架。例如：Vue、React、Angular 和 原生HTML+CSS+JS等
> - 事件驱动: 高性能事件驱动, 基于IPC通信，实现Go和Web端快速调用及数据交互
> - 资源加载: 可无需http服务支撑，直接读取本地资源或内置到执行文件的资源, 也支持http服务加载资源

### 内置依赖&集成

- [![LCL](https://img.shields.io/badge/LCL-green)](https://github.com/energye/lcl)
- [![WebView4Delphi](https://img.shields.io/badge/Windows-Webview2%20-green)](https://github.com/salvadordf/WebView4Delphi)
- [![Linux-Webkit2](https://img.shields.io/badge/Linux-Webkit2%20-green)](https://www.webkitgtk.org)
- [![MacOS-Webkit2](https://img.shields.io/badge/MacOS-Webkit2%20-green)](https://developer.apple.com/documentation/webkit)

#### 基本需求

> - Golang >= 1.20
> - 动态链接库 
>> - Windows: `WebView2Loader.dll` `liblcl`
>> - MacOS Linux: `liblcl

#### [示例](https://github.com/energye/examples/tree/main/wv)

#### 开发环境

1. 安装 [Golang](https://golang.google.cn/dl/)
2. 下载 `Webview2` 和 `LCL` 控件库动态链接库, Windows 两个动态链接库必须同时使用, MacOS, Linux 只需要 liblcl
3. Windows 将两个动态链接库配置到环境变量 `ENERGY_HOME` 目录下, 或放置到和执行文件 `exe` 同一目录
4. 创建Go项目开始使用构建桌面应用, 参考示例 [Webview2 examples](https://github.com/energye/examples/tree/main/wv)

##### Windows

- LCL+Webview2

  - [Windows32](https://sourceforge.net/projects/liblcl/files/v3.0.0/lcl_wv2_binary_windows32.zip/download)
  - [Windows64](https://sourceforge.net/projects/liblcl/files/v3.0.0/lcl_wv2_binary_windows64.zip/download)

- WebView2Loader 1.0.2277.86

  - [Windows32](https://sourceforge.net/projects/liblcl/files/v3.0.0/WebView2Loader_32_1.0.2277.86.zip/download)
  - [Windows64](https://sourceforge.net/projects/liblcl/files/v3.0.0/WebView2Loader_64_1.0.2277.86.zip/download)

##### MacOS LCL + Webkit2

  - [MacOSx64](https://sourceforge.net/projects/liblcl/files/v3.0.0/lcl_wk2_binary_macosx64.zip/download)
  - [MacOSARM64](https://sourceforge.net/projects/liblcl/files/v3.0.0/lcl_wk2_binary_macosarm64.zip/download)

##### Linux LCL + Webkit2 GTK3
  - [Linux64 GTK3](https://sourceforge.net/projects/liblcl/files/v3.0.0/lcl_wk2_binary_linux64.zip/download)
  - [LinuxARM64 GTK3](https://sourceforge.net/projects/liblcl/files/v3.0.0/lcl_wk2_binary_linuxarm64.zip/download)
 
  `GTK >= 3.24.24 and Glib2.0 >= 2.66`

##### Linux LCL + Webkit1 GTK2
 - To be added

### 相关项目
* [Go LCL](https://github.com/energye/lcl)
* [Go Webview](https://github.com/energye/wv)
* [Go CEF](https://github.com/energye/cef)
* [WebView4Delphi](https://github.com/salvadordf/WebView4Delphi)
* [CEF](https://github.com/chromiumembedded/cef)
* [CEF4Delphi](https://github.com/salvadordf/CEF4Delphi)
* [CefSharp](https://github.com/cefsharp/CefSharp)
* [Java-CEF](https://bitbucket.org/chromiumembedded/java-cef)
* [cefpython](https://github.com/cztomczak/cefpython)
* [Chromium](https://chromium.googlesource.com/chromium/src/)

---

如果你觉得此项目对你有帮助，请点亮 Star

---

### ENERGY QQ交流群 & 微信

<p align="center">
    <img src="https://energye.github.io/imgs/assets/qq-group.jpg" width="250" title="QQ交流群: 541258627" alt="QQ交流群: 541258627">
    <img src="https://energye.github.io/imgs/assets/we-chat.jpg" width="250" title="微信: sniawmdf" alt="微信: sniawmdf" style="margin-left: 30px;">
</p>

---

### 鸣谢 Jetbrains

<p align="center">
    <a href="https://www.jetbrains.com?from=energy">
        <img src="https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.svg" alt="JetBrains Logo (Main) logo.">
    </a>
</p>

---

### 项目截图
##### Windows-10
<img src="https://assets.yanghy.cn/webview2-simple.png">

----

### 开源协议

[![license](https://img.shields.io/github/license/energye/wv.svg?logo=git&logoColor=green)](http://www.apache.org/licenses/LICENSE-2.0)
