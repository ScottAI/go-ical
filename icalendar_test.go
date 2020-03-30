package go_ical

import (
	"strings"
	"testing"
	"time"
)

func toCRLF(s string) string {
	return strings.ReplaceAll(s, "\n", "\r\n")
}

var exampleCalendarStr = toCRLF(`BEGIN:VCALENDAR
PRODID:-//xyz Corp//Scott WORK Calendar Version 1.0//CN
VERSION:2.0
BEGIN:VEVENT
CATEGORIES:CONFERENCE
DESCRIPTION;ALTREP="cid:part1.0001@example.org":Test Test Test event
DTEND:19960920T220000Z
DTSTAMP:19960704T120000Z
DTSTART:19960918T143000Z
ORGANIZER:mailto:jsmith@example.com
STATUS:CONFIRMED
SUMMARY;FOO=bar,"b:az":Test summary
UID:uid1@example.com
END:VEVENT
END:VCALENDAR
`)

var exampleCalendar = NewCalendar()

func prepareExampleCalendar() *Calendar {

	event := NewEvent()
	event.NameObj = CompEvent
	event.SetProperty(PropDatetimeStamp,"19960704T120000Z")
	event.SetProperty(PropUID,"uid1@example.com")
	event.SetProperty(PropOrganizer,"mailto:jsmith@example.com")
	event.SetProperty(PropDatetimeStart,"19960918T143000Z")
	event.SetProperty(PropDatetimeEnd,"19960920T220000Z")
	event.SetProperty(PropStatus,"CONFIRMED")
	event.SetProperty(PropCategories,"CONFERENCE")
	event.SetProperty(PropSummary,"Test summary")
	param := NewParamItem(Paramaltrep,[]string{"cid:part1.0001@example.org"})
	event.SetProperty(PropDescription,"Test Test Test event",param)
	exampleCalendar.AddComponent(Component(event))
	return exampleCalendar
}

func TestCalendar(t *testing.T) {
	cal := prepareExampleCalendar()
	events := cal.GetEvents()
	if len(events) != 1 {
		t.Fatalf("len(Calendar.Events()) = %v, want 1", len(events))
	}

	wantSummary := "Test summary"
	var sm,desc string
	var ds,start,end time.Time

	for _,p := range exampleCalendar.Properties(){
		switch p.Name {
		case PropSummary:

			if s,err := p.GetToText();err != nil{
					t.Errorf("get summary value from prop err:%v",err)
			}else{
				sm = s
			}
		case PropDescription:
			if s,err := p.GetToText();err != nil{
				t.Errorf("get Description value from prop err:%v",err)
			}else{
				desc = s
			}
		case PropDatetimeStamp:
			if s,err := p.GetToDatetime();err != nil {
				t.Errorf("get datestap value from prop err:%v",err)
			}else {
				ds = s
			}
		case PropDatetimeStart:
			if s,err := p.GetToDatetime();err != nil{
				t.Errorf("get datetimestart value from prop err:%v",err)
			}else {
				start = s
			}
		case PropDatetimeEnd:
			if s,err := p.GetToDatetime();err != nil{
				t.Errorf("get datetimeend value from prop err:%v",err)
			}else {
				end = s
			}
		}

	}
	if sm != wantSummary {
		t.Errorf("Event summary  = %v, want %v", sm, wantSummary)
	}

	wantDesc := "Test Test Test event"
	if desc != wantDesc {
		t.Errorf("Event Description = %v, want %v", desc, wantDesc)
	}

	wantDTStamp := time.Date(1996, 07, 04, 12, 0, 0, 0, time.UTC)
	if ds != wantDTStamp {
		t.Errorf("Event PropDateTimeStamp = %v, want %v", ds, wantDTStamp)
	}

	wantDTStart := time.Date(1996, 9, 18, 14, 30, 0, 0, time.UTC)
	if start != wantDTStart {
		t.Errorf("Event DateTimeStart = %v, want %v", start, wantDTStart)
	}

	wantDTEnd := time.Date(1996, 9, 20, 22, 0, 0, 0, time.UTC)
	if end != wantDTEnd {
		t.Errorf("Event DateTimeEnd = %v, want %v", end, wantDTEnd)
	}
}

