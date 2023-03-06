package main

// trace编程过程
// 1. 创建文件
// 2. 启动
// 3. 停止

// func main() {
// 	// 1. 创建trace文件
// 	f, err := os.Create("trace.out")
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	defer f.Close()
//
// 	// 2. 启动trace
// 	err = trace.Start(f)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	// 业务调试
// 	fmt.Println("hello GMP")
//
// 	// 3. 停止trace
// 	trace.Stop()
// }
