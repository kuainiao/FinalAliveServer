package ffExcel

import "strings"

// 配置形式为server或client或server|client或直接为空
type valueRequired struct {
	required bool
	origin   string
}

func (vr valueRequired) String() string {
	return vr.origin
}

func newValueRequired(v string) *valueRequired {
	v = strings.ToLower(v)
	required := "required" == v
	optional := "optional" == v
	if optional && len(v) > 0 {
		return nil
	}
	return &valueRequired{
		origin:   v,
		required: required,
	}
}
