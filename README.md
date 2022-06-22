# bubble清单

一个基于gin+gorm开发的练手小项目(清单)。

前端页面基于vue和ElementUI开发。


### 配置MySQL
1. 在你的数据库中执行以下命令，创建本项目所用的数据库：
```sql
CREATE DATABASE bubble DEFAULT CHARSET=utf8mb4;
```
2. 在`bubble/conf/config.ini`文件中按如下提示配置数据库连接信息。

```ini
port = 9000
release = false

[mysql]
user = 你的数据库用户名
password = 你的数据库密码
host = 你的数据库host地址
port = 你的数据库端口
db = bubble
```

![image](https://user-images.githubusercontent.com/74490865/174916020-eb039412-7762-4414-a21e-632d33603881.png)


