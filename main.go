//====== Standard go main file for Golang apps ==============================
//........ By Atchom Bena gho Benaloga nipie Aloganjoh piar Ndou´makong........
//........ 15/03/2022 , KindingNde, Cameroon .......................
//============================================================


//package definition
package main

//external librairies
import "fmt"

//variables declaration
var version = "dev"
const TEST_CONST_00001 = "Hello M. Bena welcome"

//function call in main go application process
func main() {

	fmt.Printf("\nVersion: %s\n", version)
	fmt.Printf(hello())
	fmt.Printf("\n\n")
}

//function declaration

func hello() string {
	return TEST_CONST_00001  
}
