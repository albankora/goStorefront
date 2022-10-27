package bridge

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

func (c Contentful) Blob() {
	fmt.Println(c.name)
}
