# 三环境编译插件

## 使用方法
### 引入conf包，初始化：

```
import "conf"

conf.Setup({
    "name": "project_name",
    "directory": "config file directory",
})
```

### 使用  
有两种方法设置变量
1. 通过执行参数
```
# 以测试环境配置执行代码
./<exec_file> --env="test"
```

2. 通过环境变量  
```
$ vi ~/.bash_profile
export <product_name>="develop"
$ source ~/.bash_profile

# 自动使用全局变量中设置的环境
$ go run <product_name>
```

## 说明
在项目执行时根据

**注意：参数优先于环境变量，同时配置情况下以参数为准**