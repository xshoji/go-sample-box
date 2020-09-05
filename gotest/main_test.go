package main

import "testing"

func TestReplaceStringSuccess(t *testing.T) {
	t.Run("DataDriven", func(t *testing.T) {

		testCases := []struct {
			TestCase  string
			Input     string
			BaseName  string
			ArrayPart string
		}{
			{
				TestCase:  "Success",
				Input:     "test[0...10]",
				BaseName:  "test",
				ArrayPart: "[0...10]",
			},
			{
				TestCase:  "Unexpected format",
				Input:     "test",
				BaseName:  "test",
				ArrayPart: "",
			},
		}

		for i := range testCases {
			param := testCases[i]
			actualBaseName, actualArrayPart := ReplaceString(param.Input)
			t.Logf("TestCase:%v\n", param.TestCase)

			if actualBaseName != param.BaseName {
				t.Errorf("  Failed: actualBaseName -> %v(%T), expected -> %v(%T)\n", actualBaseName, actualBaseName, param.BaseName, param.BaseName)
			}
			if actualArrayPart != param.ArrayPart {
				t.Errorf("  Failed: actualArrayPart -> %v(%T), expected -> %v(%T)\n", actualArrayPart, actualArrayPart, param.ArrayPart, param.ArrayPart)
			}
		}
	})
}
