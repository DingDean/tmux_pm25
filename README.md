# tmux_pm25
在Tmux的状态栏中显示空气质量指数，主要是pm2.5指数。    
![效果图](./tmux25.png)

## 安装指南

### 使用 [Tmux Plugin Manager](https://github.com/tmux-plugins/tpm) (推荐)

在`.tmux.conf`文件中，将下行代码加入你的TPM插件列表中:

set -g @plugin 'DingDean/tmux_pm25'
    
### 手动安装
    
复制此库:
    
$ git clone https://github.com/DingDean/tmux_pm25 ~/clone/path
        
在`.tmux.conf`文件中，将下行代码加入你的TPM插件列表中:
        
run-shell ~/clone/path/tmux_pm25.tmux

### 重载Tmux配置
``` bash
tmux source ~/.tmux.conf
```

## 使用指南

### 配置.tmux_25_config.json
空气指数数据有两个数据源，每个都需要apiKey:

1. 来自于[阿里云](https://market.aliyun.com/products/57126001/cmapi014302.html?spm=5176.730005.0.0.5OH11d#sku=yuncode830200000)。免费版的可以使用10000次请求, 按照现在每天24次的请求数量，绝对够使用一年。
2. 来自于[PM25.in](http://www.pm25.in), 请前往网站申请apiKey。

申请得到apiKey后，在您的`$HOME`目录下创建.tmux_25_config.json, 内容如下：
``` Json
{
  "apiKey": "你申请得到的apiKey",
  "city": "城市的中文名字或者拼音",
  "source": "aliyun或者pm25.in"
}
```

### 配置.tmux.cof

此插件会在tmux的环境中添加一个新的*format name*, `pm25`。

只要在你想要此信息出现的地方加上这个*format name*即可，比如：
``` 
set -gq status-right '#{pm25} %d/%m/%y'
```

## TODO

- [X] 走通远端api
- [X] 建立缓存防止api连接数超过上限
- [X] 根据IP确定要查找的城市空气质量数据
- [X] 用户指定一个城市后，不再根据其IP自动获取数据
- [X] 使用godep
- [X] 美化信息在状态栏上的呈现方式
- [X] 增加手动刷新数据的快捷键
- [ ] 显示更多的空气数据
