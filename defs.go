package go_ical

//Parameters defined in RFC 5545 3.2
const (
	Paramaltrep = "ALTREP"
	Paramcn = "CN"
	Paramcutype = "CUTYPE"
	Paramdelfrom = "DELEGATED-FROMx"
	Paramdelto = "DELEGATED-TO"
	Paramdir = "DIR"
	Paramencoding = "ENCODING"
	Paramfmttype = "FMTTYPE"
	Paramfbtype = "FBTYPE"
	Paramlanguage = "LANGUAGE"
	Parammember = "MEMBER"
	Parampartstat = "PARTSTAT"
	Paramrange = "RANGE"
	Paramtrigrel = "RELATED"
	Paramreltype = "RELTYPE"
	Paramrole = "ROLE"
	Paramrsvp = "RSVP"
	Paramsentby = "SENT-BY"
	Paramtzid = "TZID"
	Paramvaluetypeparam = "VALUE"
)

//Value types defined in RFC 5545 3.3
const (
	VDTdefault = ""
	VDTbinary = "BINARY"
	VDTbool = "BOOL"
	VDTcalendaraddress = "CAL-ADDRESS"
	VDTdate = "DATE"
	VDTdatetime = "DATE-TIME"
	VDTduration = "DURATION"
	VDTfloat = "FLOAT"
	VDTint = "INTEGER"
	VDTperiod = "PERIOD"
	VDTrecurrence = "RECUR"
	VDTtext = "TEXT"
	VDTtime = "TIME"
	VDTuri = "URI"
	VDTutcoffset = "UTC-OFFSET"
)

//propertis defined in RFC 5545 3.7
const  (
	PropCalendarScale = "CALSCALE"
	PropMethod = "METHOD"
	PropProductIdentifier = "PRODID"
	PropVersion = "VERSION"
)
//component properties defined in RFC 3.8
const (
	//Descriptive Component Properties
	PropAttachment = "ATTACH"
	PropCategories = "CATEGORIES"
	PropClassification = "CLASS"
	PropComment = "COMMENT"
	PropDescription = "DESCRIPTION"
	PropGeographicPosition = "GEO"
	PropLocation = "LOCATION"
	PropPercentComplete = "PERCENT-COMPLETE"
	PropPriority = "PRIORITY"
	PropResources = "RESOURCES"
	PropStatus = "STATUS"
	PropSummary = "SUMMARY"

	//Date and Time Component Properties
	PropDatetimeCompleted = "COMPLETED"
	PropDatetimeEnd = "DTEND"
	PropDatetimeDue = "DUE"
	PropDatetimeStart = "DTSTART"
	PropDuration = "DURATION"
	PropFreeBusy = "FREEBUSY"
	PropTimeTransparency = "TRANSP"

	//Time Zone Component Properties
	PropTimeZoneIdentifier = "TZID"
	PropTimeZoneName = "TZNAME"
	PropTimeZoneOffsetFrom = "TZOFFSETFROM"
	PropTimeZoneOffsetTo = "TZOFFSETTO"
	PropTimeZoneURL = "TZURL"

	//Relationship Component Properties
	PropAttendee = "ATTENDEE"
	PropContact = "CONTACT"
	PropOrganizer = "ORGANIZER"
	PropRecurrenceId = "RECURRENCE-ID"
	PropRelatedTo = "RELATED-TO"
	PropURL = "URL"
	PropUID = "UID"

	//Recurrence Component Properties
	PropExceptionDatetime = "EXDATE"
	PropRecurrenceDatetime = "RDATE"
	PropRecurrenceRule = "RRULE"

	//Alarm Component Properties
	PropAction = "ACTION"
	PropRepeatCount = "REPEAT"
	PropTrigger = "TRIGGER"

	//Change Management Component Properties
	PropDatetimeCreated = "CREATED"
	PropDatetimeStamp = "DTSTAMP"
	PropLastModified = "LAST-MODIFIED"
	PropSequenceNumber = "SEQUENCE"

	//Miscellaneous Component Properties
	//IANA Properties
	//TODO:
	/*
	Property Name:  An IANA-registered property name

	   Value Type:  The default value type is TEXT.  The value type can be
	      set to any value type.
	*/

	//Non-Standard Properties
	//TODO:
	/*
	Property Name:  Any property name with a "X-" prefix

	   Purpose:  This class of property provides a framework for defining
	      non-standard properties.

	   Value Type:  The default value type is TEXT.  The value type can be
	      set to any value type.

	   Property Parameters:  IANA, non-standard, and language property
	      parameters can be specified on this property.

	   Conformance:  This property can be specified in any calendar
	      component.
	 */
	//

	//Request Status
	PropRequestStatus = "REQUEST-STATUS"
/*
   Property Name:  REQUEST-STATUS

      Purpose:  This property defines the status code returned for a
         scheduling request.

      Value Type:  TEXT

      Property Parameters:  IANA, non-standard, and language property
         parameters can be specified on this property.

      Conformance:  The property can be specified in the "VEVENT", "VTODO",
         "VJOURNAL", or "VFREEBUSY" calendar component.
 */
)

