package word1

import "testing"

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated") = false`)
	}

	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") = false`)
	}
}

func TestNonPalindrome(t *testing.T)  {
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrome") = true`)
	}
}

/*
# go test
PASS
ok      gopl/11.Testing/word1    0.295s
*/

func TestFrenchPalindrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Error(`IsPalindrome("été") = false`)
	}
}

func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrome(%q) = false`, input)
	}
}

/*
# go test
--- FAIL: TestFrenchPalindrome (0.00s)
    word_test.go:29: IsPalindrome("été") = false
--- FAIL: TestCanalPalindrome (0.00s)
    word_test.go:36: IsPalindrome("A man, a plan, a canal: Panama") = false
FAIL
exit status 1
FAIL    11.Testing/word1    0.275s

选项 -v 可以输出包中每个测试用例的名称和执行时间

# go test -v
=== RUN   TestIsPalindrome
--- PASS: TestIsPalindrome (0.00s)
=== RUN   TestNonPalindrome
--- PASS: TestNonPalindrome (0.00s)
=== RUN   TestFrenchPalindrome
    TestFrenchPalindrome: word_test.go:29: IsPalindrome("été") = false
--- FAIL: TestFrenchPalindrome (0.00s)
=== RUN   TestCanalPalindrome
    TestCanalPalindrome: word_test.go:36: IsPalindrome("A man, a plan, a canal: Panama") = false
--- FAIL: TestCanalPalindrome (0.00s)
FAIL
exit status 1
FAIL    11.Testing/word1    0.226s

-run 的参数是一个正则表达式，它可以使得 go test 只运行那些测试函数名称匹配给定模式的函数

# go test -v -run="French|Cannal"
=== RUN   TestFrenchPalindrome
    TestFrenchPalindrome: word_test.go:29: IsPalindrome("été") = false
--- FAIL: TestFrenchPalindrome (0.00s)
FAIL
exit status 1
FAIL    11.Testing/word1    0.283s

*/