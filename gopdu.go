package gopdu

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"
)

var p = fmt.Println

type SMS map[any]SMSone

type SMSone struct {
	Pdu                        string `json:"pdu"`
	ServiceCenterAddressLength int
	ServiceCenterAddressType   string
	ServiceCenterAddress       string
	MessageType                string
	ReplyPath                  string
	UserDataHeaderIncluded     string
	StatusReportRequest        string
	ValidityPeriodFormat       string
	RejectDuplicates           string
	MessageTypeIndicator       string
	MessageReference           string
	AddressLength              int
	AddressType                string
	Address                    string
	ProtocolIdentifier         string
	DataCodingScheme           string
	DataCodingSchemeIs7        bool
	ValidityPeriod             int
	ValidityPeriodSecond       int
	UserDataLength             string
	UserData                   string
	FirstSymbol                string
	Text                       string `json:"text"`
	Date                       string `json:"date"`
	Parts                      int    `json:"parts"`
	Part                       int    `json:"part"`
	Point                      int
	UserDataHeaderLength       int
	UserDataHeader             string
}

// randomString(int) string
// return random value. Double random
func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		switch randInt(1, 7) {
		case 1:
			bytes[i] = byte(randInt(48, 57))
		case 2:
			bytes[i] = byte(randInt(65, 90))
		case 3:
			bytes[i] = byte(randInt(97, 122))
		default:
			bytes[i] = byte(randInt(97, 122))
		}
	}
	return string(bytes)
}

// randInt(int64, int64) int
// return random value
func randInt(min, max int) int {
	r, _ := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	if r.Int64() < int64(min) {
		return int(r.Int64() + int64(min))
	} else {
		return int(r.Int64())
	}
}

// getPart([]string, int, bool, *int) string
// get part of []string by first point and replace symbol "f" if needed
func getPart(arr []string, add int, rep_f bool, point *int) string {
	min := *point
	max := min + add
	if len(arr) < max {
		max = len(arr) - 1
	}
	*point = max
	if rep_f {
		return strings.Replace(strings.Join(arr[min:max], ""), "f", "", -1)
	}
	return strings.Join(arr[min:max], "")
}

// castling(string) string
// castling of odd and even
func castling(s string) string {
	arr_forward := ""
	for k, _ := range s {
		if (k+1)%2 == 0 {
			arr_forward += string(s[k]) + string(s[k-1])
		}
	}
	return arr_forward
}

// decTo(string, int) string
// digital convert from decimal
func decTo(s string, base int) string {
	toi, _ := strconv.Atoi(s)
	return strconv.FormatInt(int64(toi), base)
}

// toDec(string, int) string
// digital convert to decimal
func toDec(i string, base int) string {
	ret, _ := strconv.ParseInt(i, base, 32)
	return strconv.Itoa(int(ret))
}

// toCapacity(interface{}, int, int, int) string
// to any capacity
func toCapacity(val interface{}, fromBase int, toBase int, count int) string {
	prfx := func(value string) string {
		prefix := ""
		if len(value) < count {
			prefix = strconv.Itoa(int(math.Pow(10, float64(count-len(value)))))[1:]
		}
		return prefix
	}
	ret := ""
	switch val.(type) {
	case string:
		def := val.(string)
		r := decTo(toDec(def, fromBase), toBase)
		ret = prfx(r) + r
	case int:
		def := val.(int)
		v := strconv.Itoa(def)
		r := decTo(toDec(v, fromBase), toBase)
		ret = prfx(r) + r
	case uint8:
		def := val.(uint8)
		v := strconv.Itoa(int(def))
		r := decTo(toDec(v, fromBase), toBase)
		ret = prfx(r) + r
	case int16:
		def := val.(int16)
		v := strconv.Itoa(int(def))
		r := decTo(toDec(v, fromBase), toBase)
		ret = prfx(r) + r
	case []byte:
		def := val.([]byte)
		r := decTo(toDec(string(def), fromBase), toBase)
		ret = prfx(r) + r
	case []string:
		def := val.([]string)
		newsrt := ""
		for _, v := range def {
			r := decTo(toDec(v, fromBase), toBase)
			newsrt += prfx(r) + r
		}
		ret = newsrt
	case []int:
		def := val.([]int)
		newsrt := ""
		for _, v := range def {
			v := strconv.Itoa(v)
			r := decTo(toDec(v, fromBase), toBase)
			newsrt += prfx(r) + r
		}
		ret = newsrt
	default:
		fmt.Println("Error! ToStr")
	}
	return ret
}

