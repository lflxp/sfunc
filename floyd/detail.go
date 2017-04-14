package floyd
/*
现在需要一个数据结构来存储图的信息，我们仍然可以用一个 4*4 的矩阵（二维数组 e）来存储。比如 1 号城市到 2 号城市的路程为 2，则设 e[1][2]的值为 2。2 号城市无法到达 4 号城市，则设置 e[2][4]的值为 ∞。另外此处约定一个城市自己是到自己的也是 0，例如 e[1][1]为 0，具体如下。
*/

import (
	"fmt"
	//"strings"
)

//用inf(infinity的缩写)存储一个我们认为的正无穷值
const Inf int = 99999999

type Data struct {
	//读入n和m，n表示顶点个数，m表示边的条数
	N 	int
	M 	int
}

func (this *Data) Init() [10][10]int {
	var a,b,c int
	//datas := [this.N][this.M]int{}
	var datas [10][10]int
	//初始化
	for x:=1;x<=this.N;x++ {
		for y:=1;y<=this.N;y++ {
			if x == y {
				datas[x][y] = 0
			} else {
				datas[x][y] = Inf
			}
		}
	}
	//读入边
	fmt.Println("请输入节点号,下一个节点号,路径长度")
	for n:=1;n<=this.M;n++{
		fmt.Scanf("%d %d %d",&a,&b,&c)
		datas[a][b] = c
	}
	return datas
}

func (this *Data) Parse(data [10][10]int) {
	fmt.Println("原始数据")
	for d:=0;d<len(data);d++ {
		fmt.Printf("%d %d %d %d %d %d %d %d %d %d\n",data[d][0],data[d][1],data[d][2],data[d][3],data[d][4],data[d][5],data[d][6],data[d][7],data[d][8],data[d][9])
		//tmp := []string{}
		//for _,y := range data[d] {
		//	tmp = append(tmp,string(y))
		//}
		//fmt.Println(strings.Join(tmp," "))
	}
	//Floyd-Warshall算法核心
	//赛选最短路径
	for k:=1;k<=this.N;k++ {
		for i:=1;i<this.N;i++ {
			for j:=1;j<this.N;j++ {
				if data[i][j]>data[i][k]+data[k][j] {
					fmt.Printf("%d->%d [%d] 路径最优化路径修改:%d->%d->%d [%d]\n",i,j,data[i][j],i,k,j,data[i][k]+data[k][j])
					data[i][j] = data[i][k]+data[k][j]
				}
			}
		}
	}
	//输出最终的结果
	for i:=1;i<=this.N;i++ {
		for j:=1;j<=this.N;j++ {
			fmt.Printf("%d->%d 最短路径 %d\n",i,j,data[i][j])
		}
	}
}