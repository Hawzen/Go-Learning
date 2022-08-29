package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Glass())
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
	fmt.Println(sssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss)

	fmt.Println()
	fmt.Println(math.Inf(1), math.Inf(-1), math.Inf(0), Pi)
	fmt.Println(math.SmallestNonzeroFloat64, strconv.FormatUint(math.MaxUint64, 10))
	fmt.Println(math.Sqrt(-1), math.Sqrt(math.Inf(-1)))

	fmt.Println()
	time_now := time.Now().Format(time.RFC822)
	fmt.Println(time_now)
	duration, _ := time.ParseDuration(time_now)
	rand.Seed(duration.Nanoseconds())
	fmt.Println(rand.Int63())

	fmt.Print("\n\r")
	fmt.Println(add_and_do_undefined_weird_buggy_behaviour_that_will_break_your_program_now_i_am_stalling_for_words_to_see_how_rediciouosdlsd_I_can_take_this(6, 9))
}

func add(x, y int) int {
	return x + y
}

func add_and_do_undefined_weird_buggy_behaviour_that_will_break_your_program_now_i_am_stalling_for_words_to_see_how_rediciouosdlsd_I_can_take_this(x, y int) (int, int) {
	// The best function
	return int(math.Mod(float64(x), float64(float32(y)))), 240
}
