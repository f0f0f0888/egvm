package types

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *LambdaJob) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "script":
			z.Script, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Script")
				return
			}
		case "cert":
			z.Cert, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Cert")
				return
			}
		case "config":
			z.Config, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Config")
				return
			}
		case "inputs":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Inputs")
				return
			}
			if cap(z.Inputs) >= int(zb0002) {
				z.Inputs = (z.Inputs)[:zb0002]
			} else {
				z.Inputs = make([][]byte, zb0002)
			}
			for za0001 := range z.Inputs {
				z.Inputs[za0001], err = dc.ReadBytes(z.Inputs[za0001])
				if err != nil {
					err = msgp.WrapError(err, "Inputs", za0001)
					return
				}
			}
		case "state":
			z.State, err = dc.ReadBytes(z.State)
			if err != nil {
				err = msgp.WrapError(err, "State")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *LambdaJob) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 5
	// write "script"
	err = en.Append(0x85, 0xa6, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74)
	if err != nil {
		return
	}
	err = en.WriteString(z.Script)
	if err != nil {
		err = msgp.WrapError(err, "Script")
		return
	}
	// write "cert"
	err = en.Append(0xa4, 0x63, 0x65, 0x72, 0x74)
	if err != nil {
		return
	}
	err = en.WriteString(z.Cert)
	if err != nil {
		err = msgp.WrapError(err, "Cert")
		return
	}
	// write "config"
	err = en.Append(0xa6, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67)
	if err != nil {
		return
	}
	err = en.WriteString(z.Config)
	if err != nil {
		err = msgp.WrapError(err, "Config")
		return
	}
	// write "inputs"
	err = en.Append(0xa6, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Inputs)))
	if err != nil {
		err = msgp.WrapError(err, "Inputs")
		return
	}
	for za0001 := range z.Inputs {
		err = en.WriteBytes(z.Inputs[za0001])
		if err != nil {
			err = msgp.WrapError(err, "Inputs", za0001)
			return
		}
	}
	// write "state"
	err = en.Append(0xa5, 0x73, 0x74, 0x61, 0x74, 0x65)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.State)
	if err != nil {
		err = msgp.WrapError(err, "State")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *LambdaJob) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "script"
	o = append(o, 0x85, 0xa6, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74)
	o = msgp.AppendString(o, z.Script)
	// string "cert"
	o = append(o, 0xa4, 0x63, 0x65, 0x72, 0x74)
	o = msgp.AppendString(o, z.Cert)
	// string "config"
	o = append(o, 0xa6, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67)
	o = msgp.AppendString(o, z.Config)
	// string "inputs"
	o = append(o, 0xa6, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Inputs)))
	for za0001 := range z.Inputs {
		o = msgp.AppendBytes(o, z.Inputs[za0001])
	}
	// string "state"
	o = append(o, 0xa5, 0x73, 0x74, 0x61, 0x74, 0x65)
	o = msgp.AppendBytes(o, z.State)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *LambdaJob) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "script":
			z.Script, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Script")
				return
			}
		case "cert":
			z.Cert, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Cert")
				return
			}
		case "config":
			z.Config, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Config")
				return
			}
		case "inputs":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Inputs")
				return
			}
			if cap(z.Inputs) >= int(zb0002) {
				z.Inputs = (z.Inputs)[:zb0002]
			} else {
				z.Inputs = make([][]byte, zb0002)
			}
			for za0001 := range z.Inputs {
				z.Inputs[za0001], bts, err = msgp.ReadBytesBytes(bts, z.Inputs[za0001])
				if err != nil {
					err = msgp.WrapError(err, "Inputs", za0001)
					return
				}
			}
		case "state":
			z.State, bts, err = msgp.ReadBytesBytes(bts, z.State)
			if err != nil {
				err = msgp.WrapError(err, "State")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *LambdaJob) Msgsize() (s int) {
	s = 1 + 7 + msgp.StringPrefixSize + len(z.Script) + 5 + msgp.StringPrefixSize + len(z.Cert) + 7 + msgp.StringPrefixSize + len(z.Config) + 7 + msgp.ArrayHeaderSize
	for za0001 := range z.Inputs {
		s += msgp.BytesPrefixSize + len(z.Inputs[za0001])
	}
	s += 6 + msgp.BytesPrefixSize + len(z.State)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *LambdaResult) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "outputs":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Outputs")
				return
			}
			if cap(z.Outputs) >= int(zb0002) {
				z.Outputs = (z.Outputs)[:zb0002]
			} else {
				z.Outputs = make([][]byte, zb0002)
			}
			for za0001 := range z.Outputs {
				z.Outputs[za0001], err = dc.ReadBytes(z.Outputs[za0001])
				if err != nil {
					err = msgp.WrapError(err, "Outputs", za0001)
					return
				}
			}
		case "state":
			z.State, err = dc.ReadBytes(z.State)
			if err != nil {
				err = msgp.WrapError(err, "State")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *LambdaResult) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "outputs"
	err = en.Append(0x82, 0xa7, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Outputs)))
	if err != nil {
		err = msgp.WrapError(err, "Outputs")
		return
	}
	for za0001 := range z.Outputs {
		err = en.WriteBytes(z.Outputs[za0001])
		if err != nil {
			err = msgp.WrapError(err, "Outputs", za0001)
			return
		}
	}
	// write "state"
	err = en.Append(0xa5, 0x73, 0x74, 0x61, 0x74, 0x65)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.State)
	if err != nil {
		err = msgp.WrapError(err, "State")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *LambdaResult) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "outputs"
	o = append(o, 0x82, 0xa7, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Outputs)))
	for za0001 := range z.Outputs {
		o = msgp.AppendBytes(o, z.Outputs[za0001])
	}
	// string "state"
	o = append(o, 0xa5, 0x73, 0x74, 0x61, 0x74, 0x65)
	o = msgp.AppendBytes(o, z.State)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *LambdaResult) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "outputs":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Outputs")
				return
			}
			if cap(z.Outputs) >= int(zb0002) {
				z.Outputs = (z.Outputs)[:zb0002]
			} else {
				z.Outputs = make([][]byte, zb0002)
			}
			for za0001 := range z.Outputs {
				z.Outputs[za0001], bts, err = msgp.ReadBytesBytes(bts, z.Outputs[za0001])
				if err != nil {
					err = msgp.WrapError(err, "Outputs", za0001)
					return
				}
			}
		case "state":
			z.State, bts, err = msgp.ReadBytesBytes(bts, z.State)
			if err != nil {
				err = msgp.WrapError(err, "State")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *LambdaResult) Msgsize() (s int) {
	s = 1 + 8 + msgp.ArrayHeaderSize
	for za0001 := range z.Outputs {
		s += msgp.BytesPrefixSize + len(z.Outputs[za0001])
	}
	s += 6 + msgp.BytesPrefixSize + len(z.State)
	return
}
