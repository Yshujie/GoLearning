package stringDataType

import "fmt"

func TestQuotes() {

	var a1 = 'a'
	// var aa1 = 'aa'
	var aa1 = '9'
	var b1 = "a"
	var bb1 = "aa"

	printTest(a1, aa1, b1, bb1)

	var a2 = '我'
	// var aa2 = '我们'
	var aa2 = 'w'

	var b2 = "我"
	var bb2 = "我们"

	printTest(a2, aa2, b2, bb2)

}

func printTest(x1, xx1 rune, y1, yy1 string) {
	fmt.Println(x1)
	fmt.Println(xx1)
	fmt.Println(y1)
	fmt.Println(yy1)
}
