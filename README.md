# ecnu-datasync-cli

## 简介
ecnu-datasync-cli 是一个开源的轻量级工具，用于[华东师范大学数据开放平台](https://developer.ecnu.edu.cn/doc/#/)的数据接口获取。

它提供了一种灵活、快速、方便的方式，允许在不用编写代码的情况下，将数据直接同步至 csv 或 xlsx 文件。

## 优势

无需代码编写，只需要编写相关配置文件，即可通过命令行命令实现同步数据至文件。

结合 datax 等第三方 ETL 工具，可以在无需编写代码的情况下实现各种同步任务。

## 用法

### 获取 ecnu-datasync-cli 安装程序可执行文件
1. 获取可执行文件
* 方法一：下载 ecnu-datasync-cli 可执行文件 [Releases page](https://github.com/ECNU/ecnu-datasync-cli/releases/latest)

  下载解压后可直接使用。

* 方法二：从源代码生成二进制文件
  ```shell
  git clone https://github.com/ECNU/ecnu-datasync-cli.git
  cd datasync-sdk-cli
  go build
  ```
    > 需要 go 1.20+ 环境  
      win 环境下请自行添加 .exe 后缀

  生成的二进制文件在 datasync-sdk-cli 中

2. 验证文件可用
    ```shell
    cd {PATH_TO_CLI}/datasync-sdk-cli
    ./ecnu-datasync-cli -v
    ```
   应输出版本信息

### 通过 ecnu-datasync-cli 同步数据到 csv 或 xlsx 文件

#### 不使用配置文件进行同步

不使用配置文件进行同步时，直接在命令行中指定参数。

##### 命令

> 请确保处于校园网环境或使用校园网VPN

```shell
./ecnu-datasync-cli -c {client_id} -s {client_secret} -a {api_path} -o {output_file}
```
> -o 的值既可以填写 csv 文件也可以填写 xlsx 文件，会自动根据文件的后缀名来判断导出的类型
##### 例子

* 使用示例密钥访问接口 https://api.ecnu.edu.cn/api/v1/sync/fakewithts?ts=0 ，同步数据到 path_to_csv/test.csv 文件

  ```shell
  ./ecnu-datasync-cli -c=123456 -s=abcdef -a='/api/v1/sync/fakewithts?ts=0' -o='path_to_csv/test.csv'
  ```

* 使用示例密钥访问接口 https://api.ecnu.edu.cn/api/v1/sync/fakewithts?ts=0 ，同步数据到 path_to_xlsx/test.xlsx 文件

  ```shell
  ./ecnu-datasync-cli -c=123456 -s=abcdef -a='/api/v1/sync/fakewithts?ts=0' -o='path_to_xlsx/test.xlsx'
  ```

#### 使用配置文件进行同步

您还可以使用指定配置文件的方式进行数据同步。

> 如果同时指定了 -config 和其他参数，那么会忽略其他参数，仅以配置文件中的配置进行接口调用。

##### 命令
  ```shell
  ./ecnu-datasync-cli -config {config_file}
  ```
##### 例子
1. 首先，在某目录 path_to_json 下创建一个示例配置文件 cfg.json ,并根据您的环境修改配置文件 cfg.json

```json
    {
      "oauth2_config":{
        "client_id":"client_id",
        "client_secret":"client_secret",
        "base_url":"https://api.ecnu.edu.cn",
        "scopes":["ECNU-Basic"],
        "timeout":10,
        "debug":false
      },
      "api_config":{
        "api_path":"/api/v1/sync/fakewithts?ts=0",
        "page_size":2000
      },
      "output_file":"./test.csv"    
    }
```
> 配置文件中，output_file 的值既可以填写 csv 文件也可以填写 xlsx 文件

2. 然后，使用配置文件 path_to_json/cfg.json 中的配置进行数据同步。

```shell
   ./ecnu-datasync-cli -config path_to_json/cfg.json
```
   
