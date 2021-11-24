package parallel

import (
	"fmt"
	"testing"
	"time"
)

func TestParallel(t *testing.T) {
	// this will run teardown and deferred teardown after sub1 BUT before sub2!

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
