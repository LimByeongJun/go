package main

import (
	"fmt"
	"math/rand"
	"time"
)

//난수 추출된 수의 소수 판정 프로그램 v0.6
//소수 : 1과 자기 자신외에는 나누어 떨어지지 않는 수 0과 1은 제외
func main() {
	//seed 설정
	seed := time.Now().Unix()
	rand.Seed(seed)

	isPrime := true

	number := rand.Intn(150) + 2 //0과 1제외 2~151 사이의 수
	fmt.Println("임의로 추출된 수 : ", number)

	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
			break
		}
		//fmt.Print(i, " ")
	}

	if isPrime { //비교 연산자 제거
		fmt.Println(number, "는(은) 소수입니다!")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다.")
	}
}
