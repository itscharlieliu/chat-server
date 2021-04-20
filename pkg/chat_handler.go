package pkg

import "fmt"

func ChatHandler(c chan []byte) {
	for {
		str := <-c
		fmt.Println(string(str))
		fmt.Println("End msg")
	}
}
