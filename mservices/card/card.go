package card

type CardEntity struct {
	Question string `protobuf:"bytes,1,opt,name=Question,proto3" json:"Question,omitempty"`
	Answer string `protobuf:"bytes,2,opt,name=Answer,proto3" json:"Answer,omitempty"`
	Interval int32 `protobuf:"float,3,opt,name=Interval,proto3" json:"Interval,omitempty"`
	Repetitions int32 `protobuf:"int,4,opt,name=Repetitions,proto3" json:"Repetitions,omitempty"`
	eFactor float64 `protobuf:"float,5,opt,name=eFactor,proto3" json:"eFactor,omitempty"`
}
