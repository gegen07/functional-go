package basic

func Zip() string {
	return `
func Zip<FINPUT_TYPE1><FINPUT_TYPE2>(list1 []<INPUT_TYPE1>, list2 []<INPUT_TYPE2>) map[<INPUT_TYPE1>]<INPUT_TYPE2> {
	newMap := make(map[<INPUT_TYPE1>]<INPUT_TYPE2>)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}
`
}
