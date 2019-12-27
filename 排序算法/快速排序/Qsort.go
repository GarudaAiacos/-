package main

import (
	"fmt"
)

/**
快速排序：分治法+递归实现
随意取一个值A,从end开始找比A小的，从start开始找比A大的，若start和end不相遇，则交换start和end所对应的值
直到start和end相遇，则将end找到的值与A交换，第一轮排序结束，此时A左边的值都比A小，A右边的值都比A大
接下来对A两边的数据继续上面的步骤
*/
func qsort(arr []int, first, last int) {

	flag := first //基准值
	left := first
	right := last //数组长度

	if first >= last {
		return
	}

	//将大于arr[flag]的都放在右边，小于的，都放在左边(第一轮交换)
	for first < last {
		// 如果flag从左边开始，那么是必须先从右边开始找比flag小的值
		for first < last {
			if arr[last] >= arr[flag] {
				last--
				continue
			}
			swap(arr, last, flag)
			flag = last
			break
		}

		//从左边开始找比flag大的值
		for first < last {
			if arr[first] <= arr[flag] {
				first++
				continue
			}
			swap(arr, first, flag)
			flag = first
			break
		}
	}
	//递归调用
	qsort(arr, left, flag-1)
	qsort(arr, flag+1, right)
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	Arr := []int{23, 65, 13, 27, 42, 15, 38, 21, 4, 10}
	qsort(Arr, 0, len(Arr)-1)
	fmt.Println(Arr)
}
