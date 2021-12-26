# HideSeek-Server

[![Linux](https://svgshare.com/i/Zhy.svg)](https://svgshare.com/i/Zhy.svg)  [![macOS](https://svgshare.com/i/ZjP.svg)](https://svgshare.com/i/ZjP.svg)  [![Windows](https://svgshare.com/i/ZhY.svg)](https://svgshare.com/i/ZhY.svg)




![](preview.png)

# What is HideSeek-Server?

HideSeek-Server is a part of HideSeek game. Obviously it implements all the important logic to make the work of HideSeek-Client available

---

Requirenments:

- Go(>=1.7.5)
- Bazel(>=4.2.2)
- Make(>=3.81)

---

"HideSeek-Server" can be run as a simple executable or a service. Makefile provides installation as a service or as an executable

Go to the root path for the game server and run

```
$ make

#If you want to install it as an executable
$ make install --executable

#If you want to install it as a service
$ make install --service
```