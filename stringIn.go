package checkMemberList

func StringIn(sli []string, s string) bool {
	for _, sl := range sli {
		if sl == s {
			return true
		}
	}
	return false
}
