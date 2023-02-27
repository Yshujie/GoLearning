package stringType

import "fmt"

// TestQuotes 测试单引号、双引号
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

// TextPrintEnglishByIndex 测试根据字符序号打印英文字符串
func TextPrintEnglishByIndex() {
	str := "Hello World"

	println(str[0])
	println(str[1])
	println(str[2])
	println(str[3])
	println(str[4])
	println(str[5])
	println(str[6])
	println(str[7])
	println(str[8])
	println(str[9])

}

// TextPrintChinesByIndex 测试根据字符序号打印中文字符串
func TextPrintChinesByIndex() {
	str := "十里春风不如你"

	println(str[0])
	println(str[1])
	println(str[2])
	println(str[3])
	println(str[4])
	println(str[5])
	println(str[6])
	println(str[7])
	println(str[8])
	println(str[9])
	println(str[10])
	println(str[11])
	println(str[12])
	println(str[13])
	println(str[14])
	println(str[15])
}

// TextTraversalEnglishString 遍历英文字符串
func TextTraversalEnglishString() {
	str := "only god know, how much I love you"

	for i := 0; i < len(str); i++ {
		fmt.Printf("%v(%c)", str[i], str[i])
		fmt.Println()
	}
}

// TextTraversalChinesString 遍历中文字符串
func TextTraversalChinesString() {
	str := "只有上帝知道，我有多爱你"

	for _, r := range str {
		fmt.Printf("%v(%c)", r, r)
		fmt.Println()
	}
}

func printTest(x1, xx1 rune, y1, yy1 string) {
	fmt.Println(x1)
	fmt.Println(xx1)
	fmt.Println(y1)
	fmt.Println(yy1)
}
