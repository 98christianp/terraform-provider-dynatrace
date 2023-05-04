/**
* @license
* Copyright 2020 Dynatrace LLC
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package hcl

import (
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Marshaler has no documentation
type Marshaler interface {
	MarshalHCL(Properties) error
}

type Unmarshaler interface {
	UnmarshalHCL(decoder Decoder) error
}

type Schemer interface {
	Schema() map[string]*schema.Schema
}

type Preconditioner interface {
	HandlePreconditions() error
}

func UnmarshalHCL(m Unmarshaler, d Decoder) error {
	if err := m.UnmarshalHCL(d); err != nil {
		return err
	}
	if pc, ok := m.(Preconditioner); ok {
		if err := pc.HandlePreconditions(); err != nil {
			return errors.New(d.Path() + ": " + err.Error())
		}
	}
	return nil
}
