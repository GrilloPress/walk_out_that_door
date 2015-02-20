package main

import "fmt"

var challenger = "Good Hero Knight Guy"

var (
	// another way of defining things
	champion = "Evil Blue Dragon"
	// you can declare the type, though when it write it out GoLang will guess
	age int = 4567
	location string = "Cave"
)

func main(){
	
	// Writing this style, as we saw before also declares a variable... though this type can only ever ever be written inside a function. I think := looks like a closed pair of scissors

	winner := "Evil Blue Dragon"
	
	fmt.Printf("The very smashing %s came along, and was challenged to a battle by ... \n %s (%d years young), in the dank, deep smelly %s \n they fought hard and long \n but the winner of the battle was: ... \n %s! \n", challenger, champion, age, location, winner )
	
}