// Code generated by go2go; DO NOT EDIT.


//line lb_test.go2:1
package chans

import (
	"sync"
	"testing"
	"math/rand"
)

//line lb_test.go2:31
func getInputChan() <-chan int {
	input := make(chan int, 20)
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	go func() {
		for num := range numbers {
			input <- num
		}
		close(input)
	}()
	return input
}

func TestFanin(t *testing.T) {
	chs := make([]<-chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = getInputChan()
	}

	out := instantiate୦୦Fanin୦int(chs...)
	count := 0
	for range out {
		count++
	}
	if count != 100 {
		t.Fatalf("Fanin failed")
	}
}

func TestLB(t *testing.T) {
	ins := make([]<-chan int, 10)
	for i := 0; i < 10; i++ {
		ins[i] = getInputChan()
	}
	outs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		outs[i] = make(chan int, 10)
	}
//line lb_test.go2:67
 instantiate୦୦LB୦int(func(m int) int {
//line lb_test.go2:69
  return rand.Intn(m)
	}, ins, outs)
}
//line lb_test.go2:10
func instantiate୦୦Fanin୦int(chans ...<-chan int,) <-chan int {
	buf := 0
	for _, ch := range chans {
		if len(ch) > buf {
//line lb_test.go2:13
   buf = len(ch)
//line lb_test.go2:13
  }
	}
	out := make(chan int, buf)
	wg := sync.WaitGroup{}
	wg.Add(len(chans))
	for _, ch := range chans {
		go func(ch <-chan int,) {
			for v := range ch {
//line lb_test.go2:20
    out <- v
//line lb_test.go2:20
   }
					wg.Done()
		}(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
//line lb.go2:21
func instantiate୦୦LB୦int(randomizer func(max int) int, ins []<-chan int, outs []chan int,) {
//line lb.go2:21
 instantiate୦୦Fanout୦int(randomizer, instantiate୦୦Fanin୦int(ins...), outs...)
//line lb.go2:23
}
//line lb.go2:12
func instantiate୦୦Fanout୦int(randomizer func(max int) int, In <-chan int, Outs ...chan int,) {
	l := len(Outs)
	for v := range In {
		i := randomizer(l)
		if i < 0 || i > l {
//line lb.go2:16
   i = rand.Intn(l)
//line lb.go2:16
  }
				go func(v int,) { Outs[i] <- v }(v)
	}
}

//line lb.go2:19
type _ sync.Cond
//line lb.go2:19
type _ testing.B
//line lb.go2:19
type _ rand.Rand
