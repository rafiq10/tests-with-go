package signal

import (
	"testing"
)

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
