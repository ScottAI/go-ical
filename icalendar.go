package go_ical

type Calendar struct {
	ComponentObj
}

func NewCalendar() *Calendar {
	cal := &Calendar{
		ComponentObj{NameObj:CompCalendar,PropertiesObj:[]Property{},SubComponentsObj:[]Component{}},
	}
	
}

func (cal *Calendar) SetVersion(st string,pis ...ParamItem)  {

}

func (cal *Calendar) setProperty(prop Property,val string,pis ...ParamItem)  {

}
