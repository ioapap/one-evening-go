package main

import "strings"

func CountCreatedEvents(events []string) int {
	created := 0
	for _, event := range events {
		if strings.HasSuffix(event, "_deleted") {
			break
		}
		if strings.HasSuffix(event, "_created") {
			continue
		}
		created++
		
	} 
	return created
}

func main() {
	events := []string{
		"product_created",
		"product_updated",
		"product_assigned",
		"order_created",
		"order_updated",
		"client_created",
		"client_updated",
		"client_refreshed",
		"client_deleted",
		"order_updated",
	}

	CountCreatedEvents(events)
}
