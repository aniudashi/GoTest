package goTest

import (
	"reflect"
	"testing"
)

func TestDoMergeIntervals1(t *testing.T) {
	srcValues := pairSlice{{
		left:  25,
		right: 30,
	}, {
		left:  20,
		right: 40,
	}, {
		left:  60,
		right: 70,
	}, {
		left:  65,
		right: 90,
	}}

	expectedValues := pairSlice{{
		left:  20,
		right: 40,
	}, {
		left:  60,
		right: 90,
	}}
	dstValues := DoMergeIntervals(srcValues)

	if reflect.DeepEqual(expectedValues, dstValues) {
		t.Logf("DoMergeIntervals 结果[%v]符合预期", dstValues)
	} else {
		t.Fatalf("DoMergeIntervals 结果[%v]不符合预期", dstValues)
	}

}
