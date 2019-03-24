package inspector

import (
	"fmt"
	"github.com/napsy/go-css"
	"io/ioutil"
)

func InspectFile(path string) {

	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	text := string(b)

	stylesheet, err := css.Unmarshal([]byte(text))
	if err != nil {
		panic(err)
	}

	for _, b := range stylesheet {
		//rule := k
		for key, value := range b {
			keyword := key + ": " + value
			replacement := SearchRule(keyword)
			fmt.Println(replacement)
		}
	}
}
