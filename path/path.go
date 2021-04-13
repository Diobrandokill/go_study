package main

import (
	"fmt"
	"path"
)

//go语言path包的学习

func main() {
	//返回路径的最后一个元素 ✅
	fmt.Println(path.Base("./a/b/c"))
	//如果路径为空字符串，返回. ✅
	fmt.Println(path.Base(""))
	//如果路径只有斜线，返回/ ✅
	fmt.Println(path.Base("///"))

	//返回等价的最短路径 ✅
	//1.用一个斜线替换多个斜线
	//2.清除当前路径.
	//3.清除内部的..和他前面的元素
	//4.以/..开头的，变成/
	fmt.Println(path.Clean("/..////a/b/../"))

	//返回路径最后一个元素的目录 ✅
	//路径为空则返回.
	fmt.Println(path.Dir("./a/b/c"))

	//返回路径中的扩展名 ✅
	//如果没有点，返回空
	fmt.Println(path.Ext("./a/b/c/d.md"))

	//判断路径是不是绝对路径 ✅
	fmt.Println(path.IsAbs("./a/b/c"))
	fmt.Println(path.IsAbs("/a/b/c"))

	//连接路径，返回已经clean过的路径 ✅
	fmt.Println(path.Join("./a", "b/c", "../d/"))

	//匹配文件名，完全匹配则返回true ✅
	// '*' - 匹配零个或多个非/的字符
	// '?' - 匹配一个非/的字符
	// '[abc]' or '[^abc]' or '[a-c]' - 匹配字符组；^为非的意思
	// 'c' - 匹配字符（ c != '*', '?', '\\', '[' ）
	// '\\' - 匹配字符 （ 可以为 '*', '?', '\\', '[' ）
	fmt.Println(path.Match("*", "a"))
	fmt.Println(path.Match("*", "a/b/c"))
	fmt.Println(path.Match("*/*/*", "aaa/b/c"))

	fmt.Println(path.Match("a?c", "abc"))
	fmt.Println(path.Match("a?c", "a/c"))

	fmt.Println(path.Match("[abc][123]", "a1"))
	fmt.Println(path.Match("[a-c][1-3]", "b3"))
	fmt.Println(path.Match("[a-c][^1-3]", "c9"))

	fmt.Println(path.Match("\\?\\b", "?b")) # \\c用于识别c

	//分割路径中的目录与文件 ✅
	fmt.Println(path.Split("./a/b/c/d.jpg"))
}
