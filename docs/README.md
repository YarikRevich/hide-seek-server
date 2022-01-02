# HideSeek-Server

[![Linux](https://svgshare.com/i/Zhy.svg)](https://svgshare.com/i/Zhy.svg)  [![macOS](https://svgshare.com/i/ZjP.svg)](https://svgshare.com/i/ZjP.svg)



![](preview.png)

# What is HideSeek-Server?

HideSeek-Server is a part of HideSeek game. Obviously it implements all the important logic to make the work of HideSeek-Client available

Before gameserver installation you should install [Bazel](https://docs.bazel.build/versions/main/install.html)

"HideSeek-Server" can be run as a simple executable or a service. Makefile provides installation as a service or as an executable

Go to the **ROOT PATH** for the game server and run

```
$ make prepare
$ sudo make install_deps
$ sudo make build

#The default installation type is executable, it
#means that there will be installed only executable file
#which you should manage to run

$ sudo make install

#If you want to install it as a service run
$ sudo make install type=service
```