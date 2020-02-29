package main

import "testing"

func TestReplaceStringSuccess(t *testing.T) {
	t.Run("DataDriven", func(t *testing.T) {

		testCases := []struct {
			TestCase  string
			Input     string
			Expected1 string
			Expected2 string
		}{
			{
				TestCase:  "case 1",
				Input:     "test[0...10]",
				Expected1: "test",
				Expected2: "[0...10]",
			},
			{
				TestCase:  "case 2",
				Input:     "test",
				Expected1: "test",
				Expected2: "",
			},
		}

		for i := range testCases {
			param := testCases[i]
			actual1, actual2 := ReplaceString(param.Input)
			t.Logf("TestCase:%v\n", param.TestCase)
			if actual1 != param.Expected1 {
				t.Errorf("  Failed: actual1 -> %v(%T), expected1 -> %v(%T)\n", actual1, actual1, param.Expected1, param.Expected1)
			}
			if actual2 != param.Expected2 {
				t.Errorf("  Failed: actual2 -> %v(%T), expected2 -> %v(%T)\n", actual2, actual2, param.Expected2, param.Expected2)
			}
		}
	})
}
