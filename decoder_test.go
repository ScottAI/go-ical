package go_ical

/*
func TestDecodeProperty(t *testing.T) {
	tests := []struct {
		Input    string
		Expected func(output *Property) bool
	}{
		{Input: "ATTENDEE;RSVP=TRUE;ROLE=REQ-PARTICIPANT;CUTYPE=GROUP:mailto:employee-A@example.com", Expected: func(output *Property) bool {
			return output.Name == "ATTENDEE" && output.Value == "mailto:employee-A@example.com"
		}},
		{Input: "ATTENDEE;RSVP=\"TRUE\";ROLE=REQ-PARTICIPANT;CUTYPE=GROUP:mailto:employee-A@example.com", Expected: func(output *Property) bool {
			return output.Name == "ATTENDEE" && output.Value == "mailto:employee-A@example.com"
		}},
		{Input: "ATTENDEE;RSVP=T\"RUE\";ROLE=REQ-PARTICIPANT;CUTYPE=GROUP:mailto:employee-A@example.com", Expected: func(output *Property) bool { return output == nil }},
	}
	for i, test := range tests {
		output,err := decodeProperty(ContentLine(test.Input))
		if err != nil{
			t.Logf("err: %#v",err)
		}
		if !test.Expected(output) {
			t.Logf("Got: %#v", output)
			t.Logf("Failed %d %#v", i, test)
			t.Fail()
		}
	}
}

 */


/*
func TestDecoder_Decode(t *testing.T) {
	dec := NewDecoder(strings.NewReader(exampleCalendarStr))

	cal,err := dec.Decode()
	if err != nil{
		t.Fatalf("DecodeCal err:%v",err)
	}

	if !reflect.DeepEqual(cal,exampleCalendar){
		t.Errorf("Decode() got \n %v,but expect \n %v",cal,exampleCalendar)
	}

	if _,err := dec.Decode();err != io.EOF{
		t.Errorf("Decode() = %v,want io.EOF ",err)
	}
}


 */
