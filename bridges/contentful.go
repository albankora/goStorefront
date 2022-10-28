package bridges

import (
	"fmt"
)

type Contentful struct {
	name string
}

func (c Contentful) Setup() {
	fmt.Println(c.name)
}

func (c Contentful) HomePage() {
	fmt.Println(c.name)
}

func (c Contentful) Blog() {
	fmt.Println(c.name)
}
