package main

import (
	"fmt"
)

func superAdd(numbers ...int) int {
	for index := range numbers {
		fmt.Println(numbers[index])
	}
	return 1
}
func main() {
	superAdd(1, 2, 3, 4, 5, 6)
}

/*
	main package만 compile 가능
	import를 이용해 package를 포함
	return 값으로 여러 타입 가능
	defer을 이용해 func 종료 후 명령 지정 가능
	loop는 오로지 'for'로 표현 가능 (fuc에서 쓰는 게 가장 best)
	like Python -> range 변수 형태로 사용
*/
