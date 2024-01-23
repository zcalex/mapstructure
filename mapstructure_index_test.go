package mapstructure

import "testing"

type IndexStructRaw struct {
	IntIndex    []int
	UintIndex   []uint
	FloatIndex  []float64
	StringIndex []string
}

var raw = IndexStructRaw{
	IntIndex:    []int{-1, 0, 1, 2},
	UintIndex:   []uint{0, 1, 2, 3},
	FloatIndex:  []float64{0.1, 0.2},
	StringIndex: []string{"a", "b"},
}

type IndexStructTo struct {
	IntIndex0Int   int   `mapstructure:"intindex,index=0"`
	IntIndex0Int8  int8  `mapstructure:"intindex,index=0"`
	IntIndex0Int16 int16 `mapstructure:"intindex,index=0"`
	IntIndex0Int32 int32 `mapstructure:"intindex,index=0"`
	IntIndex0Int64 int64 `mapstructure:"intindex,index=0"`
	IntIndex1      int32 `mapstructure:"intindex,index=1"`
	IntIndex2      int32 `mapstructure:"intindex,index=2"`
	IntIndex3      int32 `mapstructure:"intindex,index=3"`

	UintIndex0Uint   uint   `mapstructure:"uintindex,index=0"`
	UintIndex0Uint8  uint8  `mapstructure:"uintindex,index=0"`
	UintIndex0Uint16 uint16 `mapstructure:"uintindex,index=0"`
	UintIndex0Uint32 uint32 `mapstructure:"uintindex,index=0"`
	UintIndex0Uint64 uint64 `mapstructure:"uintindex,index=0"`
	UintIndex1       uint32 `mapstructure:"uintindex,index=1"`
	UintIndex2       uint32 `mapstructure:"uintindex,index=2"`
	UintIndex3       uint32 `mapstructure:"uintindex,index=3"`

	FloatIndex0Float32 float32 `mapstructure:"floatindex,index=0"`
	FloatIndex0Float64 float64 `mapstructure:"floatindex,index=0"`
	FloatIndex1        float64 `mapstructure:"floatindex,index=1"`

	StringIndex0 string `mapstructure:"stringindex,index=0"`
	StringIndex1 string `mapstructure:"stringindex,index=1"`
}

func TestStructureIndex(t *testing.T) {
	out := &IndexStructTo{}
	decoder, err := NewDecoder(&DecoderConfig{
		Metadata: nil,
		Result:   out,
	})

	err = decoder.Decode(raw)
	if err != nil {
		t.Errorf("NewDecoder error: %v", err)
	}

	if out.IntIndex0Int != -1 {
		t.Errorf("IntIndex0Int error: %v", out.IntIndex0Int)
	}

	if out.IntIndex0Int8 != -1 {
		t.Errorf("IntIndex0Int8 error: %v", out.IntIndex0Int8)
	}

	if out.IntIndex0Int16 != -1 {
		t.Errorf("IntIndex0Int16 error: %v", out.IntIndex0Int16)
	}

	if out.IntIndex0Int32 != -1 {
		t.Errorf("IntIndex0Int32 error: %v", out.IntIndex0Int32)
	}

	if out.IntIndex0Int64 != -1 {
		t.Errorf("IntIndex0Int64 error: %v", out.IntIndex0Int64)
	}

	if out.IntIndex1 != 0 {
		t.Errorf("IntIndex1 error: %v", out.IntIndex1)
	}

	if out.IntIndex2 != 1 {
		t.Errorf("IntIndex2 error: %v", out.IntIndex2)
	}

	if out.IntIndex3 != 2 {
		t.Errorf("IntIndex3 error: %v", out.IntIndex3)
	}

	if out.UintIndex0Uint != 0 {
		t.Errorf("UintIndex0Uint error: %v", out.UintIndex0Uint)
	}

	if out.UintIndex0Uint8 != 0 {
		t.Errorf("UintIndex0Uint8 error: %v", out.UintIndex0Uint8)
	}

	if out.UintIndex0Uint16 != 0 {
		t.Errorf("UintIndex0Uint16 error: %v", out.UintIndex0Uint16)
	}

	if out.UintIndex0Uint32 != 0 {
		t.Errorf("UintIndex0Uint32 error: %v", out.UintIndex0Uint32)
	}

	if out.UintIndex0Uint64 != 0 {
		t.Errorf("UintIndex0Uint64 error: %v", out.UintIndex0Uint64)
	}

	if out.UintIndex1 != 1 {
		t.Errorf("UintIndex1 error: %v", out.UintIndex1)
	}

	if out.UintIndex2 != 2 {
		t.Errorf("UintIndex2 error: %v", out.UintIndex2)
	}

	if out.UintIndex3 != 3 {
		t.Errorf("UintIndex3 error: %v", out.UintIndex3)
	}

	if out.FloatIndex0Float32 < 0.1-0.000001 || out.FloatIndex0Float32 > 0.1+0.000001 {
		t.Errorf("FloatIndex0Float32 error: %v", out.FloatIndex0Float32)
	}

	if out.FloatIndex0Float64 < 0.1-0.000001 || out.FloatIndex0Float64 > 0.1+0.000001 {
		t.Errorf("FloatIndex0Float64 error: %v", out.FloatIndex0Float64)
	}

	if out.FloatIndex1 < 0.2-0.000001 || out.FloatIndex1 > 0.2+0.000001 {
		t.Errorf("FloatIndex1 error: %v", out.FloatIndex1)
	}

	if out.StringIndex0 != "a" {
		t.Errorf("StringIndex0 error: %v", out.StringIndex0)
	}

	if out.StringIndex1 != "b" {
		t.Errorf("StringIndex1 error: %v", out.StringIndex1)
	}
}

func TestStructureIndexTooLarge(t *testing.T) {
	out := struct {
		Value int32 `mapstructure:"intindex,index=100"`
	}{}
	decoder, err := NewDecoder(&DecoderConfig{
		Metadata: nil,
		Result:   &out,
	})

	if err != nil {
		t.Errorf("NewDecoder error: %v", err)
		return
	}

	err = decoder.Decode(raw)
	if err == nil {
		t.Error("Decode error success, but need err")
	}
}

func TestStructureIndexDiffValue(t *testing.T) {
	out := struct {
		Value int32 `mapstructure:"uintindex,index=0"`
	}{}
	decoder, err := NewDecoder(&DecoderConfig{
		Metadata: nil,
		Result:   &out,
	})

	if err != nil {
		t.Errorf("NewDecoder error: %v", err)
		return
	}

	err = decoder.Decode(raw)
	if err == nil {
		t.Error("Decode error success, but need err")
	}
}
