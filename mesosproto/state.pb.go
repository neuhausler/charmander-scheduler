// Code generated by protoc-gen-go.
// source: state.proto
// DO NOT EDIT!

package mesosproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Operation_Type int32

const (
	Operation_SNAPSHOT Operation_Type = 1
	Operation_DIFF     Operation_Type = 3
	Operation_EXPUNGE  Operation_Type = 2
)

var Operation_Type_name = map[int32]string{
	1: "SNAPSHOT",
	3: "DIFF",
	2: "EXPUNGE",
}
var Operation_Type_value = map[string]int32{
	"SNAPSHOT": 1,
	"DIFF":     3,
	"EXPUNGE":  2,
}

func (x Operation_Type) Enum() *Operation_Type {
	p := new(Operation_Type)
	*p = x
	return p
}
func (x Operation_Type) String() string {
	return proto.EnumName(Operation_Type_name, int32(x))
}
func (x *Operation_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Operation_Type_value, data, "Operation_Type")
	if err != nil {
		return err
	}
	*x = Operation_Type(value)
	return nil
}

// Describes a state entry, a versioned (via a UUID) key/value pair.
type Entry struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Uuid             []byte  `protobuf:"bytes,2,req,name=uuid" json:"uuid,omitempty"`
	Value            []byte  `protobuf:"bytes,3,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}

func (m *Entry) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Entry) GetUuid() []byte {
	if m != nil {
		return m.Uuid
	}
	return nil
}

func (m *Entry) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// Describes an operation used in the log storage implementation.
type Operation struct {
	Type             *Operation_Type     `protobuf:"varint,1,req,name=type,enum=mesosproto.Operation_Type" json:"type,omitempty"`
	Snapshot         *Operation_Snapshot `protobuf:"bytes,2,opt,name=snapshot" json:"snapshot,omitempty"`
	Diff             *Operation_Diff     `protobuf:"bytes,4,opt,name=diff" json:"diff,omitempty"`
	Expunge          *Operation_Expunge  `protobuf:"bytes,3,opt,name=expunge" json:"expunge,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *Operation) Reset()         { *m = Operation{} }
func (m *Operation) String() string { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()    {}

func (m *Operation) GetType() Operation_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Operation_SNAPSHOT
}

func (m *Operation) GetSnapshot() *Operation_Snapshot {
	if m != nil {
		return m.Snapshot
	}
	return nil
}

func (m *Operation) GetDiff() *Operation_Diff {
	if m != nil {
		return m.Diff
	}
	return nil
}

func (m *Operation) GetExpunge() *Operation_Expunge {
	if m != nil {
		return m.Expunge
	}
	return nil
}

// Describes a "snapshot" operation.
type Operation_Snapshot struct {
	Entry            *Entry `protobuf:"bytes,1,req,name=entry" json:"entry,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Operation_Snapshot) Reset()         { *m = Operation_Snapshot{} }
func (m *Operation_Snapshot) String() string { return proto.CompactTextString(m) }
func (*Operation_Snapshot) ProtoMessage()    {}

func (m *Operation_Snapshot) GetEntry() *Entry {
	if m != nil {
		return m.Entry
	}
	return nil
}

// Describes a "diff" operation where the 'value' of the entry is
// just the diff itself, but the 'uuid' represents the UUID of the
// entry after applying this diff.
type Operation_Diff struct {
	Entry            *Entry `protobuf:"bytes,1,req,name=entry" json:"entry,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Operation_Diff) Reset()         { *m = Operation_Diff{} }
func (m *Operation_Diff) String() string { return proto.CompactTextString(m) }
func (*Operation_Diff) ProtoMessage()    {}

func (m *Operation_Diff) GetEntry() *Entry {
	if m != nil {
		return m.Entry
	}
	return nil
}

// Describes an "expunge" operation.
type Operation_Expunge struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Operation_Expunge) Reset()         { *m = Operation_Expunge{} }
func (m *Operation_Expunge) String() string { return proto.CompactTextString(m) }
func (*Operation_Expunge) ProtoMessage()    {}

func (m *Operation_Expunge) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func init() {
	proto.RegisterEnum("mesosproto.Operation_Type", Operation_Type_name, Operation_Type_value)
}
