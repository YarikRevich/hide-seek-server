# HideSeek-Server

[![Linux](https://svgshare.com/i/Zhy.svg)](https://svgshare.com/i/Zhy.svg)  [![macOS](https://svgshare.com/i/ZjP.svg)](https://svgshare.com/i/ZjP.svg)



![](preview.png)

# What is HideSeek-Server?

HideSeek-Server is a part of HideSeek game. Obviously it implements all the important logic to make the work of HideSeek-Client available

Before gameserver installation you should install [Bazel](https://docs.bazel.build/versions/main/install.html)

"HideSeek-Server" can be run as a simple executable or a service. Makefile provides installation as a service or as an executable

Go to the **ROOT PATH** for the game server and run

```
# Adds not installed deps to the list
$ bazel build prepare

# Installs deps stated in a list
$ sudo bazel build install_deps

# Runs hide-seek-server ecosystem in docker containers(server, monitoring tools)
$ sudo bazel build run
```
