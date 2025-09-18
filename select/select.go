package select_package

import (
	"fmt"
	"net/http"
	"time"
)

// func Racer(a, b string) (winner string) {
// 	aDuration := measureResponseTime(a)
// 	bDuration := measureResponseTime(b)

// 	if aDuration < bDuration {
// 		return a
// 	}

// 	return b
// }

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }

func Racer(a, b string) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

// A channel receive (<-ch) becomes ready in two cases:
// When a value is sent into it.
// When it is closed (even with no value).

// 		ping
// We have defined a function ping which creates a chan struct{} and returns it.
// In our case, we don't care what type is sent to the channel, we just want to signal we are done and closing the channel works perfectly!
// Why struct{} and not another type like a bool? Well, a chan struct{} is the smallest data type available from a memory perspective so we get no allocation versus a bool. Since we are closing and not sending anything on the chan, why allocate anything?
// Inside the same function, we start a goroutine which will send a signal into that channel once we have completed http.Get(url).

//      Always make channels   **********
// Notice how we have to use make when creating a channel; rather than say var ch chan struct{}. When you use var the variable will be initialised with the "zero" value of the type. So for string it is "", int it is 0, etc.
// For channels the zero value is nil and if you try and send to it with <- it will block forever because you cannot send to nil channels

//     select
// You'll recall from the concurrency chapter that you can wait for values to be sent to a channel with myVar := <-ch. This is a blocking call, as you're waiting for a value.
// select allows you to wait on multiple channels. The first one to send a value "wins" and the code underneath the case is executed.
// We use ping in our select to set up two channels, one for each of our URLs. Whichever one writes to its channel first will have its code executed in the select, which results in its URL being returned (and being the winner).
