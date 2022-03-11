package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"

	"golang.org/x/tour/pic"
)

type Vertex struct {
	X, Y int
}

var (
	v1     = Vertex{1, 2}               // has type Vertex
	v2     = Vertex{X: 1}               // Y:0 is implicit
	v3     = Vertex{}                   // X:0 and Y:0
	p      = &Vertex{1, 2}              // has type *Vertex
	primes = [6]int{2, 3, 5, 7, 11, 13} //배열 크기 고정

)

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v \n", s, len(x), cap(x), x)
}

func main() {

	//포인터
	fmt.Println(v1, p, *p, p.X, (*p).X, v2, v3)

	//배열
	var a [2]string
	a[0] = "Hi"
	a[1] = "Lucca"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	//슬라이스
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

	fmt.Println("========================")
	emptys := ""
	//슬라이스 길이와 용량
	printSlice(emptys, s)

	//길이 0으로 만든 뒤 길이와 용량 확인
	//용량이란 슬라이스의 "첫 번쨰 요소"부터 계산하는 "기본 배열의 요소의 개수"
	s = s[:0] //시작을 지정 안 하면 0부터 시작하니까 카파는 기본배열 카파가 나오겠네
	printSlice(emptys, s)

	//길이 확장
	s = s[:4]
	printSlice(emptys, s) //

	s = s[2:] //기본배열의 2에서 시작해서 카파 변화 있음.
	printSlice(emptys, s)

	//s = s[:6] //에러 capa가 4인데 6내놓으라고 해서
	//printSlice(s)

	fmt.Println("========================")
	//용량 0인 배열 nil
	var null_ []int
	fmt.Println(null_, len(null_), cap(null_))
	if null_ == nil {
		fmt.Print("nil! \n")
	}

	fmt.Println("========================")
	// make 함수로 동적 크기의 배열을 생성하기
	da := make([]int, 5)
	printSlice("da", da)

	//용량 지정
	db := make([]int, 0, 5)
	printSlice("db", db)

	dc := db[:2]
	printSlice("dc", dc)

	dd := dc[2:5]
	printSlice("dd", dd)

	fmt.Println("========================")
	//슬라이스는 다른 슬라이스를 포함하여 모든 타입을 담을 수 있음
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	fmt.Println("========================")
	//원소 추가해서 원래 배열 사이즈보다 커지면 새 배열 만들어 할당함
	var ss_ []int
	printSlice(emptys, ss_)

	ss_ = append(ss_, 0)
	printSlice(emptys, ss_)

	ss_ = append(ss_, 0, 2, 3, 4)
	printSlice(emptys, ss_)
	fmt.Println("========================")
	//range는 for 요소들을 순환
	for i, v := range primes {
		fmt.Printf("%d: %d\n", i, v)
	}
	fmt.Println("========================")
	pow := make([]int, 10)
	//인덱스만 가져오고 싶을 때 두번째 인자 생략
	for i := range pow {
		pow[i] = 1 << uint(i) // 2**i
	}
	for _, v := range pow { //기본적으로 생략하려면 "_" 써줌
		fmt.Printf("%d\n", v)
	}

	fmt.Println("========================")
	// 그리그리기..배열 실습
	pic.Show(Pic)
}

func dotVal(x, y float64) uint8 {
	//fmt.Println(x, y)
	return uint8(math.Min(x*y, 255))
}
func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy)
	/* 0으로 채운
	for i := range img{
		img[i] = make([]uint8, dx)
	}
	*/
	// 랜덤함수로 채운
	for i := range img {
		img[i] = make([]uint8, dx)
		for j := range img[i] {
			//rand.Seed(time.Now().UnixNano())
			img[i][j] = dotVal(float64(rand.Intn(20)), float64(rand.Intn(30)))
		}
	}

	return img
}
