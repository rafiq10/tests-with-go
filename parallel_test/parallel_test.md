Example use cases for parallel tests:
1. Simulating a real-world scenario
2. Verify that a type is truly thread-safe

(1) A web app with many users
(2) Verify that your in-memory cache can handle multiple concurrent web requests using it

Parallelism could also mean more work:
- Tests can't use as many hard-coded values; eg.: unique email constraints
- Tests might try to use shared resources incorrectly eg.:
  image manipulation on the same image or sharing a DB that doesn't suppor multiple
  concurrent connections
  
### Setup and Teardown in Parallel SubTests
The next examples will sow how to use Setup and Teardown for parallel subtests

  ```go
  func TestParallelSubtestTeardown(t *testing.T) {
	// This shows how it will run 
    // teardown and deferred teardown after sub1 BUT before sub2!

	fmt.Println("setup")
	defer fmt.Println("deferred teardown")
	t.Run("sub1", func(t *testing.T) {
		t.Parallel()
		time.Sleep(time.Second)
		fmt.Println("su1 done")
	})
	t.Run("sub2", func(t *testing.T) {
		t.Parallel()
		time.Sleep(time.Second)
		fmt.Println("sub2 done")
	})
	fmt.Println("teardown")
}

```

To correctly contruct the teardown you need to put the parallel subtests
in the separate NON-PARALLEL subtest like below

```go
func TestParallelSubtestCorrectSetup(t *testing.T) {
	// this will run teardown and deferred teardown after BOTH sub1 AND sub2

	fmt.Println("setup")
	defer fmt.Println("deferred teardown")
	t.Run("group", func(t *testing.T) {

		t.Run("sub1", func(t *testing.T) {
			t.Parallel()
			time.Sleep(time.Second)
			fmt.Println("su1 done")
		})
		t.Run("sub2", func(t *testing.T) {
			t.Parallel()
			time.Sleep(time.Second)
			fmt.Println("sub2 done")
		})

	})

	fmt.Println("teardown")
}
```
### Closure issues with subtests

A common error is to pass the parameters from outside loop to the subtest as a copy/value:
```go
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
```

One way to cope with that issue is to 
copy the item into local loop variable before each subtest Run:
```go
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
		//copy the testCase into the loop variable
		localTc := tc
		t.Run(fmt.Sprintf("arg=%d", localTc.arg), func(t *testing.T) {
			t.Parallel()
			t.Logf("Testing with: arg=%d, want=%d", localTc.arg, localTc.want)
			if localTc.arg*localTc.arg != localTc.want {
				t.Errorf("%d^2 != %d", localTc.arg, localTc.want)
			}
		})
	}
}
```

Another way to deal with the issue is to return a testing function as a result of subtest;
with a 
```go   
var tc struc{}t
```
as an argument and 
```go   
func(t *testing.T) {return func(t *testing.T) {}} (tc)
```
as a result, where '(tc)' part is calling the function func(tc testCase) ...
and passing te struct BEFORE t.Parallel inside the t.Run is called

```go
func TestGotchaCorrectAnotherWay(t *testing.T) {
    //for _, tc := range testCases {
    //  t.Run("bla bla", func (tc struct {}) func(t *testing.T) {
    //    return func(t *testing.T) {}} (tc))} 

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
```