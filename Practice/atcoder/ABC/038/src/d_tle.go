// 横幅で昇順ソートすると、「同じ横幅の場合、縦幅は小さい方がいい」となる。
// なので横幅でソートしたあと、縦幅はLISを使える
// atcoderの過去問の言語バージョンが未だに古いままなので仕方なく19089475893年前の実装でやってやる

package main

import (
	"fmt"
	"sort"
)

type Present struct {
	H, W int
}

type ByAndMoreHnMoreW []Present

func (p ByAndMoreHnMoreW) Len() int {
	return len(p)
}
func (p ByAndMoreHnMoreW) Less(i, j int) bool {
	if p[i].H < p[j].H {
		return true
	} else if p[i].H == p[j].H && p[i].W < p[j].W {
		return true
	} else {
		return false
	}
}
func (p ByAndMoreHnMoreW) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	var N int
	fmt.Scan(&N)
	Prezents := make([]Present, N)
	for i := 0; i < N; i++ {
		var h, w int
		fmt.Scan(&h, &w)
		Prezents[i] = Present{H: h, W: w}
	}
	sort.Sort(ByAndMoreHnMoreW(Prezents))
	// NlogNがどうやっても実装できんのだが？？？？？？？？
	// しかたねぇからまずN^2をやる
	L := make([]int, N)
	ans := 0
	for i := 0; i < N; i++ {
		L[i] = 1
		for j := 0; j < i; j++ {
			if Prezents[j].H < Prezents[i].H && Prezents[j].W < Prezents[i].W {
				L[i] = max(L[i], L[j]+1)
			}
		}
		ans = max(ans, L[i])
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
