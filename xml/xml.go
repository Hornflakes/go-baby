package main

import (
	"encoding/xml"
	"fmt"
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
	blueHydrangea := &Plant{
		Id: 23, Name: "BlueHydrangea"}
	blueHydrangea.Origin = []string{"China", "Korea", "Japan"}

	out, _ := xml.MarshalIndent(blueHydrangea, " ", " ")
	fmt.Println(string(out))

	fmt.Println(xml.Header + string(out))

	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	basil := &Plant{Id: 17, Name: "Basil"}
	basil.Origin = []string{"India", "California"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{blueHydrangea, basil}

	out, _ = xml.MarshalIndent(nesting, " ", " ")
	fmt.Println(string(out))
}
