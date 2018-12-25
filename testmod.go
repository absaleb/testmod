package testmod

import "fmt"

func Sum(a int, b int, c int, name string) string {
	return fmt.Sprintf("for %s sum of %d and %d and %c is %d !!!", name, a, b, c, a+b+c)
}
