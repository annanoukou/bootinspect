package inspector

import (
	"fmt"
	"github.com/napsy/go-css"
)

func DdoThis() {
	ex1 := `.foo {
	font-weight: bold;
	height: 100%;
	}`

	stylesheet, err := css.Unmarshal([]byte(ex1))
	if err != nil {
		panic(err)
	}

	for _, b := range stylesheet {
		//rule := k
		for key, value := range b {
			keyword := key + ":" + value
			replacement := SearchRule(keyword)
			fmt.Println(replacement)
		}
	}
}
