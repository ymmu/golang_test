package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, *p, p.X, (*p).X, v2, v3)
	var a [2]string
	a[0] = "Hi"
	a[1] = "Lucca"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13} //배열 크기 고정
	fmt.Println(primes)
	var s []int = primes[1:3]
	fmt.Println(s)

	// 슬라이스는 배열의 내용을 보여주는 것
	//슬라이스를 통해 값을 바꾸면 배열값이 바뀜
	//바뀐 배열값을 가지는 슬라이스들에서도 바뀐값을 확인할 수 있음
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	aa := names[0:2]
	bb := names[1:3]
	fmt.Println(aa, bb)

	bb[0] = "XXX"
	fmt.Println(aa, bb)
	fmt.Println(names)

	fmt.Println("========================")
	//다음처럼 참조변수 만들 수 있음. 배열수 지정 안 하고
	q := []int{2, 3, 5}
	fmt.Println(q)

	r := []bool{true, true}
	fmt.Println(r)

	//구조체를 담은 배열
	ss := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
	}

	fmt.Println(ss)
}
