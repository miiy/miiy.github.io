/*
在一个包目录中，以_test.go结尾的文件不是go build命令编译的目标，而是go test编译的目标。
*/
package Testinga

import (
	"testing" // 每个测试函数必须导入testing包
)

func main()  {

}


// 测试函数的名字必须以Test开头，可选的后缀必须以答谢字母开头
func TestName(t *testing.T) {
	// ...
}