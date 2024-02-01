package main

import "fmt"

// 初期設定で利用される
func init() {
	fmt.Println("init!")
}

func hoge() {
	fmt.Println("fuga")
}

func main() {
	// main関数外にある関数は以下のようにして呼び出すことができる
	hoge()
	fmt.Println("hoge", "test")
}