// splitArray(string, int, bool) map[int]string
// split slice by parameter
func splitArray(text string, count int, filling bool) map[int]string {
	r := make(map[int]string)
	for k, v := range text {
		r[k/count] += string(v)
	}
	if filling {
		for kk, vv := range r {
			if len(vv) < count {
				r[kk] = vv + strings.Repeat("0", count-len(vv))
			}
		}
	}
	return r
}

// joinArray(arr0 map[int]string)
// join slice
func joinArray(arr0 map[int]string) string {
	r := ""
	for i := 0; i < len(arr0); i++ {
		r = r + arr0[i]
	}
	return r
}

// trLine(string) string
// transform line to reverse
func trLine(s string) string {
	arr_forward := ""
	for k, _ := range s {
		arr_forward += s[len(s)-k-1 : len(s)-k]
	}
	return arr_forward
}

// trSlice([]string) []string
// transform slice to reverse
func trSlice(arr []string) []string {
	ret := make([]string, len(arr))
	for k, v := range arr {
		ret[len(arr)-1-k] = v
	}
	return ret
}

// trMap(map[int]string) map[int]string
// transform map to reverse
func trMap(r map[int]string) map[int]string {
	arr1 := make(map[int]string)
	for k, v := range r {
		arr1[len(r)-k-1] = string(v)
	}
	return arr1
}

// toStr(values ...interface{}) string
// quick 'to string'
func toStr(values ...interface{}) string {
	var retsrt string
	for _, value := range values {
		if value != nil {
			switch value.(type) {
			case string:
				retsrt += value.(string)
			case bool:
				retsrt += strconv.FormatBool(value.(bool))
			case int:
				retsrt += strconv.Itoa(value.(int))
			case uint8:
				retsrt += strconv.Itoa(int(value.(uint8)))
			case int16:
				retsrt += strconv.Itoa(int(value.(int16)))
			case []byte:
				retsrt += string(value.([]byte))
			case error:
				retsrt += value.(error).Error()
			case []string:
				newsrt := ""
				for _, v := range value.([]string) {
					newsrt += v
				}
				retsrt += newsrt
			case []int:
				newsrt := ""
				for _, v := range value.([]int) {
					newsrt += strconv.Itoa(v)
				}
				retsrt += newsrt
			default:
				fmt.Println("Error! ToStr")
			}
		}
	}
	return retsrt
}

// toInt(values ...interface{}) int
// quick 'to int'
func toInt(values ...interface{}) int {
	var retsrt int
	for _, value := range values {
		if value != nil {
			switch value.(type) {
			case string:
				tmp, _ := strconv.Atoi(value.(string))
				retsrt += tmp
			case bool:
				if value.(bool) {
					retsrt += 1
				} else {
					retsrt += 0
				}
			case int:
				retsrt += value.(int)
			case uint8:
				retsrt += int(value.(uint8))
			case int16:
				retsrt += int(value.(int16))
			case []byte:
				tmp, _ := strconv.Atoi(string(value.([]byte)))
				retsrt += tmp
			case []string:
				newsrt := 0
				for _, v := range value.([]string) {
					tmp, _ := strconv.Atoi(v)
					newsrt += tmp
				}
				retsrt += newsrt
			case []int:
				newsrt := 0
				for _, v := range value.([]int) {
					newsrt += v
				}
				retsrt += newsrt
			default:
				fmt.Println("Error! ToInt")
			}
		}
	}
	return retsrt
}

