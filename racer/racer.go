//Package race contains all the resources needed
//to run a good Ol' Fashioned processor race.
package racer

import (
	"math/rand"
	"time"
)

//Racer is a man whose sole purpose in life is
//to race harder than any man has raced before.
type Racer struct {
	Name       string
	Number     int
	Race_limit int
}

//Status encapulates the state of a racer. This is used
//for reporting to other functions that may use racer.
type Status struct {
	Name   string
	Status int
}

//Race causes the racer to begin his race
//based on the established rule. The racer will
//then report the status of his race to the go channel.
//Returns -1 if the racer has passed the race_limit
func (racer Racer) Race(c chan Status) {
	sum := 0
	rand.Seed(int64(racer.Number) * time.Now().UTC().UnixNano())

	for sum < racer.Race_limit {
		time.Sleep(time.Second)
		sum += rand.Intn(15) //random int from [0-15]
		c <- Status{racer.Name, sum}
	}

	c <- Status{racer.Name, -1}
}


