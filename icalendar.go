package go_ical

type Calendar struct {
	ComponentObj
}

func NewCalendar() *Calendar {
	cal := &Calendar{
		ComponentObj{NameObj:CompCalendar,PropertiesObj:[]Property{},SubComponentsObj:[]Component{}},
	}
	cal.SetVersion("2.0")
	cal.SetProductId("-//xyz Corp//Scott WORK Calendar Version 1.0//CN")
	return cal
}

func (cal *Calendar) SetVersion(st string,pis ...ParamItem)  {
	cal.setProperty(PropVersion,ToText(st),pis...)
}

func (cal *Calendar) SetProductId(st string,pis ...ParamItem)  {
	cal.setProperty(PropProductIdentifier,ToText(st),pis...)
}

func (cal *Calendar) SetMethod(st string,pis ...ParamItem)  {
	cal.setProperty(PropMethod,ToText(st),pis ...)
}

func (cal *Calendar) SetDescription(st string,pis ...ParamItem)  {
	cal.setProperty(PropDescription,ToText(st),pis ...)
}

func (cal *Calendar) SetLastModified(st string,pis ...ParamItem)  {
	cal.setProperty(PropLastModified,ToText(st),pis ...)
}

func (cal *Calendar) SetProperty(pname string,val string,pis ...ParamItem)  {
	p := Property{Name:pname,Value:val}
	params := Parameters{}
	for _,pi := range pis{
		params.SetItem(pi)
	}
	p.Params = params
	cal.PropertiesObj = append(cal.PropertiesObj,p)
}

func (cal *Calendar) setProperty(pname string,val string,pis ...ParamItem)  {
	p := Property{Name:pname,Value:val}
	params := Parameters{}
	for _,pi := range pis{
		params.SetItem(pi)
	}
	p.Params = params
	cal.PropertiesObj = append(cal.PropertiesObj,p)
}

func (cal *Calendar) AddComponent(coms ...Component)  {
	cal.SubComponentsObj = append(cal.SubComponentsObj,coms ...)
}

func (cal *Calendar) GetEvents() []Component {
	events := make([]Component,0,len(cal.SubComponents()))
	for _,sub := range cal.SubComponents(){
		if sub.Name() == CompEvent{
			events = append(events,sub)
		}
	}
	return events
}