// Decode(string)
// decode pdu string
func (obj SMS) Decode(pdu string) {
	o := SMSone{}
	o.Pdu = pdu
	o.Point = 0
	decoded, _ := hex.DecodeString(pdu)
	Forward, Return, Binary := make([]string, 0), make([]string, 0), make([]string, 0)
	for _, v := range decoded {
		hexline := hex.EncodeToString([]byte{v})
		Forward = append(Forward, hexline)
		Return = append(Return, hexline[1:2]+hexline[0:1])
		Binary = append(Binary, toStr(toInt(strconv.FormatInt(int64(v), 2)) + 100000000)[1:])
	}
	p(o.Point)
	o.ServiceCenterAddressLength = toInt(toCapacity(getPart(Forward, 1, false, &o.Point), 16, 10, 2))
	if o.ServiceCenterAddressLength > 0 {
		o.ServiceCenterAddressLength = o.ServiceCenterAddressLength - 1
		o.ServiceCenterAddressType = getPart(Forward, 1, false, &o.Point)
		o.ServiceCenterAddress = getPart(Return, o.ServiceCenterAddressLength, true, &o.Point)
	}
	o.MessageType = getPart(Binary, 1, false, &o.Point)
	o.ReplyPath = o.MessageType[0:1]
	o.UserDataHeaderIncluded = o.MessageType[1:2]
	o.StatusReportRequest = o.MessageType[2:3]
	o.ValidityPeriodFormat = o.MessageType[3:5]
	o.RejectDuplicates = o.MessageType[5:6]
	o.MessageTypeIndicator = o.MessageType[6:8]
	if o.MessageTypeIndicator == "01" {
		o.MessageReference = toCapacity(getPart(Forward, 1, false, &o.Point), 16, 10, 2)
	}
	o.AddressLength = toInt(toCapacity(getPart(Forward, 1, false, &o.Point), 16, 10, 2))
	o.AddressType = getPart(Forward, 1, false, &o.Point)
	o.Address = getPart(Return, toInt((o.AddressLength+1)/2), true, &o.Point)
	o.ProtocolIdentifier = getPart(Forward, 1, false, &o.Point)
	o.DataCodingScheme = getPart(Binary, 1, false, &o.Point)
	o.DataCodingSchemeIs7 = o.DataCodingScheme[4:6] == "00"
	if o.MessageTypeIndicator == "00" {
		obj_DateSending, _ := time.Parse("060102150405", getPart(Return, 7, false, &o.Point)[:12])
		o.Date = obj_DateSending.Format("2006-01-02 15:04:05")
	} else {
		if o.ValidityPeriodFormat == "10" {
			o.ValidityPeriod = toInt(toCapacity(getPart(Forward, 1, false, &o.Point), 16, 10, 2))
			switch {
			case o.ValidityPeriod > 0 && o.ValidityPeriod <= 143:
				o.ValidityPeriodSecond = o.ValidityPeriod * 5 * 60
			case o.ValidityPeriod > 144 && o.ValidityPeriod <= 167:
				o.ValidityPeriodSecond = (o.ValidityPeriod-143)*30*60 + (12 * 60 * 60)
			case o.ValidityPeriod > 168 && o.ValidityPeriod <= 196:
				o.ValidityPeriodSecond = (o.ValidityPeriod - 166) * 24 * 60 * 60
			case o.ValidityPeriod > 197 && o.ValidityPeriod <= 255:
				o.ValidityPeriodSecond = (o.ValidityPeriod - 192) * 24 * 60 * 60 * 7
			}
		}
		if o.ValidityPeriodFormat == "11" {
			o.ValidityPeriod = toInt(toCapacity(getPart(Forward, 7, false, &o.Point), 16, 10, 2))
		}
	}
	o.UserDataLength = toCapacity(getPart(Forward, 1, false, &o.Point), 16, 10, 2)
	if o.UserDataHeaderIncluded == "1" {
		o.UserDataHeaderLength = toInt(toCapacity(getPart(Forward, 1, false, &o.Point), 16, 10, 2))
		if o.DataCodingSchemeIs7 && o.UserDataHeaderLength%2 == 1 {
			o.UserDataHeaderLength++
		}
		o.UserDataHeader = getPart(Forward, int(o.UserDataHeaderLength), false, &o.Point)
		o.MessageReference = toCapacity(o.UserDataHeader[4:6], 16, 10, 2)
		o.Parts = toInt(o.UserDataHeader[6:8])
		if len(o.UserDataHeader) != 8 {
			o.Part = toInt(o.UserDataHeader[8:10])
		}
		if o.DataCodingSchemeIs7 {
			o.FirstSymbol = string(rune(toInt(toCapacity(o.UserDataHeader[10:12], 16, 10, 2)) / 2))
		}
	}
	if o.DataCodingSchemeIs7 {
		o.UserData = strings.Join(Forward[o.Point:], "")
		BinaryUserData := strings.Join(trSlice(Binary[o.Point:]), "")
		Decoded := []string{}
		for i := len(BinaryUserData); i > 0; i = i - 7 {
			if i >= 7 {
				Decoded = append(Decoded, BinaryUserData[i-7:i])
			} else {
				Decoded = append(Decoded, BinaryUserData[:i])
			}
		}
		for _, v := range Decoded {
			val, _ := strconv.ParseInt(v, 2, 32) //32
			if val > 0 {
				o.Text += string(rune(val))
			}
		}
	} else {
		Decoded := Forward[o.Point:]
		for i := 0; i <= len(Decoded)-1; i = i + 2 {
			val, _ := strconv.ParseInt(strings.Join(Decoded[i:i+2], ""), 16, 32)
			if val > 0 {
				o.Text += string(rune(val))
			}
		}
	}
	o.Text = o.FirstSymbol + o.Text
	obj[o.MessageReference+o.MessageType+o.Address+toStr(o.Parts)+toStr(o.Part)] = o
}

