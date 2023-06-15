/*
	Justin L. CSC 391

	Darts PI Program

*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Function that takes a channel and amount thrown
Returns the proportion of darts thrown in the 1st quad. in the unit square

@Params
ch : channel (float64) - The channel that will be fed the value proportion of the program.
numOfDartsToThrow : int - Amount of darts that is prompt to be thrown.
*/
func estimatePi(ch chan float64, numOfDartsToThrow int) {

	var numOfDartsThrown int = 0
	var numOfDartsInside int = 0

	for numOfDartsToThrow > numOfDartsThrown {
		// Throwing a dart to the upper right quadrant of the unit square
		var x float64 = rand.Float64()
		var y float64 = rand.Float64()

		// Det if it is inside the circle
		if x*x+y*y < 1.0 {
			numOfDartsInside++
		}
		numOfDartsThrown++
	}

	var insideProportion float64 = float64(numOfDartsInside) / float64(numOfDartsToThrow)

	ch <- insideProportion * 4
}

func main() {

	/*
		Iteration 3 with channels
		We are testing the change in time with 1,000,000 darts being thrown per thread.
	*/

	ch := make(chan float64)
	var start time.Time = time.Now()

	// i = the amount of threads we are using
	for i := 0; i < 10; i++ {
		// Call a go routine, passing in the channel
		go estimatePi(ch, 100000)
		go estimatePi(ch, 100000)
		go estimatePi(ch, 100000)
		go estimatePi(ch, 100000)
		go estimatePi(ch, 100000)

		go estimatePi(ch, 100000)
		go estimatePi(ch, 100000)
		go estimatePi(ch, 100000)
		go estimatePi(ch, 100000)
		go estimatePi(ch, 100000)

	}

	result := 0.0 // Feeds from Channel

	for i := 0; i < 10; i++ {
		// Read each value out of the channel
		// that was placed there in the calls to estimatePi()
		result += <-ch
	}
	var elapsed2 time.Duration = time.Since(start)
	fmt.Println("time taken:", elapsed2)
	fmt.Println("Test Pi:", result/10)

	/* // Iteration 2
	start := time.Now()
	pi := math.Pi
	fmt.Println("10,000:", (estimatePi(10000)/pi)*100, "took:", time.Since(start))
	fmt.Println("100,000:", (estimatePi(100000)/pi)*100, "took:", time.Since(start))
	fmt.Println("1,000,000:", (estimatePi(1000000)/pi)*100, "took:", time.Since(start))
	fmt.Println("10,000,000:", (estimatePi(10000000)/pi)*100, "took:", time.Since(start))
	fmt.Println("100,000,000:", (estimatePi(100000000)/pi)*100, "took:", time.Since(start))
	fmt.Println("1,000,000,000:", (estimatePi(1000000000)/pi)*100, "took:", time.Since(start))
	fmt.Println("10,000,000,000:", (estimatePi(10000000000)/pi)*100, "took:", time.Since(start))
	*/
}
