// Code generated by "stringer -type Action lifecycle.go"; DO NOT EDIT.

package lifecycle

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NoneAction-0]
	_ = x[DeleteAction-1]
}

const _Action_name = "NoneActionDeleteAction"

var _Action_index = [...]uint8{0, 10, 22}

func (i Action) String() string {
	if i < 0 || i >= Action(len(_Action_index)-1) {
		return "Action(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Action_name[_Action_index[i]:_Action_index[i+1]]
}
