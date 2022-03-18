//====== Standard go main_test file for CD/CI ==============================
//........ By Atchom Bena gho Benaloga nipie Aloganjoh piar Ndou´makong........
//........ 15/03/2022 , KindingNde, Cameroon .......................
//============================================================

//package definition
package main

//external librairies
import "testing"

//variables declaration
//var version = "dev"
//const TEST_CONST_00001 = "Hello_M._Bena_welcome_to_Go_main"

//Test function declaration for theHello main app defined in main.go

func TestHello(t *testing.T) {
	
	//wanted output result as defined in  main.go
	want := "Hello M. Bena welcome2" 

	//function to be tested
	got := hello()

	//processing the test
	if want != got {
		t.Fatalf("want %s, got %s\n", want, got)
	}
}
