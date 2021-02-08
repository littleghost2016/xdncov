# xdncov
本仓库仅用作练习`gocolly`爬虫，不可用于其他用途，若出现任何问题，与作者本人无关。
如果觉得本程序对您有帮助，可以点一个`star`对作者进行鼓励，非常感谢！

## 准备工作

### 文件准备

创建`configs`文件夹，按照`example.config`，改成自己的配置文件并更改文件名以`.toml`为后缀（文件名随意）。

例如`11111111111.toml`

```toml
# 已返校同学可只更改前三项
name     	= "张三"	    # 姓名
id       	= 11111111111   # 学号
password 	= "11111111"    # 统一认证登录密码

# 南校区的以下四项内容可不更改，其他定位请自行修改
province    = "陕西省"
city        = "西安市"
area        = "陕西省 西安市 长安区"
address     = "陕西省西安市西沣路兴隆段266号"

# 请根据个人实际情况填写
szgjcs      = ""
szcs        = ""
szgj        = ""
zgfxdq      = "0"
mjry        = "0"
csmjry      = "0"
tw          = "2"
sfcxtz      = "0"
sfjcbh      = "0"
sfcxzysx    = "0"
qksm        = ""
sfyyjc      = "0"
jcjgqr      = "0"
remark      = ""
sfzx        = "0"
sfjcwhry    = "0"
sfjchbry    = "0"
sfcyglq     = "0"
gllx        = ""
glksrq      = ""
jcbhlx      = ""
jcbhrq      = ""
ismoved     = "0"
bztcyy      = ""
sftjhb      = "0"
sftjwh      = "0"
sfjcjwry    = "0"
jcjg        = ""

# 以下内容会由程序自动更改，如果不清楚如何正确修改，可不用理会
cookie   	= "" # 用作持久化
path     	= "" # 文件保存路径
lastupdatetime = 2020-01-01T00:00:00Z # 最后一次更新时间

# 需要微信通知（Server酱）的填写
# http://sc.ftqq.com
SCKEY       = "SCU89912...f4a70230"
```

*提交时将优先使用`cookie`，即可自行修改`cookie`值，而不泄露登陆密码。但此方法存在`cookie`过期的风险。*

## 运行

直接下载对应平台的可执行文件，执行即可。无需克隆整个仓库；或者说，克隆以后，除了`configs`文件夹和下载好的二进制文件，其余皆可删除。

**非Winsows平台执行文件前，请先给予执行权限（+x）。**

### PowerShell on Windows

```powershell
.\xdncov_yqt_windows_amd64.exe
```

### bash/zsh on Linux

```bash
./xdncov_yqt_linux_amd64
```

### bash/zsh on MacOS

```bash
./xdncov_yqt_darwin_amd64
```

## 使用.service时

请根据需求自行更改以下内容

- User
- Group
- WorkingDirectory  即可执行程序所在的文件夹。

当更改了User和Group并再次启动服务后，可能会因系统自动生成的`.log`文件读写权限问题导致任务无法进行。若想保留配置，则使用`chown`命令更改`.log`所属用户和用户组；若对以前已经保存的`.log`文件不感冒，则直接删除`.log`即可。

## 去哪里下载？

*`GitHub`最近在改版，`releases`被移动到了右边。*

在本仓库的主页面，拉到最上，往右看，第一个是`About`，再接着往下，就能看到`Releases`，对，编译好的二进制文件就在这里下载。

## 关于server酱的使用

本功能来自[@SewellDinG](https://github.com/SewellDinG)的PR [#7](https://github.com/littleghost2016/xdncov/pull/7)。

1. http://sc.ftqq.com/ 获取GitHub用户名，登录。
2. 点击“微信推送”，绑定微信后，获取SCKEY，填入自己的`toml`文件`SCKEY`的值中。
3. 每次提交（无论成功与否）均会通过微信推送结果。

## 后续工作

- [x] 持久化存储
- [x] 定时执行
- [x] toml添加最后一次提交时间
- [x] 日志输出
- [x] systemd进程守护
- [x] server酱  感谢师兄[@SewellDinG](https://github.com/SewellDinG)的PR
- [ ] 邮件提醒
- [ ] 未完成队列
