package parallel

import (
	"fmt"
	"testing"
)

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

func Test_subtestsParallel(t *testing.T) {
	// two subtests run in parallel
	// it will give he result:

	// === RUN   TestB
	// --- PASS: TestB (0.00s)
	// === RUN   Test_subtestsParallel
	// === RUN   Test_subtestsParallel/Sub1
	// === PAUSE Test_subtestsParallel/Sub1
	// === RUN   Test_subtestsParallel/Sub2
	// === PAUSE Test_subtestsParallel/Sub2
	// === CONT  Test_subtestsParallel/Sub1
	// === CONT  Test_subtestsParallel/Sub2
	// --- PASS: Test_subtestsParallel (0.00s)
	// 	--- PASS: Test_subtestsParallel/Sub1 (0.00s)
	// 	--- PASS: Test_subtestsParallel/Sub2 (0.00s)
	// === CONT  TestSomething
	// --- PASS: TestSomething (0.00s)
	// === CONT  TestA
	// --- PASS: TestA (0.00s)
	// PASS

	t.Run("Sub1", func(t *testing.T) {
		t.Parallel()
	})
	t.Run("Sub2", func(t *testing.T) {
		t.Parallel()
	})
}

func TestGotcha(t *testing.T) {
	// this will print '"Testing with arg=4, want=16' for each iteration
	// because we are coping the value of the struct into the closeure each time
	testCases := []struct {
		arg  int
		want int
	}{
		{2, 5},
		{2, 9}, {4, 16},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("arg=%d", tc.arg), func(t *testing.T) {
			t.Parallel()
			t.Logf("Testing with: arg=%d, want=%d", tc.arg, tc.want)
			if tc.arg*tc.arg != tc.want {
				t.Errorf("%d^2 != %d", tc.arg, tc.want)
			}
		})
	}
}

func TestGotchaCorrect(t *testing.T) {
	// copy the item into local loop variable before each subtest Run

	testCases := []struct {
		arg  int
		want int
	}{
		{2, 4},
		{3, 9},
		{4, 16},
	}
	for _, tc := range testCases {
		localTc := tc //copy the testCase into the loop (local) variable - DO NOT DELETE!
		t.Run(fmt.Sprintf("arg=%d", localTc.arg), func(t *testing.T) {
			t.Parallel()
			t.Logf("Testing with: arg=%d, want=%d", localTc.arg, localTc.want)
			if localTc.arg*localTc.arg != localTc.want {
				t.Errorf("%d^2 != %d", localTc.arg, localTc.want)
			}
		})
	}
}

func TestGotchaCorrectAnotherWay(t *testing.T) {
	// copy the item into local loop variable before each subtest Run

	type testCase struct {
		arg  int
		want int
	}
	testCases := []testCase{
		{2, 4},
		{3, 9},
		{4, 16},
	}
	for _, tc := range testCases {

		t.Run(fmt.Sprintf("arg=%d", tc.arg),
			func(tc testCase) func(t *testing.T) {
				return func(t *testing.T) {
					t.Parallel()
					t.Logf("Testing with: arg=%d, want=%d", tc.arg, tc.want)
					if tc.arg*tc.arg != tc.want {
						t.Errorf("%d^2 != %d", tc.arg, tc.want)
					}
				}
			}(tc))
	}
}
