package leetcode

import (
	"fmt"
	"testing"
)

func Test123(t *testing.T) {
	emails := []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}
	fmt.Println(numUniqueEmails(emails))
}
