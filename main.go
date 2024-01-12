package main

import (
	"fmt"

	"github.com/alexohneander/xs2a/pkg/xs2a"
)

func main() {
	// Define new xs2a API client
	client := xs2a.NewClient("https://xs2a.tech26.de/sandbox", "PSDDE-BAFIN-000001", "w6uP8Tcg6K2QR905Rms8iXTlksL6OD1KOWBxTK7wxPI", 6000000000)

	// Authorize client
	location, err := client.Authorize()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(location)
}
