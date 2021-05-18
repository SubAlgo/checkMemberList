package checkMemberList

func Float32In(sli []float32, v float32) bool {
	for _, s := range sli {
		if s == v {
			return true
		}
	}
	return false
}
