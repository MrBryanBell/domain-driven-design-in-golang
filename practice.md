package main

import "fmt"

func main() {
	{
		/* a normal execution */
		var result, error = divive(4, 2)
		if error != nil {
			fmt.Println(error)
		} else {	
			fmt.Printf("result: %d \n", result)
		}
	}

	{
		/* a normal execution */
		var result, error = divive(4, 0)
		if error != nil {
			fmt.Println(error)
		} else {	
			fmt.Printf("result: %d \n", result)
		}
	}
}

func divive(a, b int) (int, error) {
	if b == 0 {
		return fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}



<!-- MAKING OUR CODE MORE SEMANTIC -->
package User

import "fmt"

type User struct {
	name  string
	email string
}

const BLOCKED_NAME = "Alex"

func Create(name, email string) (User, error) {

	if name == BLOCKED_NAME {
		return User{}, fmt.Errorf("name cannot be %s", BLOCKED_NAME)
	}

	return User{
		name:  name,
		email: email,
	}, nil
}


package main

import (
	"fmt"
	"mymodule/User"
)

func main() {
	{
		/* Creating an invalid user */
		var user, err = User.Create("Alex Flores", "woods@my.com")

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%+v\n", user)
		}
	}

	{
		/* Creating a valid user */
		var user, err = User.Create("Bryan Bell", "woods@my.com")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%+v\n", user)
		}
	}
}
