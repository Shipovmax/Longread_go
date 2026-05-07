package atourofgo

import "fmt"

// do — функция, которая принимает любой тип (interface{})
func do(i interface{}) {
	// type switch — специальный switch, который проверяет реальный тип значения
	switch v := i.(type) {

	case int: // если внутри лежит int
		fmt.Printf("Twice %v is %v\n", v, v*2)

	case string: // если внутри лежит string
		fmt.Printf("%q is %v bytes long\n", v, len(v))

	default: // если тип не подошёл ни под один case
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func type_switches() {
	do(21)      // int
	do("hello") // string
	do(true)    // bool → default
}
