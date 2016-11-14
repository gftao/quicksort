package main

import (
	"fmt"
	"math/rand"
	"time"
)

var TOT = int(20)

func main() {
	startt := time.Now().UnixNano()
	var arry []int
	rand.Seed(startt)//time.Now().UnixNano())

	for i:= 0; i < TOT; i++{
		arry = append(arry, rand.Intn(TOT))
	}
	fmt.Println("Before sort arry:")
	pfarry(arry)
	quicksort(arry)
	fmt.Println()
	fmt.Println("After sort arry:")
	pfarry(arry)
	endt := time.Now().UnixNano()

	//输出执行时间，单位为毫秒。
	fmt.Println("\nDuring:", (endt - startt) / 1000000)

}

func pfarry (list []int){

	len := int(1)
	for total := TOT; total > 0; total /= 10 {
		len ++
	}
	for i := 0; i < TOT; i++{
		fmt.Printf("%*d", len, list[i])
	}
}

func  quicksort(list []int)  {
	l:= len(list)
	if l <= 1{
		return
	}
	b := 0;
	e := l - 1
	key := list[e/2]
	for {
		for list[b] < key{
			b++
		}
		for list[e] > key {
			e--
		}

		if b >= e {
			break
		}
		swap(list, b, e)
		b++;e--		//每次较换后向后取值，否则（切片中有相同值时）会出现死循环
	}

	quicksort(list[:b])
	quicksort(list[e+1:])
}

func  swap(list []int, a, b int)  {
	list[a], list[b] = list[b] ,list[a]
}
