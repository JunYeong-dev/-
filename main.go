package main

import (
	"fmt"
	"time"
)

func main() {
	// goroutines - 동시 작업을 가능하게 해주는 아주 awesome 한 기능
	go printPerson("nick")
	go printPerson("judy")
	time.Sleep(time.Second * 10)

	// main function은 goroutines를 기다려 주지 않기 때문에
	// 이런 식으로 모든 것을 goroutines에 올려놓게 되면 main function은
	// 따로 할 작업이 없기 때문에 종료되고 goroutines도 실행되지 않는다
	// 그렇기 때문에 다른 작업을 통해 실행될 수 있도록 함
	// 이 예제에서는 Sleep로 main function의 작업을 부여했지만 실제 개발에서는 이렇게 쓰지 않음
	// go printPerson("nick")
	// go printPerson("judy")
}

func printPerson(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, " : ", i)
		time.Sleep(time.Second)
	}
}
