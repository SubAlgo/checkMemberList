package checkMemberList

func Int64In(sli []int64, i int64) bool {
	for _, ii := range sli {
		if ii == i {
			return true
		}
	}
	return false
}
