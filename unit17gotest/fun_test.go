package unit17gotest

import "testing"

func TestAddFunc01(t *testing.T) {
	res := AddFun(1, 2)
	// if res != 3 {
	// 	t.Error("测试结果错误")

	// }
	if res != 3 {
		t.Errorf("测试结果错误，实际结果为：%v", res)

	}
}
func TestAddFunc02(t *testing.T) {
	res := AddFun(1, 2)
	// if res != 3 {
	// 	t.Error("测试结果错误")

	// }
	if res != 2 {
		t.Errorf("测试结果错误，实际结果为：%v", res)

	}
}
func TestAddFunc03(t *testing.T) {
	var tests = []struct {
		inputX  int64
		inputY  int64
		outputZ int64
	}{
		{2, 3, 5},
		{7, 8, 15},
		{2, 4, 6},
		{3, 3, 6},
		{1, 8, 9},
	}
	for _, test := range tests {
		if res := AddFun(test.inputX, test.inputY); res != test.outputZ {
			t.Errorf("AddFun测试失败，实际结果为：%v,预期结果为：%v", res, test.outputZ)

		}
	}
}
