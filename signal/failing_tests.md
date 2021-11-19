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

./signal1_test.go