package main

import (
	"fmt"
	"sort"
)

type NumArr struct {
	Index int
	Value int
}

type ByVal []NumArr

func (a ByVal) Len() int           { return len(a) }
func (a ByVal) Less(i, j int) bool { return a[i].Value > a[j].Value }
func (a ByVal) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func SwapVariable(a int,b int ) {

	a = a-b
	b = a+b
	a = b-a
	fmt.Println("First number :",a)
	fmt.Println("Second number :",b)
}


func printValue(arr []int) {

	listArr := []NumArr{}

	dict := make(map[int]int)

	for _ , num :=  range arr {
		dict[num] = dict[num]+1
	}

	for idx, val := range dict {
		listArr = append(listArr, NumArr{
			Index:   idx,
			Value: val,
		})
	}

	sort.Sort(ByVal(listArr))

	fmt.Println("number --> total")

	for _, val := range listArr {

		fmt.Println(val.Index," --> ",val.Value)
	}

}

func MergeString(s1 string, s2 string) (result string) {

	asS1 := []rune(s1)
	asS2 := []rune(s2)
	// For every index in the strings

	for i := 0; i < len(s1) || i < len(s2); i++	{

		// First choose the ith character of the
		// first string if it exists
		if i < len(s1) {
			result += string(asS1[i])
		}

		// Then choose the ith character of the
		// second string if it exists
		if i < len(s2) {
			result += string(asS2[i])
		}

	}

	fmt.Println(result)
	return

}

func TwoString(a string, b string) (answer string) {

	answer = ""
	a += "~"
	b += "~"
	i := 0
	j := 0

	asS1 := []rune(a)
	asS2 := []rune(b)

	for string(asS1[i]) != "~" || string(asS2[j]) != "~" {

		if string(asS1[i]) != "~" &&  a[i:] < b[j:] {
			answer += string(asS1[i])
			i += 1
		} else {
			answer += string(asS2[j])
			j += 1
		}

	}

	fmt.Println(answer)

	return
}


func arrayManipulation(n int32, queries [][]int32) int64 {

	arr := make([]int64, n+1)

	for _, i := range queries {
		a := int(i[0])
		b := i[1]
		k := int64(i[2])
		arr[a] += k
		if (b + 1) <= n {
			arr[b+1] -= k
		}
	}

	var x, max int64 = 0, 0

	for i := 1; i <= int(n); i++ {
		x += arr[i]
		if max < x {
			max = x
		}
	}

	return max
}

func main() {

	//Soal 1
	fmt.Println("============= Soal 1")
	SwapVariable(3,5)
	fmt.Println("============= Soal 2")
	//Soal 2
	inputArray := []int{4,6,3,5,4,6,7,8,3,4,6,7,5,4,6,4,4,5,6}
	printValue(inputArray)
	//Soal 3
	fmt.Println("============= Soal 3")
	MergeString("saya" , "kamu")
	MergeString("kosong" , "ada")
	MergeString("ada" , "kosong")
	//Soal 5
	fmt.Println("============= Soal 5")
	TwoString("jack","daniel")
	//Soal 4
	fmt.Println("============= Soal 4")
	var inputArr = [][]int32{ {1,5,3}, {4,8,7}, {6,9,1}}
	fmt.Println(arrayManipulation(10,inputArr))

}