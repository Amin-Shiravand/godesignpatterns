package main

import "fmt"

func main() {

	var emailbuilder = newEmailBuilder()
	emailbuilder.SetTitle("Hiring")
	emailbuilder.SetAddress("UK")
	emailbuilder.SetDescription("")
	emailbuilder.SetAttachments(make([]string, 0))
	email, err := emailbuilder.Build()

	if err != nil {
		fmt.Println("Error creating email", err)
	} else {
		fmt.Printf("Email: %+v", email)
	}
}
