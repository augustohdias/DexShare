package decoder

type StringDecoder struct{}

func (s StringDecoder) Decode(bytes []byte) interface{} {
	str := ""
	for _, b := range bytes {
		str += s.decodeByte(b)
	}
	return str
}

func (s StringDecoder) decodeByte(b byte) string {
	decodeMap := map[byte]string{
		187: "A", 188: "B", 189: "C", 190: "D", 191: "E", 192: "F", 193: "G", 194: "H",
		195: "I", 196: "J", 197: "K", 198: "L", 199: "M", 200: "N", 201: "O", 202: "P",
		203: "Q", 204: "R", 205: "S", 206: "T", 207: "U", 208: "V", 209: "W", 210: "X",
		211: "Y", 212: "Z",
	}
	return decodeMap[b]
}
