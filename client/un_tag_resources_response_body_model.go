// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnTagResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnTagResourcesResponseBody
	GetRequestId() *string
}

type UnTagResourcesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C46FF5A8-C5F0-4024-8262-B16B639225A0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnTagResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UnTagResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnTagResourcesResponseBody) SetRequestId(v string) *UnTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnTagResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
