package main

import (
	"fmt"

	"edpasenidis.tech/cypher/internal"
)

func main() {
	login := internal.Credentials(internal.Login())
	internal.Room(login.Domain, login.Token)
	fmt.Println("Thanks for using Cypher! ~ Cypher Dev Team")
}
