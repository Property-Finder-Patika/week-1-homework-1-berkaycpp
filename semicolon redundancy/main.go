// unnecessary semicolons. after format, they would dissappear
/*
package main
import "fmt";
func main() {
 var aString string = "Selam";
 fmt.Printf("%s", aString);
 anInteger := 5;
 fmt.Printf("%d", anInteger);
 return;
}
*/
package main

import "fmt"

func main() {
	var aString string = "Selam"
	fmt.Printf("%s", aString)
	anInteger := 5
	fmt.Printf("%d", anInteger)
	return
}
