package mongodb_zero_encoders

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestDefaultZeroEncoders_RespectOmitEmpty(t *testing.T) {
	def := DefaultZeroEncoders()

	type obj struct {
		Bool    bool    `bson:"b"`
		String  string  `bson:"s"`
		Int     int     `bson:"i"`
		Int8    int8    `bson:"i8"`
		Int16   int16   `bson:"i16"`
		Int32   int32   `bson:"i32"`
		Int64   int64   `bson:"i64"`
		Uint    uint    `bson:"ui"`
		Uint8   uint8   `bson:"ui8"`
		Uint16  uint16  `bson:"ui16"`
		Uint32  uint32  `bson:"ui32"`
		Uint64  uint64  `bson:"ui64"`
		Float32 float32 `bson:"f32"`
		Float64 float64 `bson:"f64"`
	}

	var o obj

	bytes, err := bson.MarshalWithRegistry(def, o)
	if err != nil {
		fmt.Println(err.Error())
	}
	require.Nil(t, err)

	var target map[string]interface{}

	err = bson.Unmarshal(bytes, &target)
	require.Nil(t, err)
	assert.Len(t, target, 14)
}
func TestDefaultZeroEncoders(t *testing.T) {
	def := DefaultZeroEncoders()

	type obj struct {
		Bool    bool    `bson:"b,omitempty"`
		String  string  `bson:"s,omitempty"`
		Int     int     `bson:"i,omitempty"`
		Int8    int8    `bson:"i8,omitempty"`
		Int16   int16   `bson:"i16,omitempty"`
		Int32   int32   `bson:"i32,omitempty"`
		Int64   int64   `bson:"i64,omitempty"`
		Uint    uint    `bson:"ui,omitempty"`
		Uint8   uint8   `bson:"ui8,omitempty"`
		Uint16  uint16  `bson:"ui16,omitempty"`
		Uint32  uint32  `bson:"ui32,omitempty"`
		Uint64  uint64  `bson:"ui64,omitempty"`
		Float32 float32 `bson:"f32,omitempty"`
		Float64 float64 `bson:"f64,omitempty"`
	}

	var o obj

	bytes, err := bson.MarshalWithRegistry(def, o)
	if err != nil {
		fmt.Println(err.Error())
	}
	require.Nil(t, err)

	var target map[string]interface{}

	err = bson.Unmarshal(bytes, &target)
	require.Nil(t, err)
	assert.Empty(t, target)
}