//property default vaule data type
var DefaultVDT = map[string]string{
	PropCalendarScale:VDTtext,
	PropMethod:VDTtext,
	PropProductIdentifier:VDTtext,
	PropVersion:VDTtext,
	PropAttachment:VDTuri,// can use binary
	PropCategories:VDTtext,
	PropClassification:VDTtext,
	PropComment:VDTtext,
	PropDescription:VDTtext,
	PropGeographicPosition:VDTfloat,
	PropLocation:VDTtext,
	PropPercentComplete:VDTint,
	PropPriority:VDTint,
	PropResources:VDTtext,
	PropStatus:VDTtext,
	PropSummary:VDTtext,
	PropDatetimeCompleted:VDTdatetime,
	PropDatetimeEnd:VDTdatetime,//can use date
	PropDatetimeDue:VDTdatetime,//can use date
	PropDatetimeStart:VDTdatetime,//can use date
	PropDuration:VDTduration,
	PropFreeBusy:VDTperiod,
	PropTimeTransparency:VDTtext,
	PropTimeZoneIdentifier:VDTtext,
	PropTimeZoneName:VDTtext,
	PropTimeZoneOffsetFrom:VDTutcoffset,
	PropTimeZoneOffsetTo:VDTutcoffset,
	PropTimeZoneURL:VDTuri,
	PropAttendee:VDTcalendaraddress,
	PropContact:VDTtext,
	PropOrganizer:VDTcalendaraddress,
	PropRecurrenceId:VDTdatetime,
	PropRelatedTo:VDTtext,
	PropURL:VDTuri,
	PropUID:VDTtext,
	PropExceptionDatetime:VDTdatetime,//can use date
	PropRecurrenceDatetime:VDTdatetime,//can use date or period
	PropRecurrenceRule:VDTrecurrence,
	PropAction:VDTtext,
	PropRepeatCount:VDTint,
	PropTrigger:VDTduration,//can use datetime
	PropDatetimeCreated:VDTdatetime,
	PropDatetimeStamp:VDTdatetime,
	PropLastModified:VDTdatetime,
	PropSequenceNumber:VDTint,
	PropRequestStatus:VDTtext,
}

const (
	DateFormat = "20060102"
	DatetimeFormat = "20060102T150405"
	DatetimeFormat2 = "20060102T150405Z"
)

//RFC 5545 3.6
const (
	CompCalendar = "VCALENDAR"
	CompEvent = "VEVENT"
	CompTodo = "VTODO"
	CompJournal = "VJOURNAL"
	CompFreebusy = "VFREEBUSY"
	CompTimezone = "VTIMEZONE"
	CompTimezoneStandard = "STANDARD"
	CompTimezoneDaylight = "DAYLIGHT"
	CompAlarm = "VALARM"


)

var OnlyOnePropMap = map[string][]string{
	CompCalendar:[]string{PropProductIdentifier,PropVersion},
	CompEvent:[]string{PropDatetimeStamp,PropUID},
	CompTodo:[]string{PropDatetimeStamp,PropUID},
	CompJournal:[]string{PropDatetimeStamp,PropUID},
	CompFreebusy:[]string{PropDatetimeStamp,PropUID},
	CompTimezone:[]string{PropTimeZoneIdentifier},
	CompTimezoneStandard:[]string{PropDatetimeStart,PropTimeZoneOffsetTo,PropTimeZoneOffsetFrom},
	CompTimezoneDaylight:[]string{PropDatetimeStart,PropTimeZoneOffsetTo,PropTimeZoneOffsetFrom},
	CompAlarm:[]string{PropAction,PropTrigger},
}

var OneOrZeroPropMap = map[string][]string{
	CompCalendar:[]string{PropCalendarScale,PropMethod},
	CompEvent:[]string{PropDatetimeStart,PropClassification,PropDatetimeCreated,PropDescription,PropGeographicPosition,PropLastModified,PropLocation,
		PropOrganizer,PropPriority,PropSequenceNumber,PropStatus,PropSummary,PropTimeTransparency,PropURL,PropRecurrenceId,PropDatetimeEnd,PropDuration},
	CompTodo:[]string{PropClassification,PropDatetimeCompleted,PropDatetimeCreated,PropDescription,PropDatetimeStart,PropGeographicPosition,
		PropLastModified,PropLocation,PropOrganizer,PropPercentComplete,PropPriority,PropRecurrenceId,PropSequenceNumber,PropStatus,PropSummary,PropURL,PropDatetimeDue,PropDuration},
	CompJournal:[]string{PropClassification,PropDatetimeCreated,PropDatetimeStart,PropLastModified,PropOrganizer,PropRecurrenceId,PropSequenceNumber,PropStatus,PropSummary,PropURL},
	CompFreebusy:[]string{PropContact,PropDatetimeStart,PropDatetimeEnd,PropOrganizer,PropURL},
	CompTimezone:[]string{PropLastModified,PropTimeZoneURL},
	//alarm is special
}
