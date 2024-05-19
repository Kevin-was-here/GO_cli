package main

import "fmt"

//commands to implement: https://aws.amazon.com/what-is/cli/#seo-faq-pairs#what-are-some-common-command-line-interface-commands

func main() {
	lock := 1

	var user_input string

	for lock <= 5 {
		fmt.Print("Enter a command: ")
		fmt.Scan(&user_input)
		fmt.Println("Executing command:", user_input)
		run_command(user_input)
	}

}

func run_command(cmd string){
	
}
