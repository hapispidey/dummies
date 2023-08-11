package main

import "fmt"


func main(){
	var n string = "Nay Lin Aung"
	var a int = 18
	fmt.Println("Hello, world!")
	greeting(n, a)
	agetest(a)
}

func greeting(name string, age int) {
	fmt.Printf("I'm %s, I am %d years old.\n", name, age)
}

func agetest(age int) {
	if age >= 18 && age < 110 {
		fmt.Println("You are old enough to watch hentai.")
	}else if age < 18 && age > 0 {
		fmt.Println("WTF are you trying to watch.")
	}else {
		fmt.Println("LoL, You aren't even a human being.")
	}
}