// MergeTextToFirst()
// merge multipart sms text into first
func (obj *SMS) MergeTextToFirst() {
	for _, o := range *obj {
		if o.Part == 1 {
			field := o.MessageReference + o.MessageType + o.Address + toStr(o.Parts)
			for i := 2; i <= o.Parts; i++ {
				if ok := (*obj)[field+toStr(i)]; ok != (SMSone{}) {
					o.Text += ok.Text
				}
			}
			(*obj)[field+toStr(o.Part)] = o
		}
	}
}

// Encode(string, string) SMS
// encode to pdu
func Encode(text string, phone string) SMS {
	DataCodingSchemeIs7 := true
	MaxLength := 160
	Alphabet := "00"
	for _, v := range text {
		if int(v) > 255 {
			DataCodingSchemeIs7 = false
			Alphabet = "10"
			MaxLength = 67
		}
	}

	sms := make(SMS)
	chunks := splitArray(text, MaxLength, false)
	Ref := toCapacity(randInt(1, 255), 10, 16, 2)
	Ref2 := toCapacity(randInt(1, 65535), 10, 16, 2)
	items := make(SMS)
	for k, v := range chunks {
		item := SMSone{}
		item.Address = phone
		item.Text = v
		item.Part = k + 1
		item.Parts = len(chunks)
		items[k] = item
	}

	for _, v := range items {
		// --> Service Center Address (ServiceCenterAddress)
		// 0 - use default ServiceCenterAddress.
		v.ServiceCenterAddress = toStr(0)
		v.Pdu += toCapacity(v.ServiceCenterAddress, 10, 16, 2)
		// --> Protocol Data Unit Type (PDU Type)
		// >0<0000000 Reply Path (RP)
		// 0–not use, 1–use
		v.ReplyPath = "0"
		// 0>0<000000 Header in UD (UDH)
		// 0–only data, 1-header and data.
		v.UserDataHeaderIncluded = "0"
		if v.Parts > 1 {
			v.UserDataHeaderIncluded = "1"
		}
		// 00>0<00000 Status Report Request (SRR)
		// 0-not use, 1-use
		v.StatusReportRequest = "1"
		// 000>00<000 Validity Period Format (VPF)
		// 00 – empty VP
		// 01 – for Siemens,
		// 10 – VP standart
		// 11 – VP avsolute
		v.ValidityPeriodFormat = "10"
		// 00000>0<00 Reject Duplicates (RD)
		// --> 0–not delete, 1-delete
		v.RejectDuplicates = "0"
		// 000000>00< Message Type Indicator (MTI)
		// 00  SMS-DELIVER   REPORT SMS-DELIVER
		// 10  SMS-COMMAND   SMS-STATUS REPORT
		// 01  SMS-SUBMIT    SMS-SUBMIT REPORT
		// 11  RESERVED
		v.MessageTypeIndicator = "01"
		v.Pdu += toCapacity(v.ReplyPath+v.UserDataHeaderIncluded+v.StatusReportRequest+v.ValidityPeriodFormat+v.RejectDuplicates+v.MessageTypeIndicator, 2, 16, 2)
		// --> Message Reference (MR)
		v.MessageReference = Ref
		v.Pdu += v.MessageReference
		// --> Destination Address (DA)
		AddressExt := ""
		v.Address = "7" + strings.Replace(v.Address[len(v.Address)-10:], "+", "", -1)
		if len(v.Address)%2 > 0 {
			AddressExt = "F"
		}
		v.Address = castling(v.Address + AddressExt)
		DA := len(v.Address)
		v.Pdu += toCapacity(DA, 10, 16, 2)
		// Phone type
		AddressType0 := "1"
		// Number type
		// 000 – unknown;
		// 001 – international;
		// 010 – national;
		// 011 – local;
		// 100 – local user;
		// 101 – alphabet-digits;
		// 110 – short;
		// 111 – reserved.
		AddressType1 := "001"
		// Type
		// 0000 – unknown;
		// 0001 – ISDN;
		// 0010 – X.121;
		// 0011 – teletype;
		// 1000 – nation;
		// 1001 – private;
		// 1010 – ERMES;
		// 1111 – reserved.
		AddressType2 := "0001"
		// --> Phone number
		v.AddressType = toCapacity(AddressType0+AddressType1+AddressType2, 2, 16, 2)
		v.Pdu += v.AddressType + v.Address
		// --> Protocol Identifier (PID)
		v.ProtocolIdentifier = toStr(0)
		v.Pdu += toCapacity(v.ProtocolIdentifier, 10, 16, 2)
		// --> Data Coding Scheme (DCS)
		// 000>0<0000 Flash
		// 0 - Custom
		// 1 - Flash (Class 0, if 000000>00< is null)
		ViewType := "0"
		// 0000>00<00 Alphabet
		// 00 Custom Alphabet (7-bit);
		// 01 8 bit;
		// 10 UCS2 (16 bit) – Unicode;
		// 11 Reserved.
		//Alphabet := pdu.Alphabet
		v.Pdu += toCapacity("000"+ViewType+Alphabet+"00", 2, 16, 2)
		// --> Validity Period (VP)
		// For 10 length VP = 1 byte
		// 168 = A8 = 2 day
		// 173 = AD = 7 day
		// 195 = C3 = 29 day
		// 205 = CD = 91 day
		// 255 = FF = 443 day
		// default: want maximum 443 days
		v.ValidityPeriodFormat = toStr(255)
		v.Pdu += toCapacity(v.ValidityPeriodFormat, 10, 16, 2)
		if v.Parts > 1 {
			// length uniq number couple of message
			// 00 = 8-bit (1 octet, 255 values),
			// 08 = 16-bit (2 octet, 65535 values).
			MessageRefType := "00"
			if DataCodingSchemeIs7 {
				MessageRefType = "08"
				v.MessageReference = Ref2
			}
			MessagePartCount := toCapacity(v.Parts, 10, 10, 2)
			MessagePartCurrent := toCapacity(v.Part, 10, 10, 2)
			LenMessageRefPart := toCapacity(len(v.MessageReference+MessagePartCount+MessagePartCurrent)/2, 10, 10, 2)
			LenMessageHeader := toCapacity(len(MessageRefType+LenMessageRefPart+v.MessageReference+MessagePartCount+MessagePartCurrent)/2, 10, 10, 2)
			v.UserDataHeader = LenMessageHeader + MessageRefType + LenMessageRefPart + v.MessageReference + MessagePartCount + MessagePartCurrent
		}
		if DataCodingSchemeIs7 {
			tmp0 := ""
			for _, j := range v.Text {
				tmp0 += toCapacity(int(j), 10, 2, 7)
			}
			tmp1 := splitArray(trLine(joinArray(trMap(splitArray(tmp0, 7, true)))), 8, true)
			for i := 0; i < len(tmp1); i++ {
				v.UserData += strings.ToUpper(toCapacity(trLine(tmp1[i]), 2, 16, 2))
			}
			// --> User Data Length (UDL)
			v.UserDataLength = toCapacity(int(float32(len(v.UserData)*4/7))+(len(v.UserDataHeader)/2), 10, 16, 2)
			v.Pdu += v.UserDataLength + v.UserDataHeader + v.UserData
		} else {
			for _, j := range v.Text {
				v.UserData += toCapacity(int(j), 10, 16, 4)
			}
			// --> User Data Length (UDL)
			v.UserDataLength = toCapacity((len(v.UserData)/2)+(len(v.UserDataHeader)/2), 10, 16, 2)
			v.Pdu += v.UserDataLength + v.UserDataHeader + v.UserData
		}
		v.Pdu = strings.ToUpper(v.Pdu)
		sms[v.MessageReference+v.MessageType+v.Address+toStr(v.Parts)+toStr(v.Part)] = v
	}
	return sms
}

