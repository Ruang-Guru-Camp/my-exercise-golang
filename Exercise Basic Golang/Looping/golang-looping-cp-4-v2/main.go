package main

import (
	"fmt"
	"strings"
)

func EmailInfo(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return "Invalid email format"
	}

	domainParts := strings.Split(parts[1], ".")
	if len(domainParts) < 2 {
		return "Invalid email format"
	}

	domain := domainParts[0]
	var tld string
	if len(domainParts) > 2 {
		tld = domainParts[len(domainParts)-2] + "." + domainParts[len(domainParts)-1]
	} else {
		tld = domainParts[len(domainParts)-1]
	}

	return fmt.Sprintf("Domain: %s dan TLD: %s", domain, tld)
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(EmailInfo("admin@yahoo.co.id"))
}
