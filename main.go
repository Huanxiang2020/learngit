package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var msg string
	reader := bufio.NewReader(os.Stdin) // 标准输入输出
	msg, _ = reader.ReadString('\n')    // 回车结束
	msg = strings.TrimSpace(msg)        // 去除最后一个空格 \r\n
	fmt.Println(msg)
}
