package connectors

import (
	"fmt"
)

type Crystallize struct {
	name string
}

func (c Crystallize) Setup() {
	fmt.Println(c.name)
}
