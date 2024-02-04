package main
import "fmt"

func main() {
    // 10の要素を持つ配列を空で作成する
    var inputs [10]string

    // 10回の標準入力を行う
    for i := 0; i < 10; i++ {
    // inputs配列のi番目のポインタを取得し、
    // fmt.Scanで標準入力から値を取得し、引数で指定された変数（inputs[i]）に格納する
        fmt.Scan(&inputs[i])
    }

    for i := 0; i < 10; i++ {
        fmt.Print(inputs[i])
        if i < 9 {
            fmt.Print(" ")
        }
    }
    fmt.Println()
}