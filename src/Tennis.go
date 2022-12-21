package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	court := make(chan int)
	pointsToWin, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}
	wg.Add(2)

	scorePlayer1 := 0
	scorePlayer2 := 0

	go player("1", court, pointsToWin, scorePlayer1)
	go player("2", court, pointsToWin, scorePlayer2)

	court <- 1

	wg.Wait()
}

func player(id string, court chan int, pointsToWin int, scorePlayer int) {
	defer wg.Done()

	for {
		hit, finish := <-court
		if !finish {
			return
		}

		n := rand.Intn(100)
		if n%11 < 2 {
			fmt.Println("OH! Player ", id, " Missed :(\n")
		} else {
			scorePlayer++
			hit++
			fmt.Println("NICE! Player ", id, " Hit ", hit, " [ Score: ", scorePlayer, "]\n")
			if scorePlayer == pointsToWin {
				fmt.Printf("Player %s Won\n", id)
				close(court)
				return
			}
		}
		court <- hit
	}
}
