package main

import (
	"fmt"
	"github.com/rvanwhy/processor_race/racer"
)
//A simple example of the use of the racers
func main() {
	c := make(chan racer.Status)
	for i := 1; i < 11; i++ {
		a := racer.Racer{fmt.Sprintf("Randy%d", i), i, 100}
		go a.Race(c)
	}

	b := <-c
	for ; b.Status != -1; b = <-c {
		fmt.Printf("dude: %s says he is %d\n", b.Name, b.Status)
	}

	fmt.Printf("Bro %s won!\n", b.Name)
}
