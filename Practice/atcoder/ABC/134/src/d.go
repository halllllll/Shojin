// 初回というか半分から向こう側はいじれないからこっち側半分を考えると...
// くそむずい どうやって実装するんだこれ
package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	A := make([]int, N+1)
	for i := 0; i < N; i++ {
		var V int
		fmt.Scan(&V)
		A[i+1] = V
	}
	ANS := make([]int, N+1)
	for i := N; i > 0; i-- {
		s := 0
		for j := i + i; j < N+1; j += i {
			// s += A[j]	は？？？？？？？？？？？？？？？？？？？？？
			s += ANS[j]
		}
		if s%2 != A[i] {
			ANS[i] = 1
		}
		// fmt.Printf("i=%d, now ans %v\n", i, ANS[1:])
	}
	SUM := 0
	for _, a := range ANS {
		SUM += a
	}
	fmt.Println(SUM)
	arr := []int{}
	for i := 1; i < N+1; i++ {
		if ANS[i] == 1 {
			arr = append(arr, i)
		}
	}
	for i := 0; i < len(arr); i++ {
		if i == len(arr)-1 {
			fmt.Printf("%d\n", arr[i])
		} else {
			fmt.Printf("%d ", arr[i])
		}
	}
}
