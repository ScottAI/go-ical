package go_ical

import (
	"bytes"
	"io"
	"log"
	"time"
)

func ExampleDecoder() {
	var r io.Reader

	dec := NewDecoder(r)
	for {
		c, err := dec.Decode()
		cal := NewCalendar()
		cal.SubComponentsObj = c.SubComponents()
		cal.NameObj = c.Name()
		cal.PropertiesObj = c.Properties()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		for _, event := range cal.GetEvents() {
			props := event.Properties()
			var summary string
			for _,p := range props{
				if p.Name == PropSummary{
					s,err := p.GetToText()
					if err != nil{
						panic(err)
					}
					summary = s
				}
			}

			log.Printf("Found event: %v", summary)
		}
	}
}

func ExampleEncoder() {
	event := NewEvent()
	p := NewProperty(PropUID)
	p.Value = ToText("uid@example.org")
	event.PropertiesObj = append(event.PropertiesObj,*p)

	p = NewProperty(PropDatetimeStamp)
	p.SetFromDatetime(time.Now())
	event.PropertiesObj = append(event.PropertiesObj,*p)
	p = NewProperty(PropSummary)
	p.SetFromText("Event Test")
	event.PropertiesObj = append(event.PropertiesObj,*p)

	p = NewProperty(PropDatetimeStart)
	p.SetFromDatetime(time.Now().Add(24*time.Hour))

	cal := NewCalendar()
	cal.SubComponentsObj = append(cal.SubComponentsObj, event.SubComponents() ...)

	var buf bytes.Buffer
	if err := NewEncoder(&buf).Encode(cal); err != nil {
		log.Fatal(err)
	}

	log.Print(buf.String())
}
