# deblocus

[![Build Status](https://travis-ci.org/Lafeng/deblocus.svg?branch=master)](https://travis-ci.org/Lafeng/deblocus)
[![MIT License](https://img.shields.io/packagist/l/doctrine/orm.svg)](http://opensource.org/licenses/MIT)
[![Issues](https://img.shields.io/github/issues/Lafeng/deblocus.svg)](https://github.com/Lafeng/deblocus/issues)
[![Comment-on-Github.Party](https://img.shields.io/badge/Comment%20on-Github.Party-yellow.svg)](https://github.party/item?id=46)

Current version is v0.9.2500-beta

```
      ___         ___         ___         ___  ___         ___         ___         ___     
     /\  \       /\  \       /\  \       /\__\/\  \       /\  \       /\__\       /\  \    
    /::\  \     /::\  \     /::\  \     /:/  /::\  \     /::\  \     /:/  /      /::\  \   
   /:/\:\  \   /:/\:\  \   /:/\:\  \   /:/  /:/\:\  \   /:/\:\  \   /:/  /      /:/\ \  \  
  /:/  \:\__\ /::\~\:\  \ /::\~\:\__\ /:/  /:/  \:\  \ /:/  \:\  \ /:/  /  ___ _\:\~\ \  \ 
 /:/__/ \:|__/:/\:\ \:\__/:/\:\ \:|__/:/__/:/__/ \:\__/:/__/ \:\__/:/__/  /\__/\ \:\ \ \__\
 \:\  \ /:/  \:\~\:\ \/__\:\~\:\/:/  \:\  \:\  \ /:/  \:\  \  \/__\:\  \ /:/  \:\ \:\ \/__/
  \:\  /:/  / \:\ \:\__\  \:\ \::/  / \:\  \:\  /:/  / \:\  \      \:\  /:/  / \:\ \:\__\  
   \:\/:/  /   \:\ \/__/   \:\/:/  /   \:\  \:\/:/  /   \:\  \      \:\/:/  /   \:\/:/  /  
    \::/__/     \:\__\      \::/__/     \:\__\::/  /     \:\__\      \::/  /     \::/  /   
     ~~          \/__/       ~~          \/__/\/__/       \/__/       \/__/       \/__/    
     
     He who gives up freedom for safety deserves neither.___Benjamin Franklin
```

## Introduction

deblocus is similar to VPN, could establish encrypted tunnels with the remote gateway to achieve access to the remote network, and  provide high secure and performance local proxy (socks5/http) service. 

![deblocus-introduction.png](https://i.imgur.com/FP5A7hE.png)

## Features

There is some finished features:

- Traffic was encrypted by dynamic and unique secert-key per session (not static password).
- C/S supports multiple peers and delivers payloads parallelly.
- Use tcp multiplexer for inter endpoints of tunnel.
- Much reduced latency by Fast-open (nearly 1-RTT responded).

What does it mean?

- Don't worry about history traffic could be decrypted by others in anytime (Past, future, and present you will be at risk if lost password based on fixed-key encryption)
- Could share server with friends, and use multiple servers to offload single gateway traffic.
- Lots of short requests will be accelerated by multiplexer through tunnel (if have high packet lose rate or establish new network connection difficultly on your access network or transmission network, it will save time from without TCP 3-handshakes, and save much time  from if establish new connection when syn or ack+syn lost until tcp stack waiting timeout. Although do nothing but the acceleration is existing really, that cause of using old established connection could be able to gain more active transmitting strategies and lost detection (such as fast-restransmit, SACK) fully by kernel, then will enjoy fast responses and high throughput)

## Usage

Download [Binaries](https://github.com/Lafeng/deblocus/releases) or Compile from source, and should read [Wiki](https://github.com/Lafeng/deblocus/wiki) at first.

### GUI Manager

This tool is useful to manage deblocus process and can minimize to tray.

![deblocusMonitor.png](https://i.imgur.com/pdBpKN6m.png)

Download [deblocusMonitor](https://deblocus.codeplex.com/releases).

## Acknowledgements

[qtunnel](https://github.com/getqujing/qtunnel), [osext](https://bitbucket.org/kardianos/osext), [dhkx](https://github.com/monnand/dhkx) and [glog](https://github.com/golang/glog), thanks to those projects.

## Code License:

[MIT License](https://github.com/tvvocold/deblocus/blob/master/LICENSE)


# deblocus

## 介绍

建立与远端网关的加密通道，为本地应用提供安全高效的局部代理(Socks5/Http)服务。

可以满足在企业家庭校园内网、公共Wifi、通信窃听拦截等不可靠链路上实现高保密及高隐私的通信需要。

## 使用

详见 [Wiki](https://github.com/Lafeng/deblocus/wiki)，直接下载可从项目Releases.
