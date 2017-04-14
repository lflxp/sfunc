package floyd

type Floyd interface {
	//初始化
	Init()	[10][10]int
	//Floyd-Warshall算法核心
	Parse(data [10][10]int)
}

func Start(a Floyd) {
	result := a.Init()
	a.Parse(result)
}