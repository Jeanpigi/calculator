package calculator

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//Calc es un struct importante
type Calc struct{}

func (Calc) parseString(operator string) (int, error) {
	result, err := strconv.Atoi(operator)
	return result, err
}

//Operate es una funcion importante
func (c Calc) Operate(input string, operation string) (int, error) {
	cleanInput := strings.Split(input, operation)
	first, err := c.parseString(cleanInput[0])
	if err != nil {
		return 0, err
	}
	second, err := c.parseString(cleanInput[1])
	if err != nil {
		return 0, err
	}

	switch operation {
	case "+":
		return first + second, nil
	case "-":
		return first - second, nil
	case "*":
		return first * second, nil
	case "/":
		return first / second, nil
	default:
		log.Println(operation, "operation is not supported!")
		return 0, nil

	}

}

func main() {
	fmt.Println("Enter your input with operation and without spaces, example 23+23")
	input := ReadInput()
	fmt.Println("Enter your operation [+, -, *, /]")
	operator := ReadInput()
	processResult(input, operator)

}

//ReadInput funcion importante para exportar
func ReadInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func processResult(input string, operator string) {
	c := Calc{}
	value, err := c.Operate(input, operator)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result of", input, "equals to", value)
	}
}
