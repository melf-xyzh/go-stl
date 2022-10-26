/**
 * @Time    :2022/6/23 11:05
 * @Author  :Xiaoyu.Zhang
 */

package main

import (
	"fmt"
	"github.com/melf-xyzh/go-stl/stack"
)

func main() {
	st := stack.NewStack()
	// 入栈
	st.Push("1")
	st.Push("2")
	st.Push("3")
	st.Push("4")
	fmt.Println("打印栈:",st.AllData())
	// 出栈
	data1 := st.Pop()
	fmt.Println("出栈：",data1)
	fmt.Println("打印栈:",st.AllData())
}
