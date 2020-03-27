package go_ical

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"
)

type ContentLine string

func decodeProperty(line ContentLine) (*Property,error) {
	p := &Property{
		Params: map[string][]string{},
	}

	proNameReg, err := regexp.Compile("[A-Za-z0-9-]{1,}")
	if err != nil {
		return nil,fmt.Errorf("ical:Failed to build regex: %v", err)
	}
	namepos := proNameReg.FindIndex([]byte(line))
	if namepos == nil {
		return nil,fmt.Errorf("ical:property name format error or no property name")
	}
	n := 0
	p.Name = string(line[n+namepos[0]:n+namepos[1]])
	n += namepos[1]
	for{
		if n>=len(line){
			return nil,fmt.Errorf("ical:content line format error,can not only have property name")
		}
		switch rune(line[n]) {
		case ':':
			return decodePropertyValue(p,line,n+1)
		case ';':
			if _,err,newn := decodePropertyParam(p,line,n+1);err != nil{
				return nil,err
			}else {
				n = newn
			}
		default:
			return nil,fmt.Errorf("ical:content line format error,after property name,VALUE and PARAMETERS are allowed")
		}
	}
}

func decodePropertyValue(p *Property,line ContentLine,n int) (*Property,error) {
	paramValueReg, err := regexp.Compile("^(?:\"(?:[^\"\\\\]|\\[\"nrt])*\"|[^,;\\\\:\"]*)")
	if err != nil {
		return nil,fmt.Errorf("ical: Failed to build regex: %v", err)
	}
	vpos := paramValueReg.FindIndex([]byte(line[n:]))
	if vpos == nil {
		return nil,fmt.Errorf("ical:Can not decode PropertyValue:format may error")
	}
	p.Value = string(line[n:n+vpos[1]])
	return p,nil
}

func decodePropertyParam(p *Property,line ContentLine,n int) (*Property,error,int) {
	paraNameReg, err := regexp.Compile("[A-Za-z0-9-]{1,}")
	if err != nil {
		return nil,fmt.Errorf("Failed to build regex: %v", err),-1
	}
	npos := paraNameReg.FindIndex([]byte(line[n:]))
	if npos == nil {
		return nil,fmt.Errorf("ical:can not decode property parameter name,no parameter name or format error"),-1
	}
	key := string(line[n:n+npos[1]])
	n += npos[1]
	if rune(line[n]) != '='{
		return nil,fmt.Errorf("ical:param format need '='"),-1
	}
	paraValueReg,err :=  regexp.Compile("^(?:\"(?:[^\"\\\\]|\\[\"nrt])*\"|[^,;\\\\:\"]*)")
	if err != nil {
		return nil,fmt.Errorf("Failed to build regex: %v", err),-1
	}
	for{
		if n >= len(line){
			return nil,fmt.Errorf("ical:property format error,need value"),-1
		}
		npos = paraValueReg.FindIndex([]byte(line[n:]))
		if npos == nil{
			return nil,fmt.Errorf("ical:can not decode property param,param value format error or no value"),-1
		}
		val := string(line[n+npos[0]:n+npos[1]])
		n += npos[1]
		p.Params[key]=append(p.Params[key],val)
		if rune(line[n]) == ','{
			n += 1
		}else {
			return p,nil,n
		}
	}
}

type Decoder struct {
	r *bufio.Reader
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{bufio.NewReader(r)}
}

func (dec *Decoder) Decode() (Component,error) {
	com,err := dec.decodeComponent()
	return com,err
}

func (dec *Decoder) decodeComponent() (*ComponentObj,error) {
	firstProp,err := dec.decodeContentline()
	if err != nil{
		return nil,err
	}
	if firstProp.Name != "BEGIN"{
		return nil,fmt.Errorf("ical:malformed component,expect BEGIN property,but got %q",firstProp.Name)
	}



	return dec.generalDecodeComponent(firstProp)
}

func (dec *Decoder) generalDecodeComponent(first *Property) (*ComponentObj,error) {
	var prop *Property
	props := []Property{}
	var subComs = []Component{}
	isContinued := true
	for ;isContinued;{
		var err error
		prop,err = dec.decodeContentline()
		if err != nil{
			return nil,err
		}
		switch prop.Name {
		case "END":
			if prop.Value != first.Value{
				return nil,fmt.Errorf("ical:malformed component,expect END property %q,but got %q",first.Value,prop.Value)
			}else{
				isContinued = false
				break
			}
		case "BEGIN":
			sub,err := dec.generalDecodeComponent(prop)
			if err != nil{
				return nil,err
			}
			subComs = append(subComs,sub)
		default:
			props = append(props,*prop)
		}
	}
	if prop.Name != "END"{
		return nil,fmt.Errorf("ical:expect END property")
	}
	return &ComponentObj{
		NameObj:first.Value,
		PropertiesObj:props,
		SubComponentsObj:subComs,
	},nil
}

func (dec *Decoder) decodeContentline() (*Property,error) {
	for{
		line,err := dec.readContentline()
		if err != nil{
			return nil,err
		}
		if len(line) == 0{
			continue
		}
		return decodeProperty(line)
	}
}

func (dec *Decoder) readContentline() (ContentLine,error) {
	bs,err := dec.r.ReadSlice('\n')
	if err == io.EOF && len(bs)>0{
		err = nil
	}
	if err != nil{
		return "",err
	}
	bs = bytes.TrimRight(bs,"\r\n")
	var sb strings.Builder

	sb.Write(bs)
	for{
		c,_,err := dec.r.ReadRune()
		if err == io.EOF{
			break
		} else if err != nil {
			return "",err
		}

		if c != ' ' && c != '\t'{
			dec.r.UnreadRune()
			break
		}

		bs,err = dec.r.ReadSlice('\n')
		if err == io.EOF && len(bs)>0{
			err = nil
		}
		if err != nil{
			return "",err
		}
		bs = bytes.TrimRight(bs,"\r\n")
		sb.Write(bs)

	}
	return ContentLine(sb.String()),nil
}
