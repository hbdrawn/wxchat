package utils_test

import (
	"testing"
	. "wxchat/utils"
)

func TestCompareStr(t *testing.T){
	str1 := "ba"
	str2 := "ba"
	result := CompareStr(str1,str2)
	t.Logf("the result is: %d",result)
}

