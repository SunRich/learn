package sort

type ArraySort struct {
	List   []int64
	Length int
}

func NewArraySort(list []int64) (r *ArraySort) {
	return &ArraySort{
		List:   list,
		Length: len(list),
	}
}

func swap(result []int64, i, j int) {
	if i == j {
		return
	}
	result[i] = result[i] ^ result[j]
	result[j] = result[i] ^ result[j]
	result[i] = result[i] ^ result[j]
}
func (l *ArraySort) copy() (result []int64) {
	result = make([]int64, l.Length)
	copy(result, l.List)
	return
}

//冒泡排序
//每次比较满足条件后交换
//大循环一次后大循环次减少一次
func (l *ArraySort) BubbleSort() (result []int64) {
	result = l.copy()
	for i := 0; i < l.Length-1; i++ {
		for j := 0; j < l.Length-i-1; j++ {
			if result[j] > result[j+1] {
				swap(result, j, j+1)
			}
		}
	}
	return
}

//选择排序
//增加最小数的临时变量，对比一个循环确定最小数
func (l *ArraySort) SelectSort() (result []int64) {
	result = l.copy()
	for i := 0; i < l.Length-1; i++ {
		minIndex := i
		for j := i + 1; j < l.Length; j++ {
			if result[j] < result[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			swap(result, minIndex, i)
		}

	}
	return
}

func (l *ArraySort) InsertSort() (result []int64) {
	result = l.copy()
	for i := 1; i < l.Length; i++ {
		temp := result[i]
		j := i
		for j >= 1 && result[j-1] > temp {
			result[j] = result[j-1]
			j--
		}
		result[j] = temp
	}
	return
}

func (l *ArraySort) QuickSort() (result []int64) {
	result = l.copy()
	Quick(result, 0, l.Length-1)
	return
}
func Quick(result []int64, left, right int) {
	if left >= right {
		return
	}
	pivotIndex := partition(result, left, right)
	Quick(result, left, pivotIndex-1)
	Quick(result, pivotIndex+1, right)

}

func partition(result []int64, left, right int) int {
	pivot := left
	index := pivot + 1
	for i := index; i <= right; i++ {
		if result[i] < result[pivot] {
			swap(result, i, index)
			index++
		}
	}
	swap(result, pivot, index-1)
	return index - 1
}

func (l *ArraySort) HeapSort() (result []int64) {
	result = l.copy()
	buildHeap(result,len(result))
	for i:=len(result)-1;i>=0;i-- {
		swap(result,0,i)
		heapify(result,0,i)

	}
	Quick(result, 0, l.Length-1)
	return
}

func buildHeap(result []int64, length int) {
	for i:=(length>>1)-1;i>=0;i--{
		heapify(result,i,length)
	}

}
func heapify(result []int64, i, length int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < length && result[left] > result[largest] {
		largest = left
	}
	if right < length && result[right] > result[largest] {
		largest = right
	}
	if largest != i {
		swap(result, i, largest)
		heapify(result,largest,length)
	}
}
