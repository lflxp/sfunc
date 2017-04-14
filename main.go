package main

import (
	"fmt"
	"sf/floyd"
)

func main() {
	var n,m int
    	fmt.Println("请输入顶点个数和边的条数")
	fmt.Scanf("%d %d",&n,&m)
	init := &floyd.Data{
		N:n,
		M:m,
	}
	floyd.Start(init)
}