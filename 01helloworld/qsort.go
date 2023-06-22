package main

import "fmt"

/*
分治的思想
确定一个基准数，将所有比基准数晓得数放在左边，将所有比基准数大的数放在右边，即将基准数挪到他应该在的位置
从序列两端进行探测，先从右往左找一个小于基准数的数，然后从左往右找一个大于基准数的数，然后交换它们。
*/
func quickSort(values []int, left, right int) {
	temp := values[left] // 基准数的值
	p := left            // 基准数的位置
	i, j := left, right

	for i <= j {
		// 从右往左找到第一个小于基准数的数
		for j >= p && values[j] >= temp {
			j--
		}
		// 如果找到的小于基准数的的数在基准数的右边，将其和基准数换位置
		if j >= p {
			values[p] = values[j] // 也可写成values[p], values[j] = values[j], values[p]，这样最下面就不用values[p] = temp一句了
			p = j
		}
		for i <= p && values[i] <= temp {
			i++
		}
		if i <= p {
			values[p] = values[i] // 也可写成values[p], values[i] = values[i], values[p]，这样最下面就不用values[p] = temp一句了
			p = i
		}
		values[p] = temp
	}
	if p-left > 1 {
		quickSort(values, left, p-1)
	}
	if right-p > 1 {
		quickSort(values, p+1, right)
	}
}

func main() {
	values := []int{6, 1, 3, 5, 7, 2, 5, 10, 9, 8}
	quickSort(values, 0, len(values)-1)
	for _, v := range values {
		fmt.Println(v)
	}
}
