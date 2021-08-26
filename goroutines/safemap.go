package goroutines

type SafeMap struct {
	M map[string]int
}

type Value struct {
	K string
	V int
}

func (t *SafeMap) Iterator() (it <-chan *Value) {
	channel := make(chan *Value, 32)
	go func() {
		defer close(channel)
		for k, v := range t.M {
			channel <- &Value{k, v}
		}
	}()

	return channel
}
