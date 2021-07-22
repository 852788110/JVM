package main

import (
	"fmt"
	"io/ioutil"
	"jvmgo/jvm/rtda/heap"
)

func main() {
	path := "/Users/liujie/IdeaProjects/jvm/out/production/jvm/lucky.class"
	// 读取字节码
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	// 解析字节码
	class := heap.ParseClass(data)

	// 得到字节码结构
	fmt.Println(class)
}
