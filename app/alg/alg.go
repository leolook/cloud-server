package alg

import "fmt"

func BinarySearch(data []int, key int, start int, end int) int {

	mid := (end-start)/2 + start
	if data[mid] == key {
		return mid
	} else if start >= end {
		return -1
	}

	if key > data[mid] {
		return BinarySearch(data, key, mid+1, end)
	} else if key < data[mid] {
		return BinarySearch(data, key, start, mid-1)
	}

	return -1
}

//二分查找算法必须基于数组是有序数组
func BinarySearchExample() {

	data := []int{1, 3, 5, 7, 8, 9, 10, 12, 14, 16}

	index := BinarySearch(data, 3, 0, len(data))
	if index == -1 {
		fmt.Println("not find data")
	} else {
		fmt.Println("success to find data,index=", index)
	}

}

//冒泡排序算法
func BubbleSort(data []int) []int {

	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-1-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}

	return data
}

func BubbleSortExample() {
	data := []int{8, 3, 5, 7, 2, 4, 9, 1}
	sortedData := BubbleSort(data)
	fmt.Println(sortedData)
}

//选择排序
func SelectionSort(data []int) {

	for i := 0; i < len(data)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[minIndex] {
				minIndex = j //记录最小值的下标
			}
		}
		//将首位与最小值互换
		data[i], data[minIndex] = data[minIndex], data[i]
	}

}

func SelectionSortExample() {
	data := []int{8, 3, 5, 7, 2, 4, 9, 1}
	SelectionSort(data)
	fmt.Println(data)
}

//插入排序
//第一个元素默认已排序,将第二个元素与第一个元素比较，若 a1<a0,则a1插入到a0前面
//依次，第三个元素与第二个元素、第一个元素比较，比谁小，就插入到谁前面
//等等

//从第二个元素开始与已排序的元素比较
func InsertSort(data []int) {

	//2,4,7,6
	preIndex := 0
	for i := 1; i < len(data); i++ {

		preIndex = i - 1
		current := data[i]

		for {
			if preIndex < 0 || current > data[preIndex] {
				break
			}
			data[preIndex+1] = data[preIndex]
			preIndex--
		}

		data[preIndex+1] = current
	}
}

func InsertSortExample() {
	data := []int{8, 3, 5, 7, 2, 4, 9, 1}
	InsertSort(data)
	fmt.Println(data)
}
