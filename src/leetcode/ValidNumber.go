package leetcode
//https://leetcode.com/problems/valid-number/description/
/**Some examples:
"0" => true
" 0.1 " => true
"abc" => false
"1 a" => false
"2e10" => true
**/
import (
	"strings",
	"fmt",
	"strconv"
)
func isNumber(s string) bool {
    s = strings.TrimSpace(s)
	if(s[0] == '0' && len(s) == 1) {
		return true
	}
	notZeroIndex := 0
	for _,v := range s{
		if(v == "0") {
			notZeroIndex = notZeroIndex+1
		}
	}
	s = s[notZeroIndex:len(s)]
	if(len(s) == 0) {
		return false
	}
	var numberList = []byte{'0','1','2','3','4','5','6','7','8','9'}
	if strings.Contains(s,".") {
		if(s[0] != '0' && s[0] != '.') {
			return false
		} else {
			s = s[1:len(s)]
		}
	} else {
		if(s[0] == '0' && len(s) == 1) {
			return true
		} else if(s[0] == '0') {
			return false
		}
	}
	isNum := false
	count := 0
	for i:= 0; i < len(s); i++ {
		for _,value :=range numberList {
			if s[i] != value {
				continue
			} else {
				isNum = true
				count ++
			}
		}
	}
	return isNum && count == len(s)
}