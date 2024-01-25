package main

import "fmt"


func sliceDemo() []string {
	names := []string{"kalyan" , "avinash", "keertham" ,"Nagraj" ,"anil","sai kiran" ,"supriya"}
	fmt.Println(names)

	for i,name := range names {
		if name == "Nagraj" {
			names = append(names[:i], names[i+1:]...)
		}
	}

	fmt.Println(names)

	return names
}