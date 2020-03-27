package go_ical

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//=======================Parameters=====================
/*
one property can have multi Parameters,one Parameter is a key-value.
e.g:
ATTENDEE;RSVP=TRUE;ROLE=REQ-PARTICIPANT:mailto:
      jsmith@example.com

RSVP and ROLE are all belong to parameters

    icalparameter = altrepparam       ; Alternate text representation
                   / cnparam           ; Common name
                   / cutypeparam       ; Calendar user type
                   / delfromparam      ; Delegator
                   / deltoparam        ; Delegatee
                   / dirparam          ; Directory entry
                   / encodingparam     ; Inline encoding
                   / fmttypeparam      ; Format type
                   / fbtypeparam       ; Free/busy time type
                   / languageparam     ; Language for text
                   / memberparam       ; Group or list membership
                   / partstatparam     ; Participation status
                   / rangeparam        ; Recurrence identifier range
                   / trigrelparam      ; Alarm trigger relationship
                   / reltypeparam      ; Relationship type
                   / roleparam         ; Participation role
                   / rsvpparam         ; RSVP expectation
                   / sentbyparam       ; Sent by
                   / tzidparam         ; Reference to time zone object
                   / valuetypeparam    ; Property value data type
                   / other-param

     other-param   = (iana-param / x-param)

     iana-param  = iana-token "=" param-value *("," param-value)
     ; Some other IANA-registered iCalendar parameter.

     x-param     = x-name "=" param-value *("," param-value)
     ; A non-standard, experimental parameter.
 */

type Parameters map[string][]string

//all Parameters should implement this interface
type ParamItem interface {
	ParamItem(...interface{}) (string,[]string)
}


//used to store ONE parameter,and implement ParamKeyValues
type ParamItemObj struct {
	Key string
	Values []string
}

func (p *ParamItemObj) ParamItem(...interface{}) (string,[]string) {
	return p.Key,p.Values
}
func (p Parameters) GetItem(name string) ParamItem {
	if values := p[strings.ToUpper(name)];len(values)>0{
		return &ParamItemObj{
			Key:strings.ToUpper(name),
			Values:values,
		}
	}
	return nil
}

func (p Parameters) SetItem(pi ParamItem)  {
	k,vs := pi.ParamItem()
	p[strings.ToUpper(k)] = vs
}
//one ParamItem can have multi values,add one value to an existed ParamItem
func (p Parameters) AddItem(pi ParamItem)  {
	k,vs := pi.ParamItem()
	for _,v := range vs {
		p[strings.ToUpper(k)] = append(p[k], v)
	}
}

func (p Parameters) DelItem(pi ParamItem)  {
	k,_ := pi.ParamItem()
	delete(p,strings.ToUpper(k))
}

func (p Parameters) Get(name string) string {
	if values := p[strings.ToUpper(name)];len(values)>0{
		return values[0]
	}
	return ""
}

func (p Parameters) Set(k,v string)  {
	p[strings.ToUpper(k)] = []string{v}
}
//one ParamItem can have multi values,add one value to an existed ParamItem
func (p Parameters) Add(k,v string)  {
	p[strings.ToUpper(k)] = append(p[k], v)
}

func (p Parameters) Del(k string)  {
	delete(p,strings.ToUpper(k))
}

func NewParamItem(name string,values []string) ParamItem {
	return &ParamItemObj{
		Key:name,
		Values:values,
	}
}

func NewParamCN(cn string) ParamItem {
	return NewParamItem(Paramcn,[]string{cn})
}

func NewParamEncoding(enc string) ParamItem {
	return NewParamItem(Paramencoding,[]string{enc})
}

func NewParamFMTtype(formattype string) ParamItem {
	return NewParamItem(Paramfmttype,[]string{formattype})
}

func NewParamValue(val string) ParamItem {
	return NewParamItem(Paramvaluetypeparam,[]string{val})
}

func NewParamRSVP(rsvp bool) ParamItem {
	return NewParamItem(Paramrsvp,[]string{strconv.FormatBool(rsvp)})
}


//=============================Property======================================
/*
a property is a contentline in icalendar object
contentline   = name *(";" param ) ":" value CRLF
one name
none/one/multi parameters,parameters is key-value
one value
*/
type Property struct {
	Name string
	Params Parameters
	Value string
}

