package main

import "sync"

var Wg sync.WaitGroup

func main() {
	cmd := parseCmd()
	Wg.Add(1)
	if cmd.versionFlag {
		println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		// 创建一个JVM对象，并执行start方法
		newJVM(cmd).start()
		Wg.Wait()
	}
}
