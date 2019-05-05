# three environment dotenv package

### Installation

Include the package locally in your repository.
`npm install dotenv-go-runtime`

### Description

```
import "conf"

conf.Setup({
    "name": "project_name",
    "directory": "config file directory",
})
```

### Usage

There are two ways to set env variables
1.  By command parameters

```
# example
./<exec_file> --env="test"
```

2. Through PATH environment variables 
```
# config your system variables
$ vi ~/.bash_profile
export <product_name>="develop"
$ source ~/.bash_profile

$ go run <product_name>
```

###### Create a environment directory in your repository

###### Create three environment ini files in this directory
```
environment = "dev"
debug = true
prefixUrl = "127.0.0.1"
port = 3000

[mysql]
name = "project-name"
username = "root"
password = "123456"
host = "127.0.0.1"
port = 8306
```

### Description

**Note: Parameters take precedence over system environment variables**

### LICENSE

MIT