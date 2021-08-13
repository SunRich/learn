package likou

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	//判断n为正数还是负数
	if n < 0 {
		return 1 / process(x, -n)
	} else {
		return process(x, n)
	}

}

func process(x float64, n int) (result float64) {
	//没有问题了
	if n == 1 {
		result = 1.0 * x
		return
	}
	//判断奇数偶数
	if (n & 1) != 0 {
		result = process(x*x, n>>1) * x
	} else {
		result = process(x*x, n>>1)
	}
	return
}
