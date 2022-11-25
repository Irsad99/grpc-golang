// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: product.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *RequestProduct) Validate() error {
	if this.Category != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Category); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Category", err)
		}
	}
	return nil
}
func (this *ResponseProduct) Validate() error {
	if this.ResponseData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ResponseData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ResponseData", err)
		}
	}
	return nil
}
func (this *Product) Validate() error {
	if this.Category != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Category); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Category", err)
		}
	}
	return nil
}
func (this *Category) Validate() error {
	return nil
}
func (this *Id) Validate() error {
	return nil
}