package pkg

import "fmt"

func ChatHandler(c chan []byte) {
	for {
		str := <-c
		fmt.Println(str)
		fmt.Println("End msg")
	}
}
