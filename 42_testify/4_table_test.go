package main

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"unicode"
)

// IsAllLowerCase 是否input所有都是大写
func IsAllLowerCase(input string) (bool, error) {
	re, err := regexp.Compile(`[0-9]+`)
	if err != nil {
		return false, err
	}

	loc := re.FindStringIndex(input) // 返回最左边和最右边匹配到到配置
	if len(loc) == 0 {
		return false, nil
	}

	for _, ch := range input {
		if unicode.IsUpper(ch) { // 判断是否是小写 判断是否有小写，可以转成是否有大写
			return false, nil
		}
	}

	return true, nil
}

func TestIsAllLowerCase(t *testing.T) {
	// define test case struct
	type testCase struct {
		Name       string
		Input      string
		ExpectData bool
		ExpectErr  bool
	}

	testTable := []testCase{
		{
			Name:       "all lower case",
			Input:      "hello",
			ExpectData: true,
			ExpectErr:  false,
		},
		{
			Name:       "upper case",
			Input:      "Hello",
			ExpectData: false,
			ExpectErr:  false,
		},
		{
			Name:       "input string contains number",
			Input:      "hello123",
			ExpectData: false,
			ExpectErr:  false,
		},
	}

	for _, test := range testTable {
		lowerCase, err := IsAllLowerCase(test.Input)
		if err != nil {
			return
		}
		assert.Equal(t, test.ExpectData, lowerCase, test.Name)
	}

}
