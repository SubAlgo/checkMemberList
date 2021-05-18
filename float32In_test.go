package checkMemberList

import "testing"

func TestFloat32In(t *testing.T) {
	var testSlice = []float32{1.1, 2.2, 6.3}
	var testInt = []float32{1.1, 2.2, 5.3}

	for _, i := range testInt {
		var res bool
		res = Float32In(testSlice, i)

		// error case
		if i == 1.1 && res == false {
			t.Errorf("value %v is in slice %v should be return true but return %v ", i, testSlice, res)
		}

		// error case
		if i == 2.2 && res == false {
			t.Errorf("value %v is in slice %v should be return true but return %v ", i, testSlice, res)
		}

		// error case
		if i == 5.3 && res == true {
			t.Errorf("value %v is not in slice %v should be return false but return %v ", i, testSlice, res)
		}
	}
}
