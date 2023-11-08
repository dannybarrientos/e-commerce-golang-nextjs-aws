package variables

import "fmt"

func ShowNumber() {
	// Ejemplo por Declaracion
	var intcomun int

	// Ejemplo por referencia
	intde32 := int32(10)
	intde64 := int64(100)
	fmt.Println("intcomun", intcomun)
	fmt.Println("intde32", intde32)
	fmt.Println("intde64", intde64)

}