# xdncov
本仓库仅用作练习`gocolly`爬虫，不可用于其他用途，若出现任何问题，与本人无关。

## 准备工作

### 文件准备

创建`configs`文件夹，按照`configs`文件夹里面的`example.config`，改成自己的配置文件并更改文件名以`.toml`为后缀。

例如`11111111111.toml`

```toml
# 已返校同学可只更改前三项
name     = "张三"        # 姓名
id       = 11111111111 # 学号
password = "11111111"    # 密码

# 其他内容可不更改
province = "陕西省"
city     = "西安市"
district = "长安区"
address  = "陕西省西安市西沣路兴隆段266号"
ymtys    = 0 # 一码通颜色
tw       = 1 # 体温
sfzx     = 1 # 是否在校
sfcyglq  = 0 # 是否处于隔离期
sfyzz    = 0 # 是否有症状

# 以下内容可能会由程序自动更改
cookie   = "" # 用作持久化
path     = "" # 文件保存路径
lastupdatetime = "" # 最后一次更新时间
```

## 执行

下载对应平台的可执行文件，执行即可。

### PowerShell on Windows

```powershell
.\xdncov_windows_amd64.exe
```

### bash/zsh on Linux

```bash
./xdncov_linux_amd64
```

### bash/zsh on MacOS

```bash
./xdncov_darwin_amd64
```

## 后续工作

- [x] 持久化存储
- [x] 定时执行
- [x] toml添加最后一次提交时间
- [ ] 邮件提醒