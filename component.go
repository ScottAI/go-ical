package go_ical

import (
	"fmt"
)

/*
The value for the "component" parameter is defined as follows:

       component = "VEVENT"
                 / "VTODO"
                 / "VJOURNAL"
                 / "VFREEBUSY"
                 / "VTIMEZONE"
                 / iana-token
                 / x-name
 */
type Component interface {
	Name() string
	Properties() []Property //TODO:prepare for IANA properties   RFC 5545 8
	SubComponents() []Component
	IsAvailable() error
	encode(enc *Encoder) error
}

type ComponentObj struct {
	NameObj string
	PropertiesObj []Property
	SubComponentsObj []Component
}

func (com *ComponentObj) Name() string {
	return com.NameObj
}

func (com *ComponentObj) Properties() []Property {
	return com.PropertiesObj
}

func (com *ComponentObj) SubComponents() []Component {
	return com.SubComponentsObj
}

//changed from Encoder's encodeComponent
func (com *ComponentObj) encode(enc *Encoder) error {
	if err := com.IsAvailable();err != nil{
		return err
	}
	fmt.Fprint(enc.w,"BEGIN:"+com.Name(),"\r\n")
	for _,p := range com.Properties(){
		enc.encodeProperty(&p)
	}

	for _,c := range com.SubComponents(){
		c.encode(enc)
	}
	fmt.Fprint(enc.w,"END:"+com.Name()+"\r\n")

	return nil
}

//to check componet
func (com *ComponentObj) IsAvailable() error {
	switch com.Name() {
	case CompCalendar:
		if len(com.SubComponents()) == 0{
			return fmt.Errorf("ical: VCALENDAR is empty!,can not encode")
		}
		//if VEVENT has no Prop METHOD,VCALENDAR Must have DTSTART
		for _,sub := range com.SubComponents(){
			if sub.Name() == CompEvent{
				isMethod := false
				for _,p := range sub.Properties(){
					if p.Name == PropMethod{
						isMethod = true
					}
				}
				if isMethod{
					continue
				}
				isDTSTART := false
				for _,p := range com.Properties(){
					if p.Name == PropDatetimeStart{
						isDTSTART = true
					}
				}
				if !isDTSTART{
					return fmt.Errorf("ical: DTSTART is required,when VEVENT has no Prop METHOD")
				}
			}
		}
	case CompEvent:
		for _,sub := range com.SubComponents(){
			if sub.Name() != CompAlarm{
				return fmt.Errorf("ical:In EVENT only ALARM is allowed,but got %q",sub.Name())
			}
		}
		isEnd := false
		isDuration := false
		for _,p := range com.Properties(){
			if p.Name == PropDatetimeEnd{
				isEnd = true
			} else if p.Name == PropDuration{
				isDuration = true
			}
		}
		if isEnd && isDuration{
			return fmt.Errorf("ical:in EVENT can not use DTEND and DURATION at same time")
		}
	case CompTodo:
		for _,sub := range com.SubComponents(){
			if sub.Name() != CompAlarm{
				return fmt.Errorf("ical:In TODO only ALARM is allowed,but got %q",sub.Name())
			}
		}
		isDUE := false
		isDURATION := false
		isDTSTART := false
		for _,p := range com.Properties(){
			if p.Name == PropDatetimeDue{
				isDUE = true
			} else if p.Name == PropDuration{
				isDURATION = true
			} else if p.Name == PropDatetimeStart{
				isDTSTART = true
			}
		}
		if isDUE && isDURATION{
			return fmt.Errorf("ical:In TODO can not use DUE and DURATION at the same time")
		}
		if isDURATION && (!isDTSTART){
			return fmt.Errorf("ical:In TODO DTSTART is required when DURATION existed")
		}
	case CompJournal,CompFreebusy:
		if len(com.SubComponents()) > 0{
			return fmt.Errorf("ical:can not have subcomponents in %q",com.Name())
		}
	case CompTimezone:
		if len(com.SubComponents()) == 0{
			return fmt.Errorf("ical:TIMEZONE is empty,expect STANDARD or DAYLIGHT component")
		}
		for _,sub := range com.SubComponents(){
			if sub.Name() != CompTimezoneStandard && sub.Name() != CompTimezoneDaylight{
				return fmt.Errorf("ical:STANDARD and DAYLIGHT are allowed in TIMEZONE,but got %q",sub.Name())
			}
		}
	//case CompAlarm:
	}
	for _,pn := range OnlyOnePropMap[com.Name()]{
		n := 0
		for _,p := range com.Properties(){
			if pn == p.Name{
				n += 1
			}
		}
		if n != 1{
			return fmt.Errorf("ical:%q MUST have only one prop %q,but got %q",com.Name(),pn,n)
		}
	}
	for _,pn := range OneOrZeroPropMap[com.Name()]{
		n := 0
		for _,p := range com.Properties(){
			if pn == p.Name{
				n += 1
			}
		}
		if n > 1{
			return fmt.Errorf("ical:%q SHOULD have one or zero prop %q,but got %q",com.Name(),pn,n)
		}
	}
	return nil
}



type VEvent struct {
	ComponentObj
}


type Attendee struct {
	ComponentObj
}


type VTodo struct {
	ComponentObj
}

type VJournal struct {
	ComponentObj
}


type VFreeBusy struct {
	ComponentObj
}


type VTimezone struct {
	ComponentObj
}

type VAlarm struct {
	ComponentObj
}

type VTZSTANDARD struct {
	ComponentObj
}

type VTZDAYLIGHT struct {
	ComponentObj
}



