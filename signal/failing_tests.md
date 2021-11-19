The `testing.T` type has a `Log` and `Logf` method.
These work similar to `Print` and `Printf` in the `fmt` package.
They print to log, but we see it only when test fails 
(we can also use **`go test -v`** to print Log without failing the test).

Two basic ways to signal that a test has failed:
- Fail = fail, but keep running
- FailNow = fail now and stop test

You will rarely see these called though, because most of the time people will
call the following methods which combine failures and logging:
- Error     = Log + Fail
- Errorf    = Logf + Fail
- Fatal     = Log + FailNow
- Fatalf    = Logf + FailNow

Which to use?
- If you can let a test keep running, use Error/Errorf
- If a testis completely over and runnning further won't hel at all,
    use Fatal/Fatalf

If not using subsets, Fatal will prevent oher tests in that function from running.
When using subsets (not recommended), Fatal becomes much easier to use.

## Examples of Fatal vs Error

```go
func TestLog(t *testing.T) {
	t.Log("plain log message with t.Log()")              // works like fmt.Print but prints only when test fails
	t.Logf("message with number: %d with t.Logf()", 123) // works like fmt.Printf but prints only when test fails

}

func TestError(t *testing.T) {
	t.Log("plain log message with t.Log()")              // works like fmt.Print
	t.Logf("message with number: %d with t.Logf()", 123) // works like fmt.Printf
	t.Error("failing test with t.Error()")               // fails test and shows previous log msgs
}

func TestFail(t *testing.T) {
	t.Log("pfirst message before t.Fail()") // works like fmt.Print
	t.Fail()                                // will print next message
	t.Log("second message after t.Fail()")  // works like fmt.Printf
}

func TestFailNow(t *testing.T) {
	t.Log("first message before t.FailNow()") // printed
	t.FailNow()                               // will NOT print next message
	t.Log("second message after t.FailNow()") // not printed
}

func TestError2(t *testing.T) {
	t.Log("first message")                 // printed
	t.Error("failing test with t.Error()") // fails test and shows previous AND next log msgs
	t.Log("second message")                // printed

}

func TestFatal(t *testing.T) {
	t.Log("first message before t.Fatal()") // printed
	t.Fatal("failing test with t.Fatal()")  // fails test and shows previous but NOT next log msgs
	t.Log("second message after t.Fatal()") // NOT printed

}
```