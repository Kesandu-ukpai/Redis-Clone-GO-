package main

func EncodeSimple(s string) string{
	return "+" + s + "\r\n"
}

//error
func EncodeError(s string) string{
	return "-" + s + "\r\n"
}

//bulk message
func EncodeBulk(s string) string{
	return  "$" + s + "\r\n"
}

//integer
func EncodeInt(s integer) string{
	strconv.Itoa(s)
	return  ":" + s + "\r\n"
} 