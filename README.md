= # tmux_pm25 =
在Tmux的状态栏中显示空气质量指数，主要是pm2.5指数

## 安装指南
### 使用 [Tmux Plugin Manager](https://github.com/tmux-plugins/tpm) (推荐)

在`.tmux.conf`文件中，将下行代码加入你的TPM插件列表中:

set -g @plugin 'tmux-plugins/tmux-cpu'
    
输入 `prefix + I` 重载配置。
    
### 手动安装
    
复制此库:
    
$ git clone https://github.com/tmux-plugins/tmux-cpu ~/clone/path
        
在`.tmux.conf`文件中，将下行代码加入你的TPM插件列表中:
        
run-shell ~/clone/path/cpu.tmux
          
重载配置
``` bash
tmux source-file ~/.tmux.conf
```

## 使用指南

此插件会在tmux的环境中添加一个新的*format name*, `pm25`。

只要在你想要此信息出现的地方加上这个*format name*即可，比如：
``` 
set -gq status-right '#{pm25} %d/%m/%y'
```

## TODO

- [X] 走通远端api
- [X] 建立缓存防止api连接数超过上限
- [ ] 根据IP确定要查找的城市空气质量数据
- [ ] 美化信息在状态栏上的呈现方式
- [ ] 用个人ApiKey代替测试ApiKey
