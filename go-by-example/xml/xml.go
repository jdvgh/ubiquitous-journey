package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", p.Id, p.Name, p.Origin)
}

func main() {

	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethopia", "Brazil"}

	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	fmt.Println(xml.Header + string(out))

	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", " ")
	fmt.Println(string(out))
	/*
	    <?xml version="1.0" encoding="UTF-8"?>
	     <plant id="42">
	   	<name>Coffee</name>
	   	<origin>Ethopia</origin>
	   	<origin>Brazil</origin>
	     </plant>
	*/
	enc := xml.NewDecoder(os.Stdin)

	plant2 := Plant{}

	enc.Decode(&plant2)

	fmt.Println(plant2)

}
