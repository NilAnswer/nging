# Nging V5

![Nging's logo](https://github.com/admpub/nging/blob/master/public/assets/backend/images/nging-gear.png?raw=true)

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/admpub/nging)

> 注意：这是Nging V5源代码，旧版V4.x、V3.x、V2.x、V1.x已经转移到 [v4分支](https://github.com/admpub/nging/tree/v4) [v3分支](https://github.com/admpub/nging/tree/v3) [v2分支](https://github.com/admpub/nging/tree/v2) [v1分支](https://github.com/admpub/nging/tree/v1)

    Nging支持MySQL和SQLite3数据库

Nging是一个网站服务程序，可以代替Nginx或Apache来搭建Web开发测试环境，并附带了实用的周边工具，例如：计划任务、MySQL管理、Redis管理、FTP管理、SSH管理、服务器管理等。


## 可执行文件下载

* [最新版下载地址](http://dl.webx.top/nging/latest/)

* [最新版备用地址](http://dl2.webx.top/nging/latest/)

## 安装方式

1. 安装Nging

    1). 自动安装方式:

    ```sh
    sudo sh -c "$(wget https://raw.githubusercontent.com/admpub/nging/master/nging-installer.sh -O -)"

    # 如果是中国境内网络，可以选择采用以下命令：
    sudo sh -c "$(wget https://gitee.com/admpub/nging/raw/master/nging-installer.sh -O -)"
    ```

    或

    ```sh
    sudo wget https://raw.githubusercontent.com/admpub/nging/master/nging-installer.sh -O ./nging-installer.sh && sudo chmod +x ./  nging-installer.sh && sudo ./nging-installer.sh
    ```

    nging-installer.sh 脚本支持的命令如下

    命令 | 说明
    :--- | :---
    `./nging-installer.sh` 或 `./nging-installer.sh install` | 安装(自动下载nging并启动为系统服务)
    `./nging-installer.sh upgrade` 或 `./nging-installer.sh up` | 升级
    `./nging-installer.sh uninstall` 或 `./nging-installer.sh un` | 卸载

    2). 手动安装方式:  
    下载相应平台的安装包，解压缩到当前目录，进入目录执行名为“nging”的可执行程序(在Linux系统，执行之前请赋予nging可执行权限)。 例如在Linux64位系统，分别执行以下命令：

    ```sh
    cd ./nging_linux_amd64
    chmod +x ./nging
    ./nging
    ```

    3). [Docker 安装方式](./README_docker.md)

2. 配置Nging:  
    打开浏览器，访问网址 <http://localhost:9999/setup> ，
    在页面中配置数据库和管理员账号信息进行安装。

安装成功后，使用管理员账号登录。

## Nging手动升级步骤

0. 备份数据库和旧版可执行文件；
1. 停止旧版本程序的运行；
2. 将新版本所有文件复制到旧版文件目录里进行覆盖；
3. 启动新版本程序；
4. 登录后台检查各项功能是否正常；
5. 升级完毕

## V3 升级到 V4
将 `config/config.yaml` 文件内的 `caddy`、 `ftp`、`download` 配置块移动到 `extend` 块内(ftp改名为ftpserver)。即：
```
extend {
    caddy {
        // 内容略...
    }
    ftpserver {
        // 内容略...
    }
    download {
        // 内容略...
    }
}
```

## 开机自动运行

1. 首先，安装为服务，执行命令 `./nging service install`
2. 启动服务，执行命令 `./nging service start`

与服务相关的命令：

命令 | 说明
:--- | :---
`./nging service install` | 安装服务
`./nging service start` | 启动服务
`./nging service stop` | 停止服务
`./nging service restart` | 重启服务
`./nging service uninstall` | 卸载服务

## Ⅰ、[功能介绍](doc/feature.md)

## Ⅱ、先睹为快

### 运行

[![安装](https://gitee.com/admpub/nging/raw/master/preview/preview_cli.png?raw=true)](https://gitee.com/admpub/nging/raw/master/preview/preview_cli.png)

### 安装：

[![安装](https://gitee.com/admpub/nging/raw/master/preview/preview_install.png?raw=true)](https://gitee.com/admpub/nging/raw/master/preview/preview_install.png)

### 登录：

[![登录](https://gitee.com/admpub/nging/raw/master/preview/preview_login.png?raw=true)](https://gitee.com/admpub/nging/raw/master/preview/preview_login.png)

### 系统信息：

[![系统信息](https://gitee.com/admpub/nging/raw/master/preview/preview_sysinfo.png?raw=true)](https://gitee.com/admpub/nging/raw/master/preview/preview_sysinfo.png)

### 实时状态：

[![实时状态](https://user-images.githubusercontent.com/512718/59155431-376ebe00-8abc-11e9-8d29-cee91978e574.png)](https://user-images.githubusercontent.com/512718/59155431-376ebe00-8abc-11e9-8d29-cee91978e574.png)


### 在线编辑文件：

[![在线编辑文件](https://gitee.com/admpub/nging/raw/master/preview/preview_editfile.png?raw=true)](https://gitee.com/admpub/nging/raw/master/preview/preview_editfile.png)

### 添加计划任务：

[![添加计划任务](https://gitee.com/admpub/nging/raw/master/preview/preview_task.png?raw=true)](https://gitee.com/admpub/nging/raw/master/preview/preview_task.png)

### MySQL数据库管理：

[![MySQL数据库管理](https://gitee.com/admpub/nging/raw/master/preview/preview_listtable.png?raw=true)](https://gitee.com/admpub/nging/raw/master/preview/preview_listtable.png)


## Ⅲ、[功能介绍](doc/developer.md)


请注意，本系统的源代码基于AGPL协议发布，不管您使用本系统的完整代码还是部分代码，都请遵循AGPL协议。  
> 如果需要更宽松的商业授权协议，请联系我购买授权。
