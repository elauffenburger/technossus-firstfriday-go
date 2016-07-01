package test

import (
	"testing"
	"time"
	"github.com/elauffenburger/technossus-firstfriday-go/demo/conc"
)

func TestChannel(t *testing.T) {
	channel1 := make(chan string)
	done := make(chan interface{})

	// 500ms
	blockMs := time.Duration(500) * time.Millisecond

	go func() {
		msg, ok := <-channel1
		if !ok {
			return
		}

		// signal that you've received the message
		t.Logf("received: '%v' on channel 1", msg)

		if !testing.Short() {
			// wait for n ms
			// notice that we could have used time.Sleep, but it's more fun this way :D
			<-time.After(blockMs)

			// signal we're done
			done <- true
		}
	}()

	channel1 <- "hi channel 1!"

	// if we're running the whole thing...
	if !testing.Short() {
		// wait until we get a signal from done
		t.Logf("About to be blocked for %s ms", blockMs)
		<-done
	}

	t.Log("Done!\n")
}

func TestSemaphore(t *testing.T) {
	numChannels := 3

	sem := make(chan interface{}, 1)
	done := make(chan interface{})

	action := func(i int) {
		t.Logf("hi from goroutine %d", i)
	}

	for i := 0; i < numChannels; i++ {
		sem <- true

		go func(chanNum int) {
			action(chanNum)
			<-sem

			if chanNum == numChannels-1 {
				done <- true
			}
		}(i)
	}

	<-done

	t.Log("Done!")
}

func TestManualSlidingWindow(t *testing.T) {
	windowsize := 3

	window := make(chan interface{}, windowsize)
	done := make(chan interface{})

	numjobs := 20
    completedjobs := 0
	for i := 0; i < numjobs; i++ {
		go func(num int) {
            // insert a token.  if the window is full, this will block
            // until there is a spot available
			window <- true

            // log that we're doing some work
			t.Logf("[%d] Working...", num)
            // do some work
			time.Sleep(time.Duration(20) * time.Millisecond)
            // log that we're done with work
            t.Logf("  [%d] Done!", num)

            // mark job as complete
            completedjobs++

            // release a token.  this will let one of the blocked threads
            // get a token and do its work
			<-window

            // if we're all done with the jobs, let the calling thread know
			if completedjobs == numjobs {
				done <- true
			}
		}(i)
	}

    // blocks until the jobs are done
	<-done
}

func TestSlidingWindowType(t *testing.T) {
	windowsize := 3
	numjobs := 20

    window := conc.SlidingWindowFactory(windowsize, numjobs)

	for i := 0; i < numjobs; i++ {
		go func(num int) {
            // Get a token.  if the window is full, this will block
            // until there is a spot available
			window.GetToken()

            // log that we're doing some work
			t.Logf("[%d] Working...", num)
            // do some work
			time.Sleep(time.Duration(20) * time.Millisecond)
            // log that we're done with work
            t.Logf("  [%d] Done!", num)

            // mark job as complete. this will let one of the blocked threads
            // get a token and do its work
			window.CompleteJob()
		}(i)
	}

    // blocks until the jobs are done
	<-window.Done
}
