package floyd

import (
	"testing"
	"fmt"
)

func TestData_Init(t *testing.T) {
	var n,m int
    	fmt.Println("请输入顶点个数和边的条数")
	fmt.Scanf("%d %d",&n,&m)
	init := &Data{
		N:n,
		M:m,
	}
	data := init.Init()
	fmt.Println("原始数据")
	for d:=0;d<len(data);d++ {
		fmt.Printf("%d %d %d %d %d %d %d %d %d %d\n",data[d][0],data[d][1],data[d][2],data[d][3],data[d][4],data[d][5],data[d][6],data[d][7],data[d][8],data[d][9])
	}
}

func TestData_Parse(t *testing.T) {
	d := Data{N:4,M:8}
	var data = [10][10]int{
		{0,0,0,0,0,0,0,0,0,0},
		{0,0,2,6,4,0,0,0,0,0},
		{0,99999999,0,3,99999999,0,0,0,0,0},
		{0,7,99999999,0,1,0,0,0,0,0},
		{0,5,99999999,12,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0,0},
	}
	d.Parse(data)
}

func TestStart(t *testing.T) {
	var n,m int
    	fmt.Println("请输入顶点个数和边的条数")
	fmt.Scanf("%d %d",&n,&m)
	init := &Data{
		N:n,
		M:m,
	}
	Start(init)
}

func BenchmarkData_Parse(b *testing.B) {
	d := Data{N:4,M:8}
	var data = [10][10]int{
		{0,0,0,0,0,0,0,0,0,0},
		{0,0,2,6,4,0,0,0,0,0},
		{0,99999999,0,3,99999999,0,0,0,0,0},
		{0,7,99999999,0,1,0,0,0,0,0},
		{0,5,99999999,12,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0,0},
	}

	for i:=0;i<b.N;i++ {
		d.Parse(data)
	}
}