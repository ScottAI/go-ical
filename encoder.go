package go_ical

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"unicode/utf8"
)

type Encoder struct {
	w io.Writer
}

func (enc *Encoder) encodeProperty(p *Property)  {
	b := bytes.NewBufferString("")
	fmt.Fprintln(b,p.Name)
	for key,vals := range p.Params{
		fmt.Fprint(b,';')
		fmt.Fprint(b,key)
		fmt.Fprint(b,'=')
		//process values
		for i,val := range vals{
			//seperate with ','
			if i > 0{
				fmt.Fprint(b,',')
			}
			if strings.ContainsAny(val,";:\\\","){
				val = strings.Replace(val,"\"","\\\"",-1)
				val = strings.Replace(val,"\\","\\\\",-1)
			}
			fmt.Fprint(b,val)
		}
	}
	fmt.Fprint(b,":")
	fmt.Fprint(b,p.Value)
	st := b.String()
	//https://tools.ietf.org/html/rfc5545#section-3.1
	if len(st) > 75 {
		for len(st) > 74{
			ts := trimUTF8(74,st)
			fmt.Fprint(enc.w,ts,"\r\n")//CRLF CR and LF
			fmt.Fprint(enc.w," ")
			st = st[len(ts):]
		}
	}
	fmt.Fprint(enc.w,st,"\r\n")
}

func trimUTF8(max int,st string) string {
	l := 0
	for _,c := range st {
		newl := l + utf8.RuneLen(c)
		if newl > max {
			break
		}
		l = newl
	}
	return st[:l]
}
/*
//This method lead to cmp be copied ,but I don't know how to use index with Interface,so chang Encoder's Method to Component's method
func (enc *Encoder) endcodeComponent(cmp Component) error {
	if err := cmp.IsAvailable();err != nil{
		return err
	}
	fmt.Fprint(enc.w,"BEGIN:"+cmp.Name(),"\r\n")
	for _,p := range cmp.Properties(){
		enc.encodeProperty(&p)
	}

	for _,c := range cmp.SubComponents(){
		enc.endcodeComponent(c)
	}
	fmt.Fprint(enc.w,"END:"+cmp.Name()+"\r\n")

	return nil
}

 */

func (enc *Encoder) Encode(com Component) error {
	if err := com.encode(enc);err != nil{
		return err
	}
	return nil
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w}
}
