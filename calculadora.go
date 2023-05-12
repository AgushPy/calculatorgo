package calculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calc struct {
}

func (Calc) Operate(entrada string, operador string) (rdo int) {
	entradaLimpia := strings.Split(entrada, operador)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])
	switch operador {
	case "+":
		rdo := operador1 + operador2
		fmt.Println(rdo)
		return rdo
	case "-":
		rdo := operador1 - operador2
		fmt.Println(rdo)
		return rdo
	case "*":
		rdo := operador1 * operador2
		fmt.Println(rdo)
		return rdo
	case "/":
		rdo := operador1 / operador2
		fmt.Println(rdo)
		return rdo
	default:
		fmt.Println("El tipo de operación ingresada es inexsitente")
		return 0
	}
}

func parsear(entrada string) (operador int) {
	operador, _ = strconv.Atoi(entrada)
	return
}

func LeerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	entrada := leerEntrada()
	operador := leerEntrada()
	c := calc{}
	c.operate(entrada, operador)
}
