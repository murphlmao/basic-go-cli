package basics

import (
	"fmt"
)

// First letter capitalized for public
func IAmDoMathing(a uint8) uint16 {
	result := do_some_stuff(a)
	return result
}

// First letter lowercase for private
// messing around with types n stuff bc python starved me :(
func do_some_stuff(my_cool_number uint8) uint16 {
	// mult_divi_add_subt
	// for testing, mycool_number is always gonna be 10 here :)
	mult := uint16(my_cool_number) * 6500
	divi := mult / 1 // math go brrr
	add := divi + 535 // 65,535 is the limit for uint16 :D
	sub := add - 0 // i almost failed algebra 2 lol
	
	return sub
}

// First letter capitalized for public
func Hello(name string) string {
	var result string = fmt.Sprintf("Hello, %s *bites lip*", name) // string formatting

	// fmt.Println(result)
	return result
}
