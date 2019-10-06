package main

import "fmt"
import "encoding/json"
import "time"
type Message struct{
	Name string
	Body string
	Time int64
}
func main(){
	fmt.Println("hello world")
	fmt.Println(time.Now())
	for i:= 0; i < 100000 * 30; i++{
		m := Message{"Alice", "Hello", time.Now().UnixNano()}
		b, err := json.Marshal(m)
		if err != nil{
			fmt.Println("encode error")
			return
		}

		var dm Message

		err = json.Unmarshal(b, &dm)
		if err != nil{
			fmt.Println("decode error")
		}
	}
	fmt.Println(time.Now())
}
