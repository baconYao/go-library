package main

import (
	"fmt"
	"io/ioutil"
	"log"

    "github.com/baconYao/go-library/protobuf/src/simple"
	"github.com/golang/protobuf/proto"
)

func main() {
    sm := doSimple()
    readAndWriteDemo(sm)
}

func readAndWriteDemo(sm proto.Message) {
    wtireToFile("simple.bin", sm)
    sm2 := &simplepb.SimpleMessage{}
    readFromFile("simple.bin", sm2)
    fmt.Println("Read the content: ", sm2)
}

func wtireToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialise to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}

	fmt.Println("Data has been written!")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
    in, err := ioutil.ReadFile(fname)
    if err != nil {
        log.Fatalln("Something went wrong when reading the file", err)
        return err
    }

    err2 := proto.Unmarshal(in, pb)
    if err2 != nil {
        log.Fatalln("Couldn't put the bytes into the protocol buffers struct", err2)
        return err2
    }

	fmt.Println("Data has been read!")
	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "Bacon yao is my name",
		SampleList: []int32{1, 3, 5, 7},
	}
	fmt.Println(sm)

	sm.Name = "I renamed you"
	fmt.Println(sm)
	fmt.Println("The ID is: ", sm.GetId())

	return &sm
}
