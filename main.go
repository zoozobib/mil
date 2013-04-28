// mil project main.go
package main

import (
	"fmt"
	"math/rand"
	//"strconv"
)

var target []int = []int{}

//var cc chan int = make(chan int)

func data() []int {
	d := []int{}
	for i := 1; i < 10000001; i++ {
		d = append(d, rand.Intn(10000000))
	}
	return d
}

func has(tem []int, ddd int) bool {
	var r bool
	if len(tem) == 0 {
		r = true
	} else {
		for i := 0; i < len(tem); i++ {
			if tem[i] == ddd {
				r = false
				break
			} else {
				r = true
			}
		}
	}
	return r
}

func selected(c int, dd []int, cc chan []int) {

	temps := []int{}
	count := c * 10000

	if count+10000 <= 10000000 {
		for b := 0; b < 3; b++ {
			temp := 1
			for a := count; a < count+10000; a++ {

				if dd[a] > temp && has(temps, dd[a]) {
					temp = dd[a]
				}
			}

			if temp != 1 {
				temps = append(temps, temp)
			}
		}
		cc <- temps
	}
	//return temps
}

func main() {
	dd := data()
	var ss [][]int
	for j := 0; j < 1000; j++ {
		cc := make(chan []int)
		go selected(j, dd, cc)
		ss = append(ss, <-cc)

	}
	//fmt.Println(ss[0][1])
	temps := []int{}
	for b := 0; b < 3; b++ {
		temp := 1
		for i := 0; i < len(ss); i++ {
			for j := 0; j < len(ss[i]); j++ {
				if ss[i][j] > temp && has(temps, ss[i][j]) {
					temp = ss[i][j]
				}
			}
		}
		if temp != 1 {
			temps = append(temps, temp)
		}
	}
	fmt.Println(temps)
}
