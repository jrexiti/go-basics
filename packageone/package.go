package packageone

import "fmt"

var privateVar = "I am private"
var PublicVar = "I am public"

func notExported() {
	fmt.Println("from func notExported")
}

func Exported() {

	fmt.Println("from func Exported")

}
