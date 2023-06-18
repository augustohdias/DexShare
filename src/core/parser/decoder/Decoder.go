package decoder

type Decoder interface {
	Decode([]byte) interface{}
}
