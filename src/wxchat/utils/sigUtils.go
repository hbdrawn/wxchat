package utils

//import "fmt"


//当且仅当str1>=str2(自然排序)时返回1
func CompareStr(str1, str2 string) int {
	len1 := len(str1)
	len2 := len(str2)
	if ok := len1 - len2; ok > 0 {
		len1 = len2
	}

	for i := 0; i < len1; i++ {
		if str1[i] > str2[i] {
			return 1
		}
		if str1[i] < str2[i]{
			return -1
		}
	}
	
	if len1==len2{
		return 1
	}
	return -1
}
