# Limiting goroutines with anikhasibul/queue

Hello there, today I'm gonna write another post on golang! If you are not a golang develeper, probably you won't understand me. But I'll try my best to keep it simple.

## The problem:

I think most of the golang developers faced this problem. 
If you try to run 1000 concurrent request to a web server, you will see your PC is nearly about a crash, because, go doesn't wait for a goroutine to be finished and it also doesn't have any limitations about how many goroutines should run at once!

In this snippet:

```go
package main

import "time"

func main(){
	for i := 0; i<100; i++{
		go out(i)
	}
	time.Sleep(2*time.Second)
}

func out(n int){
	time.Sleep(time.Second)
	print(n)
}
```

All of the goroutines will run at once, you can't limit it. And you also can't say thia to your computer, *"Bro, wait for the finishing of all goroutines, then exit!"*  And that's why we've used the `time.Sleep(2*time.Second)` at the end, so that we can wait for the finishing.


## The solution:

If you are a golang programmer, probably you heard of the name of buffered channel. Yes!

Buffered channel receives a value until it reaches to the limit of the buffer, and if a value gets out of the channel, it starts receiving the new waiting values.

We will take the full advantage of it!

## The implementation:

So, I've created this library [github.com/anikhasibul/queue](https://github.com/anikhasibul/queue) using the same way!

Let's see how it works!

```go
package main

import (
	"time"
	"github.com/anikhasibul/queue"
)

// we've limited the goroutine to 1
// it creates a new channel with the buffer of 1 value
var q = queue.New(1)

func main(){
	for i := 0; i<100; i++{
	// add a new value to the channel
	    q.Add()
		go out(i)
	}
	// wait for the finishing
	q.Wait()
}

func out(n int){
	// receive a value from the channel
	defer q.Done()
	time.Sleep(time.Second)
	print(n)
}
```

Here we've linited the goroutines by 1. So, here only one goroutine can run once, you can increase the number. 

### Mechanisms:

* `queue.New(100)` - It creates a buffered channel with the capacity of 100 values. So, it can't store more than 100 values and this is how we limited the goroutines!
* `q.Add()` - Adds a value to the channel. If the amount of value added reaches the maximum buffer, it wait for a space.
* `q.Done()` - Receives a value from the channel. It removes one value and frees space for another value. So the `q.Add()` starts storing again.
* `q.Wait()` - It waits until the channel length becomes 0! So, if all the job has done, the wait time has been finished!


## Conclusion:

I hope you found this post helpful. And as always, thanks for reading.


> Drop a comment if this post has any typo or other mistakes, let me fix it ASAP.
