package main

import (
	"fmt"
	"strings"

	"github.com/kniren/gota/dataframe"
)

func main() {
	fmt.Println("HELLO")

	jsonStr := `[{"COL.2":1,"COL.3":3},{"COL.1":5,"COL.2":2,"COL.3":2},{"COL.1":6,"COL.2":3,"COL.3":1}]`
	df := dataframe.ReadJSON(strings.NewReader(jsonStr))

	fmt.Println(df)
}
