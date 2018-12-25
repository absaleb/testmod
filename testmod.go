package testmod

import "fmt"

func Sum(a int, b int, name string) string {
	return fmt.Sprintf("for %s sum of %d and %d is %d !!!", name, a, b, a+b)
}
