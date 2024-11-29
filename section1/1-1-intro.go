package main // 表示文件中的所有其余代码都属于main包
import (
	"fmt"
	"math"
	"reflect"
	"strings"
) // 表示将使用fmt包中包含的文本格式代码

// 声明变量 可以直接在后面加等号直接赋值
var quantity int
var length, width float64
var customerName string

// 当程序运行时 首先运行main函数
func main() {
	// fmt.Println("Hello, Go form main") // Go中一切都是区分大小写的 fmt下面的Println就是大写开头的函数
	hello()

	fmt.Println(math.Floor(2.75), reflect.TypeOf(math.Floor(2.75)))                             // 向下取整
	fmt.Println(strings.Title("head first go"), reflect.TypeOf(strings.Title("head first go"))) // 将每一个单词首字母变成大写
}

func hello() {
	// fmt.Println("Hello, Go from hello")
}
