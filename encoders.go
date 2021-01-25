package mongodb_zero_encoders

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"reflect"
)

//Register
//NewRegistry
var dve = bsoncodec.DefaultValueEncoders{}
var sc = bsoncodec.NewStringCodec()
var uic = bsoncodec.NewUIntCodec()

const (
	zeroBool    bool    = false
	zeroString  string  = ""
	zeroInt     int     = 0
	zeroInt8    int8    = 0
	zeroInt16   int16   = 0
	zeroInt32   int32   = 0
	zeroInt64   int64   = 0
	zeroUint    uint    = 0
	zeroUint8   uint8   = 0
	zeroUint16  uint16  = 0
	zeroUint32  uint32  = 0
	zeroUint64  uint64  = 0
	zeroFloat32 float32 = 0
	zeroFloat64 float64 = 0
)

func RegisterZeroEncoders(rb *bsoncodec.RegistryBuilder) {
	rb.
		RegisterTypeEncoder(reflect.TypeOf(zeroBool), &BoolZeroEncoder{}).
		RegisterTypeEncoder(reflect.TypeOf(zeroString), &StringZeroEncoder{}).
		RegisterTypeEncoder(reflect.TypeOf(zeroInt), &IntZeroEncoder{}).
		RegisterTypeEncoder(reflect.TypeOf(zeroInt8), &IntZeroEncoder{}).
		RegisterTypeEncoder(reflect.TypeOf(zeroInt16), &IntZeroEncoder{}).
		RegisterTypeEncoder(reflect.TypeOf(zeroInt32), &IntZeroEncoder{}).
		RegisterTypeEncoder(reflect.TypeOf(zeroInt64), &IntZeroEncoder{}).
		RegisterTypeEncoder(reflect.TypeOf(zeroUint), &UintZeroEncoder{}).
		RegisterTypeEncoder(reflect.TypeOf(zeroUint8), &UintZeroEncoder{}).
		RegisterTypeEncoder(reflect.TypeOf(zeroUint16), &UintZeroEncoder{}).
		RegisterTypeEncoder(reflect.TypeOf(zeroUint32), &UintZeroEncoder{}).
		RegisterTypeEncoder(reflect.TypeOf(zeroUint64), &UintZeroEncoder{}).
		RegisterTypeEncoder(reflect.TypeOf(zeroFloat32), &Float32ZeroEncoder{}).
		RegisterTypeEncoder(reflect.TypeOf(zeroFloat64), &Float64ZeroEncoder{})

}

func DefaultZeroEncoders() *bsoncodec.Registry {
	rb := bson.NewRegistryBuilder()
	RegisterZeroEncoders(rb)
	return rb.Build()
}

type BoolZeroEncoder struct{}

func (b *BoolZeroEncoder) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	return dve.BooleanEncodeValue(ec, vw, val)
}

func (b *BoolZeroEncoder) IsTypeZero(i interface{}) bool {
	if b, ok := i.(bool); ok {
		return b == zeroBool
	}

	return true
}

type StringZeroEncoder struct{}

func (s *StringZeroEncoder) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	return sc.EncodeValue(ec, vw, val)
}

func (s *StringZeroEncoder) IsTypeZero(i interface{}) bool {
	if s, ok := i.(string); ok {
		return s == zeroString
	}

	return false
}

type IntZeroEncoder struct{}

func (_ *IntZeroEncoder) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	return dve.IntEncodeValue(ec, vw, val)
}

func (_ *IntZeroEncoder) IsTypeZero(i interface{}) bool {
	switch v := i.(type) {
	case int8:
		return v == zeroInt8
	case int16:
		return v == zeroInt16
	case int32:
		return v == zeroInt32
	case int:
		return v == zeroInt
	case int64:
		return v == zeroInt64
	}

	return false
}

type UintZeroEncoder struct{}

func (s *UintZeroEncoder) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	return uic.EncodeValue(ec, vw, val)
}

func (s *UintZeroEncoder) IsTypeZero(i interface{}) bool {
	switch v := i.(type) {
	case uint8:
		return v == zeroUint8
	case uint16:
		return v == zeroUint16
	case uint32:
		return v == zeroUint32
	case uint:
		return v == zeroUint
	case uint64:
		return v == zeroUint64
	}

	return false
}

type Float32ZeroEncoder struct{}

func (_ *Float32ZeroEncoder) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	return dve.FloatEncodeValue(ec, vw, val)
}

func (_ *Float32ZeroEncoder) IsTypeZero(i interface{}) bool {
	if v, ok := i.(float32); ok {
		return v == zeroFloat32
	}

	return false
}

type Float64ZeroEncoder struct{}

func (s *Float64ZeroEncoder) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	return dve.FloatEncodeValue(ec, vw, val)
}

func (_ *Float64ZeroEncoder) IsTypeZero(i interface{}) bool {
	if v, ok := i.(float64); ok {
		return v == zeroFloat64
	}

	return false
}
