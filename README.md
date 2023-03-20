# Description

This is my batch handler written by golang
It will wait until timeout or specific number I want, the handler will be execute

## Installation

```sh
$ go get -u github.com/QuanLeo/batch-handler
```

## Usage

```golang
package main

import (
	"fmt"
	"time"

	batch "github.com/QuanLeo/batch-handler"
)

func handler(data []any) error {
	for _, v := range data {
		data := v.(Demo)
		fmt.Printf(" | %v : %v", data.Id, data.Value)
	}

	fmt.Println()
	return nil
}

type Demo struct {
	Id    int
	Value int
}

func main() {
	//when list item = 30 -> trigger handler function
	// t := thrott.New(0*time.Second, 30, handler)

	//Every 5seconds -> trigger handler function
	// t := thrott.New(5*time.Second, 0, handler)

	//Every 1seconds or 10 items -> trigger handler function
	t := batch.New(1*time.Second, 10, handler)
	go func() {
		i := 0
		for {
			t.Push(Demo{Id: 1, Value: i})

			//Test for terminate batch handle. If i == 100 -> terminate -> increase go routine
			if i == 100 {
				t.Terminate()
			}

			time.Sleep(100 * time.Millisecond)
			i++
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(31 * time.Second)
}
```

## License

This code is licensed under the MIT license.
