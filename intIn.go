package checkMemberList

func IntIn(sli []int, i int) bool {
	for _, ii := range sli {
		if ii == i {
			return true
		}
	}
	return false
}
