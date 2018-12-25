package testmod

import "fmt"

func Sum(a int, b int, name string) string {
	return fmt.Sprintf("for %s sum is %d !!!", name, a+b)
}
