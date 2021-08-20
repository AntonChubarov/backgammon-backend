
package utils

func Contains(s *[]string, e string) bool {
	for i := range *s {
		if (*s)[i] == e {
			return true
		}
	}
	return false
}

func AreEqualIntSlices( sl1 []int, sl2 []int) bool {

	l1:=len(sl1)
	l2:=len(sl2)
	if l1!=l2 {
		return false}

	for i:= range sl1 {
		if sl1[i]!=sl2[i] {
			return false}
	}
	return true
}
