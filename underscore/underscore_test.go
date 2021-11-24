package underscore

import (
	"testing"
)

func TestCamel_first(t *testing.T) {
	arg := "thisIsCamelCaseString"
	want := "this_is_camel_case_string"
	got := Camel(arg)
	if got != want {
		t.Errorf("Camel(%q): wanted: %q, but got: %q", arg, want, got)
	}
}

func TestCamel_simple(t *testing.T) {
	type TestCase struct {
		arg  string
		want string
	}

	testCases := &[]TestCase{
		{"thisIsCamelCaseString", "this_is_camel_case_string"},
		{"with space", "with space"},
		{"endsWithR", "ends_with_r"},
	}

	for _, tc := range *testCases {

		t.Logf("Testing: %q", tc.arg)
		got := Camel(tc.arg)
		if got != tc.want {
			t.Errorf("Camel(%q): wanted: %q, but got: %q", tc.arg, tc.want, got)
		}
	}

}

func TestCamel(t *testing.T) {
	type args struct {
		inp string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Camel(tt.args.inp); got != tt.want {
				t.Errorf("Camel() = %v, want %v", got, tt.want)
			}
		})
	}
}
