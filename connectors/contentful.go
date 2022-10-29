package connectors

import (
	"fmt"
)

type Contentful struct {
	name string
}

func (c Contentful) Setup() {
	fmt.Println(c.name)
}
