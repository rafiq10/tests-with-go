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
### Sample of when to use t.Fatal() and when to use t.Error():
Notice going from more general and important errors to more specific.

```go

type Person struct {
	Age        int
	Name       string
	Occupation string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	p := Person{
		Age:        30,
		Name:       "Bob Jones",
		Occupation: "Nurse",
	}

	data, err := json.Marshal(p)
	if err != nil {
		http.Error(w, "Internal servererror", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}


func TestHandler1(t *testing.T) {
	w := httptest.NewRecorder() //fake http response writer
	r, err := http.NewRequest(http.MethodGet, "", nil)

	if err != nil {
		t.Fatalf("http.NewRequest() err = %s", err)
	}
	Handler(w, r)

	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Fatalf("Handler() status = %d; wanted: %d", resp.StatusCode, 200) //do NOT continue the test
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Handler() Content-Type = %q; wanted: %q", contentType,
			"application/json") //DO CONTINUE with test, it is not so relevant to have json t this moment
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// fail test if not able to read body
		// will NOT put "wanted" as it is pretty obious we expect nil error
		t.Fatalf("iotil.ReadAll(resp.Bdy) err = %s", err)
	}

	var p Person
	err = json.Unmarshal(data, &p) // unmarshall into a new Person struct
	if err != nil {
		// doesn't make sense to continue with property's tests if unable to unmarshall
		t.Fatalf("json.Unmarshall(resp.Body) err = %s", err)
	}

	if p.Age != 21 {
		// it is not so important to stop the rest of the tests
		t.Errorf("person.Age = %d; wanted: %d", p.Age, 21)
	}

	if p.Name != "Rafa" {
		t.Errorf("person.Name = %s; wanted: %s", p.Name, "Rafa")
	}
}
```

### Useful error messages:

The goal of useful error mesage is to tell what went wrong in the test.
It should be meaning for the debugging.

```go
    t.Errorf("SomeFunction(%v) err = %v; wanted: %v)

    t.Fatalf(http.NewRequest(%s,%s,nil) err = %v; wanted: %v",http.MethodGet,"\\",err.Error(),"somethingDesired")

```

Kind of standard way to deal with testing:
```go
    if got != wanted{
        // do stuff
    } 
```

If the object is too big, it might have sense to print only important properties:
```go
err  SomeFunc(p)
if err != nil{
    t.Fatalf("SmeFunc(name=%s, age=%d)", p.Name, p.Age)
} 
```

```go
```