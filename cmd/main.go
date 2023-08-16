package main

import (
	"fmt"

	"github.com/o98k-ok/appscripts/template"
)

func main() {
	m3 := template.NewMenuBuilder("Chrome", []string{"文件", "分享", "备忘录"})
	fmt.Println(m3.Build())

	m2 := template.NewMenuBuilder("Chrome", []string{"文件", "新标签页"})
	fmt.Println(m2.Build())

	m1 := template.NewMenuBuilder("Chrome", []string{"Chrome"})
	fmt.Println(m1.Build())

	fmt.Println(template.NewMenuBuilder("Chrome", []string{"视图", "开发者", "JavaScript 控制台"}).Build())
	fmt.Println(template.NewMenuBuilder("Chrome", []string{"标签页", "关闭其他标签页"}).Build())
	fmt.Println(template.NewMenuBuilder("Chrome", []string{"标签页", "关闭右侧标签页"}).Build())
	fmt.Println(template.NewMenuBuilder("Chrome", []string{"标签页", "搜索标签页"}).Build())
	fmt.Println(template.NewMenuBuilder("Chrome", []string{"窗口", "扩展程序"}).Build())
	fmt.Println(template.NewMenuBuilder("Chrome", []string{"窗口", "移到“内建视网膜显示器”"}).Build())
	fmt.Println(template.NewMenuBuilder("Chrome", []string{"窗口", "移到“DELL U2720QM”"}).Build())
}