func NewProperty(name string) *Property {
	return &Property{
		Name:strings.ToUpper(name),
		Params:make(Parameters),
	}
}

//=========================Property Value Type===================================
//defined in RFC 5545 3.3
/*
The properties in an iCalendar object are strongly typed.  The
   definition of each property restricts the value to be one of the
   value data types, or simply value types, defined in this section.
   The value type for a property will either be specified implicitly as
   the default value type or will be explicitly specified with the
   "VALUE" parameter.  If the value type of a property is one of the
   alternate valid types, then it MUST be explicitly specified with the
   "VALUE" parameter.
 */


func (p *Property) GetParamValue() string {
	vdt := p.Params[Paramvaluetypeparam][0]
	if vdt == VDTdefault{
		vdt = DefaultVDT[p.Name]
	}
	return vdt
}

func (p *Property) UpdateParamValue(vdt string)  {
	nt,exist := DefaultVDT[p.Name]
	if nt == VDTdefault || (exist && nt == vdt){
		p.Params.Del(Paramvaluetypeparam)
	} else {
		p.Params.Set(Paramvaluetypeparam,vdt)
	}
}

func (p *Property) expectVDT(expect string) error {
	vdt := p.GetParamValue()
	if vdt != VDTdefault && vdt != expect{
		return fmt.Errorf("ical:property %q expect type %q,but got %q",p.Name,expect,vdt)
	}
	return nil
}
//TODO:only base64?
func (p *Property) GetToBinary() ([]byte,error) {
	if err := p.expectVDT(VDTbinary);err != nil {
		return nil,err
	}
	return base64.StdEncoding.DecodeString(p.Value)
}

func (p *Property) SetFromBinary(b []byte)  {
	p.UpdateParamValue(VDTbinary)
	p.Params.SetItem(NewParamEncoding("BASE64"))
	p.Value = base64.StdEncoding.EncodeToString(b)
}

func (p *Property) GetToBool() (bool,error) {
	if err := p.expectVDT(VDTbool);err != nil{
		return false,err
	}
	switch strings.ToUpper(p.Value) {
	case "FALSE":
		return false,nil
	case "TRUE":
		return true,nil
	default:
		return false,fmt.Errorf("ical:invalid bool: %q",p.Value)
	}
}

func (p *Property) GetToDate() (time.Time,error) {
	tz := p.Params.Get(Paramtzid)
	loc := time.UTC
	if tz != ""{
		l,err := time.LoadLocation(tz)
		if err != nil{
			return time.Time{},err
		}
		loc = l
	}
	vdt := p.GetParamValue()
	if vdt == VDTdefault || vdt == VDTdatetime{
		return time.ParseInLocation(DateFormat,p.Value,loc)
	} else {
		return time.Time{},fmt.Errorf("ical:expect date,but got %q",vdt)
	}
}


func (p *Property) GetToDatetime() (time.Time,error) {
	tz := p.Params.Get(Paramtzid)
	loc := time.UTC
	if tz != ""{
		l,err := time.LoadLocation(tz)
		if err != nil{
			return time.Time{},err
		}
		loc = l
	}
	vdt := p.GetParamValue()
	if vdt == VDTdefault || vdt == VDTdatetime{
		//TODO:need to know reason
		if t,err := time.ParseInLocation(DatetimeFormat,p.Value,loc);err == nil{
			return t,nil
		}
		return time.ParseInLocation(DatetimeFormat2,p.Value,time.UTC)
	} else {
		return time.Time{},fmt.Errorf("ical:expect datetime,but got %q",vdt)
	}
}

func (p *Property) SetFromDatetime(t time.Time)  {
	p.UpdateParamValue(VDTdatetime)
	p.Value = t.Format(DatetimeFormat2)
}

func (p *Property) SetFromDuration(d time.Duration)  {
	p.UpdateParamValue(VDTduration)
	seconds := d.Milliseconds()/1000
	sign := seconds < 0
	if seconds < 0{
		seconds = -seconds
	}
	var st string
	if sign {
		st += "-"
	}
	st += "PT"
	st += strconv.FormatInt(seconds,10)
	st += "S"
	p.Value = st
}

func (p *Property) GetToDuration() (time.Duration,error) {
	if err := p.expectVDT(VDTduration);err != nil {
		return 0,err
	}
	ds := durstr(p.Value)
	return ds.parseToDuration()
}

