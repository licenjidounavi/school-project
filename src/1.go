
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(randomGoCode())
}

func randomGoCode() string {
	return fmt.Sprintf("var %s = func() {\n\t%s := rand.Intn(%d)\n\treturn %s\n}", generateVariableName(), generateFunctionName(), 10, generateVariableName())
}

func generateVariableName() string {
	return "x" + strconv.Itoa(rand.Intn(10))
}

func generateFunctionName() string {
	return "f" + strconv.Itoa(rand.Intn(10))
}