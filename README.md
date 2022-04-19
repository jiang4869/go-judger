# 感谢 [https://github.com/QingdaoU/Judger](https://github.com/QingdaoU/Judger)

# 用法
仅在Linux系统测试过

首先安装`judger`的依赖库
```
git clone https://github.com/QingdaoU/Judger
sudo apt-get install gcc g++ make cmake libseccomp-dev
mkdir build && cd build && cmake .. && make && sudo make install
```

获取`go-judger`
```
go get  github.com/jiang4869/go-judger
```

首先先在项目跟目录下创建一个`1.in`的输入文件，一个`main.c`要执行的c代码

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/jiang4869/go-judger"
)

func main() {
    // 先把main.c编译成main
	judger.ExecShell("gcc -o main main.c")

    // 传入各种参数，调用main读取1.in中的内容，输出到1.out里面
	r, err := judger.Run(1000, 2000, 128*1024*1024, 32*1024*1024, 10000, 200, 0, 0, 0, []string{}, []string{}, "main", "1.in", "1.out", "1.out", "judger.log", "c_cpp")
	if err != nil {
		fmt.Println(err)
		return
	}
	b, _ := json.Marshal(r)
	fmt.Println(string(b))
}

```