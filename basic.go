package main

import (
	"fmt"
	"os/user"
	"time"
)

func main() {
	fmt.Println("hoge", "test", time.Now())
	fmt.Println(user.Current())
}

// コマンドでドキュメントのパッケージにアクセスする

// go doc fmt.Println
