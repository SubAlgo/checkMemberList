package checkMemberList

import "testing"

func TestInt64In(t *testing.T) {
	var testSlice = []int64{1, 2, 6}
	var testInt = []int64{1, 2, 5}

	for _, i := range testInt {
		var res bool
		res = Int64In(testSlice, i)

		// error case
		if i == 1 && res == false {
			t.Errorf("int value %v is in slice %v should be return true but return %v ", i, testSlice, res)
		}

		// error case
		if i == 2 && res == false {
			t.Errorf("int value %v is in slice %v should be return true but return %v ", i, testSlice, res)
		}

		// error case
		if i == 5 && res == true {
			t.Errorf("int value %v is not in slice %v should be return false but return %v ", i, testSlice, res)
		}
	}
}
