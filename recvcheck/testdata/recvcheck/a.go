package recvcheck

import "fmt"

type MyStruct struct {
	a int
}

func (m MyStruct) Add(b int) int {
	return m.a + b
}

func (m *MyStruct) Sub(b int) int {
	return m.a - b
}

func Print(m MyStruct) string {
	return fmt.Sprintf("%d", m.a)
}

func (n MyStruct) Divide(b int) int { // want `invalid receiver name "n" of type "MyStruct" for method "Divide", should be "m"`
	return n.a / b
}