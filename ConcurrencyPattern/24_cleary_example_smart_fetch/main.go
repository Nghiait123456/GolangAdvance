// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type (
	OneFetch struct {
		url string
	}
	OneFetchInterface interface {
		Fetch() DataOneFetch
	}

	DataOneFetch struct {
		Data string
		UUID uint64
	}
)

func (o *OneFetch) Fetch() DataOneFetch {
	id := rand.Uint64()
	data := o.url + string(id)
	return DataOneFetch{
		Data: data,
		UUID: id,
	}
}

func NewOneFetch(url string) OneFetchInterface {
	return &OneFetch{
		url: url,
	}
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type (
	ControlOneStreamFetch struct {
		Fetch OneFetchInterface
		Data  chan DataOneFetch
		Quit  chan bool
	}

	ControlOneStreamFetchInterface interface {
		loop(AllData chan DataOneFetch)
		SendQuit(data bool)
	}
)

func (c *ControlOneStreamFetch) SendQuit(data bool) {
	c.Quit <- data
}

func (c *ControlOneStreamFetch) loop(AllData chan DataOneFetch) {
	startOneFetch := time.After(1 * time.Second)
	oneFetchDone := make(chan bool, 1)
	timeOut := time.After(100 * time.Second)
	go func() {
		for {
			select {
			case <-startOneFetch:
				fmt.Println("start one stream fetch")
				go func() {
					fmt.Println("start one fetch")
					data := c.Fetch.Fetch()
					fmt.Println("push to stream data")
					c.Data <- data
					fmt.Println("push to oneFetchDone")
					oneFetchDone <- true
				}()

			case <-oneFetchDone:
				fmt.Println("---------------------------------------------------fetch done, reset timer-------------------------------------------------------")
				// reset timer
				startOneFetch = time.After(1 * time.Second)

			case d := <-c.Data:
				AllData <- d // use design patter FallIn
				fmt.Printf("data of channel = %v \n", d)

			case <-c.Quit:
				fmt.Println("quit channel")
				close(c.Data)
				close(c.Quit)
				return

			case <-timeOut:
				fmt.Println("timeout")
				close(c.Data)
				close(c.Quit)
				return
			}
		}
	}()

}

func NewControlOneStreamFetch(Fetch OneFetchInterface) ControlOneStreamFetchInterface {
	return &ControlOneStreamFetch{
		Fetch: Fetch,
		Data:  make(chan DataOneFetch),
		Quit:  make(chan bool),
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type (
	MergerAllStreamData struct {
		AllStream     []ControlOneStreamFetchInterface
		AllData       chan DataOneFetch
		QuitAllStream chan bool
	}
	MergerAllStreamDataInterface interface {
		Merger()
	}
)

func (m MergerAllStreamData) Merger() {
	// start all stream
	for _, v := range m.AllStream {
		v.loop(m.AllData)
	}
	timeOut := time.After(10 * time.Second)

	//merger data
	go func() {
		for {
			select {
			case d := <-m.AllData:
				fmt.Printf("all Data Stream, %v \n", d)
			case <-m.QuitAllStream:
				fmt.Println("send Quit ALl channel")
				for _, v := range m.AllStream {
					v.SendQuit(true)
				}
				// todo case send quit one channel
			case <-timeOut:
				fmt.Println("send Quit ALl channel")
				for _, v := range m.AllStream {
					v.SendQuit(true)
				}
			}
		}
	}()
}

func NewMergerAllStreamData(listStream ...ControlOneStreamFetchInterface) MergerAllStreamDataInterface {
	var allStream []ControlOneStreamFetchInterface
	for _, one := range listStream {
		allStream = append(allStream, one)
	}

	return &MergerAllStreamData{
		AllStream:     allStream,
		AllData:       make(chan DataOneFetch),
		QuitAllStream: make(chan bool),
	}
}
func main() {
	Streams := NewMergerAllStreamData(
		NewControlOneStreamFetch(NewOneFetch("24_cleary_example_smart_fetch")),
		NewControlOneStreamFetch(NewOneFetch("abc")),
		NewControlOneStreamFetch(NewOneFetch("def")),
	)
	Streams.Merger()

	time.Sleep(10000 * time.Second)
}
