package future_test

import (
	"fmt"
	"log"
	"time"

	"github.com/Laugusti/go2draft/generics/future"
)

func ExamplePool() {
	p, err := future.NewPool(1)
	if err != nil {
		log.Fatal(err)
	}
	f, err := p.NewFuture(func() (int, error) {
		time.Sleep(5 * time.Second)
		return 42, nil
	})
	if err != nil {
		log.Fatal(err)
	}
	v := f.MustGet()
	fmt.Println(v)
	// Output:
	// 42
}

func ExamplePool_closed() {
	p, err := future.NewPool(1)
	if err != nil {
		log.Fatal(err)
	}
	f, err := p.NewFuture(func() (int, error) {
		time.Sleep(5 * time.Second)
		return 42, nil
	})
	if err != nil {
		log.Fatal(err)
	}
	p.Stop()
	_, err = f.Get()
	fmt.Println(err)
	_, err = p.NewFuture(func() (int, error) {
		return 42, nil
	})
	fmt.Println(err)
	// Output:
	// Future pool was stopped
	// Future pool was stopped
}
