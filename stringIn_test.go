package checkMemberList

import "testing"

func TestStringIn(t *testing.T) {
	var testSlice = []string{"test1", "test2", "test3"}
	var testStr = []string{"test1", "test2", "test4"}

	for _, s := range testStr {
		var res bool
		res = StringIn(testSlice, s)
		if s == "test1" && res == false {
			t.Errorf("string value %s is in slice %v should return true but return %v ", s, testSlice, res)
		}

		if s == "test2" && res == false {
			t.Errorf("string value %s is in slice %v should return true but return %v ", s, testSlice, res)
		}

		if s == "test4" && res == true {
			t.Errorf("string value %s is not in slice %v should return false but return %v ", s, testSlice, res)
		}
	}
}