type durstr string
func (ds durstr) next(b byte) bool {
	if len(ds) ==0 || ds[0] != b{
		return false
	}
	ds = ds[1:]
	return true
}

func (ds durstr) parseToDuration() (time.Duration,error) {
	//Duration can date or datetime,so need to sigure
	isTime := false
	var end time.Duration
	sign := ds.next('-')
	if !sign {
		bl := ds.next('+')
		if bl{
			ds = ds[1:]
		} else {
			return 0,fmt.Errorf("ical: invalid duration,need '-' or '+'")
		}
	} else {
		ds = ds[1:]
	}


	if !ds.next('P'){
		return 0,fmt.Errorf("ical: invalide duration,expect 'P'")
	} else {
		ds = ds[1:]
	}

	for len(ds) > 0{
		if ds.next('T'){
			ds = ds[1:]
			isTime = true
		}

		index := strings.IndexFunc(string(ds), func(r rune) bool {
			return !(r>='0' && r<='9')
		})

		if index == 0 {
			return 0,fmt.Errorf("ical:invalid duration,shoud be digital")
		}
		if index < 0{
			index = len(ds)
		}

		n,err := strconv.ParseUint(string(ds[:index]),10,64)
		if err != nil{
			return 0,err
		}
		ds := ds[index:]

		num := time.Duration(n)

		if isTime{
			if ds.next('H'){
				ds = ds[1:]
				end += num * time.Hour
			} else if ds.next('M'){
				ds = ds[1:]
				end += num*time.Minute
			} else if ds.next('S'){
				ds = ds[1:]
				end += num*time.Second
			} else {
				return 0,fmt.Errorf("ical:invalid duration: expect 'H','M' or 'S'")
			}
		} else {
			if ds.next('D'){
				ds = ds[1:]
				end += num*24*time.Hour
			} else if ds.next('W'){
				ds = ds[1:]
				end += num*7*24*time.Hour
			} else {
				return 0,fmt.Errorf("ical:invalid duration: expect 'D' or 'W'")
			}
		}
	}

	if sign{
		end = -end
	}
	return end,nil
}

func (p *Property) GetToFloat() (float64,error) {
	if err := p.expectVDT(VDTfloat);err != nil{
		return 0,err
	}
	return strconv.ParseFloat(p.Value,64)
}

func (p *Property) GetToInt() (int,error) {
	if err := p.expectVDT(VDTint);err != nil{
		return 0,err
	}
	return strconv.Atoi(p.Value)
}

func (p *Property) GetToTextlines() ([]string,error) {
	if err := p.expectVDT(VDTtext);err != nil {
		return nil,err
	}
	var bui strings.Builder
	var ends []string
	n := len(p.Value)
	for i:=0;i<n;i++{
		switch c := p.Value[i];c {
		case '\\':
			i++
			if i >= n{
				return nil,fmt.Errorf("ical:text need end with '\\'")
			}
			switch c:=p.Value[i];c {
			case '\\',';',',':
				bui.WriteByte(c)
			case 'n','N':
				bui.WriteByte('\n')
			default:
				return nil,fmt.Errorf("ical:text invalid escape sequence '\\%v'",c)
			}
		case ',':
			ends = append(ends,bui.String())
			bui.Reset()
		default:
			bui.WriteByte(c)
		}
	}
	ends = append(ends,bui.String())
	return ends,nil
}

func (p *Property) SetFromTextlines(lines []string)  {
	p.UpdateParamValue(VDTtext)
	var bui strings.Builder
	for i,line := range lines{
		//seperate with ','
		if i > 0{
			bui.WriteByte(',')
		}

		bui.Grow(len(line))
		for _,c := range line{
			switch c {
			case '\\',';',',':
				bui.WriteByte('\\')
				bui.WriteRune(c)
			case '\n':
				bui.WriteString("\\n")
			default:
				bui.WriteRune(c)
			}
		}
	}
	p.Value = bui.String()
}

func (p *Property) GetToText() (string,error) {
	l,err := p.GetToTextlines()
	if err != nil{
		return "",err
	}
	if len(l) == 0 {
		return "",nil
	}
	return l[0],nil
}

func (p *Property) SetFromText(text string)  {
	p.SetFromTextlines([]string{text})
}
