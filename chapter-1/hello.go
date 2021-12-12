package main // this is package declare to run program as each code relates to package

import "fmt"
import "os" // these are import statements of libraries that aere reuqired. these are standard, you can download as per requirememt

func main() { // this is main method same as other lang. responsible to run the code/starting point of you app
	fmt.Println("Hello, gopher!") // print string to output
	os.Exit(0)                    // exit normally optional
}
