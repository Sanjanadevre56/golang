package main

import "fmt"

const Sanjana string = "ghabbhhjd" //this is public bcoz of first latter is a capital latter
//this Sanjana is accesible by any other file into this folder or acually in this program and you can use this anywhere

func main() {
	var username string = "sanjana"
	fmt.Println(username)
	fmt.Printf("variable is of type : %T \n", username)

	var isloggedIn bool = true
	fmt.Println(isloggedIn)
	fmt.Printf("variable is of type : %T \n", isloggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type : %T \n", smallVal)

	var smallFloat float32 = 255.0
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type : %T \n", smallFloat)

	// default values an some aliases
	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("variable is of type : %T \n", anothervariable)

	//imlicit type
	var website = "learncodeonlinenin"
	fmt.Println(website)

	// no var style
	numberOfUser := 300000.0
	fmt.Println(numberOfUser)

	fmt.Println(Sanjana)
	fmt.Printf("variable is of type : %T \n", Sanjana)
}
