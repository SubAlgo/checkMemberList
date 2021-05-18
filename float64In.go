package checkMemberList

func Float64In(sli []float64, v float64) bool {
	for _, s := range sli {
		if s == v {
			return true
		}
	}
	return false
}
