package hannota

import "fmt"

func PrintHannota(num int) {
	if num == 0 {
		return
	}
	proess(num,"left","right","mid")
}
func PrintHannota2(num int) {
	if num == 0 {
		return
	}
	proess2(num,"left","right","mid")
}

func move(num int, from, to string) {
	fmt.Printf("move num:%d from:%s to:%s \n", num, from, to)
}

func proess(num int, from, to, help string) {
	if num == 1 {
		move(num, from, to)
	}else{
		proess(num-1,from,help,to)
		move(num,from,to)

		proess(num-1,help,to,from)
	}
}
//必须借助中间的实现方法
func proess2(num int, from, to, help string) {
	if num == 1 {
		move(num, from, help)
		move(num, help, to)
	}else{
		proess2(num-1,from,to,help)//先将num-1放到右边，不考虑实现过程，交给子过程
		move(num,from,help)//两步移动num值，先将num放到中间
		proess2(num-1,to,from,help)//再把num-1放到右边，不考虑实现过程，交给子过程
		move(num,help,to)//两步移动num值，将num从中间放到右边，完成num的移动
		proess2(num-1,from,to,help)//真正的递归深入问题
	}
}
