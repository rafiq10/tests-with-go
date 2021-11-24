package parallel

import "testing"

func TestSomething(t *testing.T) {
	//the tests with t.Parallel() pause and group before running non-parallel tests
	// so the TestSomething and TestA will pause and resume after TestB is completed
	t.Parallel()
}

func TestA(t *testing.T) {
	t.Parallel()
}

func TestB(t *testing.T) {

}
