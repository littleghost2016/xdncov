# xdncov
本仓库仅用作练习`gocolly`爬虫，不可用于其他用途，若出现任何问题，与本人无关。

## 准备工作

### 文件准备

创建`configs`文件夹，按照`example.config`，改成自己的配置文件并更改文件名以`.toml`为后缀。

例如`11111111111.toml`

```toml
# 已返校同学可只更改前三项
name     = "张三"        # 姓名
id       = 11111111111  # 学号
password = "11111111"   # 统一认证登录密码

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

# 以下内容会由程序自动更改，如果不清楚如何正确修改，可不用理会
cookie   = "" # 用作持久化。
path     = "" # 文件保存路径
lastupdatetime = 2020-01-01T00:00:00Z # 最后一次更新时间
```

*提交时将优先使用`cookie`，即可自行修改`cookie`值，而不泄露登陆密码。但此方法存在`cookie`过期的风险。*

## 运行

直接下载对应平台的可执行文件，执行即可。无需克隆整个仓库；或者说，克隆以后，除了`configs`文件夹和下载好的二进制文件，其余皆可删除。

**非Winsows平台执行文件前，请先给予执行权限（+x）。**

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

## 使用.service时

请根据需求自行更改以下内容

- User
- Group
- WorkingDirectory  即可执行程序所在的文件夹。

当更改了User和Group并再次启动服务后，可能会因系统自动生成的`.log`文件读写权限问题导致任务无法进行。若想保留配置，则使用`chown`命令更改`.log`所属用户和用户组；若对以前已经保存的`.log`文件不感冒，则直接删除`.log`即可。

## 后续工作

- [x] 持久化存储
- [x] 定时执行
- [x] toml添加最后一次提交时间
- [x] 日志输出
- [x] systemd进程守护
- [ ] 邮件提醒
- [ ] 未完成队列