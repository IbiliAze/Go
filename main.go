package main

import (
	"fmt"
	"math"
)

func main() { //comments
	fmt.Println(5 + 7*2)
	fmt.Println("hello clo\nud gurus")
	fmt.Println('D')
	fmt.Println(33.0 / 4)
	fmt.Println(math.Pow(12.3, 5))
	fmt.Println(math.Sqrt(25))
	fmt.Println(math.Sqrt(26))
	fmt.Println(2 > 5)
	fmt.Println(2 == 5)
	fmt.Println('D' == 68)
	fmt.Println(math.Logb(32))
	var testVariable string = "okaay"
	fmt.Println(testVariable)
	var number1 = 20
	var var1, var2, var3 = "ada", true, 10
	fmt.Println(var1, var2, number1/var3)
	var4 := "value"
	fmt.Println(var4)
	var1 = "baba"
	names := [4]string{"a", "b", "c", "d"}
	fmt.Println(names)
	var names2 [2]string
	names2[0] = "a"
	names2[1] = "b"
	fmt.Println(names2)
	fmt.Println(names2[1])
	names3 := []string{}
	names3 = append(names3, "aa")
	names3 = append(names3, "ab")
	names3 = append(names3, "bb")
	fmt.Println(names3)
	mapss := map[string]string{
		"aa": "1212",
		"bb": "3443",
		"cc": "5445",
	}
	fmt.Println(mapss)
	delete(mapss, "bb")
	fmt.Println(mapss)
	if mapss["aa"] == "1212" {
		fmt.Println("trueeee")
	} else {
		fmt.Println("falseeee")
	}
	for i := 1; 1 <= 10; i++ {
		fmt.Println("We are counting", i)
	}
}
