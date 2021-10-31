package main

import (
	"fmt"
	"log"
)

func main() {
	variations, err := tetris("test.go")

	if err != nil {
		log.Printf("main: %s\n", err)
	}

	for _, s := range variations {
		for _, l := range s {
			fmt.Println(l)
		}
		fmt.Println("////////////////////")
	}
	fmt.Println(topWords("00000aASDASD0099,   HELOO, ,darkness &old&  old heloo   0000 ddd ,  0doaoowkj,  owkpa. [pasdp", 100))
}