// PrintDebug()
// debug
func (obj *SMS) PrintDebug() {
	for _, o := range *obj {
		p("ServiceCenterAddressLength:", o.ServiceCenterAddressLength)
		p("ServiceCenterAddressType:", o.ServiceCenterAddressType)
		p("ServiceCenterAddress:", o.ServiceCenterAddress)
		p("MessageType:", o.MessageType)
		p("ReplyPath:", o.ReplyPath)
		p("UserDataHeaderIncluded:", o.UserDataHeaderIncluded)
		p("StatusReportRequest:", o.StatusReportRequest)
		p("ValidityPeriodFormat:", o.ValidityPeriodFormat)
		p("RejectDuplicates:", o.RejectDuplicates)
		p("MessageTypeIndicator:", o.MessageTypeIndicator)
		p("AddressLength:", o.AddressLength)
		p("AddressType:", o.AddressType)
		p("Address:", o.Address)
		p("Protocol Identifier:", o.ProtocolIdentifier)
		p("Data Coding Scheme:", o.DataCodingScheme)
		p("Date:", o.Date)
		p("ValidityPeriod:", o.ValidityPeriodSecond)
		p("User Data Header Length:", o.UserDataHeaderLength)
		p("User Data Header:", o.UserDataHeader)
		p("Message Reference:", o.MessageReference)
		p("Parts:", o.Parts)
		p("Part", o.Part)
		p("User Data Length:", o.UserDataLength)
		p("User Data:", o.UserData)
		p("PDU:", o.Pdu)
		p("Text:", o.Text)
	}
}
