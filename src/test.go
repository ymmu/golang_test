package main // main함수로 인식할 수 있도록 함

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

//func add(x int, y int) int {
func add(x, y int) int { //같은 매개변수 타입일 때 변수들 모두 써주고 type을 뒤에 붙여줌
	return x + y
}

func swqp(x, y string) (string, string) {
	return y, x
}

func split(sum float32) (x, y float32) { //리턴값 이름 정할 수 있음
	x = sum * 4 / 9
	y = sum - x
	return //이름이 주어진 리턴 값을 반환. naked return이라 함
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v //v scope은 if문 안에서만
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim //v 넣으면 에러남
}

func yourOs() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

//var c, python, java bool //bool 타입 디폴트 초깃값 false

func main() {

	fmt.Println("Hello Golang")
	fmt.Println("The time is", time.Now())
	fmt.Println("radomn num is ", rand.Intn((10)))
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))
	//fmt.Println(math.pi) //export 되지 않음
	fmt.Println(math.Pi) //대문자인 애들은 패키지에서 export됨

	fmt.Println(add(2, 3))
	a, b := swqp("Lucca", "Hello")
	fmt.Println(a, b)
	fmt.Println(split(10))

	//var i int //숫자 타입은 자동 초깃값=0

	var c, python, java = true, false, "No!"
	var i, j int = 1, 2 //초깃값 넣기
	var str string      //문자열 타입 초깃값 ""
	k := 3              // 암시적 type으로 var 선언처럼 사용될 수 있음. 함수 밖에서는 사용할 수 없음

	fmt.Println(str, k, i, j, c, python, java)

	// 타입 변환
	i_ := 42
	f_ := float64(i_)
	//u_ := uint(f_)
	fmt.Printf("%T:", f_)

	//암시적 타입이면 값의 정확도에 따라 type이 결정됨
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)

	const Pi = 3.14
	fmt.Println("Hello", Pi)

	//shift
	fmt.Println((1 << 3)) // 2^3 = 100_{2}

	//for문
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//for문에서 초기화,사후 구문은 필수가 아님
	sum_2 := 1
	for sum_2 < 10 { //format에서 ;가 지워져서 while문이 됩니다만..
		sum_2 += sum_2
	}
	fmt.Println(sum_2)

	//;생략하면 for==while
	//0하면 무한루프.
	for sum_2 < 10 {
		sum_2 += sum_2
	}
	fmt.Println(sum_2)

	//for{} //무한루프

	fmt.Println(sqrt(2), sqrt(-4))

	//if문에서 변수 사용할 때 sciope 확인
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	yourOs()

}
