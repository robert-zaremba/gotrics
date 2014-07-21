package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

const jsonStream = `{
"Name": "Ed",
"Type": "A",
"Data": {
  "F1": 1
}}
{
"Name": "Ed",
"Type": "B",
"Data": {
  "F2": "krowa"
}}
`

type MessageT struct {
	Name, Type string
	Data       interface{}
}

type messageT struct {
	Name, Type string
	Data       json.RawMessage
}
type Data1 struct {
	F1 int
}
type Data2 struct {
	F2 string
}

func (m *MessageT) UnmarshalJSON(b []byte) error {
	var mp messageT
	err := json.Unmarshal(b, &mp)
	if err != nil {
		log.Fatal(err)
		return err
	}
	// decode Data into a proper type
	switch mp.Type {
	case "A":
		m.Data = new(Data1)
	case "B":
		m.Data = new(Data2)
	default:
		log.Fatal("Wrong message type")
	}
	err = json.Unmarshal(mp.Data, m.Data)
	if err != nil {
		log.Fatal(err)
		return err
	}
	// copy  fields
	m.Name = mp.Name
	m.Type = mp.Type
	return nil
}

func customizeUnmarshall() {
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		fmt.Println("New object *******************")
		var m MessageT
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Decoded data: %T\n", m.Data)
	}
}

func setMarshallIndentation() {
	m := make(map[string]interface{})
	m["1"] = 1
	m["2"] = "dwa"
	m["mapa"] = map[string]int{"1": 10, "2": 20}
	b, _ := json.MarshalIndent(m, "$$", "__")
	println(string(b))
}

func main() {
	fmt.Println("### Example1: marshalling indentaion")
	setMarshallIndentation()

	fmt.Println("\n\n### Example2: customize unmarshaller to correct decode an interface{} into correct data type")
	customizeUnmarshall()
}
