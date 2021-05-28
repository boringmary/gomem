// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gomem/mservices/card/card.proto

package card

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Empty) Validate() error {
	return nil
}
func (this *Card) Validate() error {
	return nil
}
func (this *CardsResponse) Validate() error {
	for _, item := range this.Cards {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Cards", err)
			}
		}
	}
	return nil
}
