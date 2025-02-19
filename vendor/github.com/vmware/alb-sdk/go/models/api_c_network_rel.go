// Copyright 2021 VMware, Inc.
// SPDX-License-Identifier: Apache License 2.0
package models

// This file is auto-generated.

// APICNetworkRel API c network rel
// swagger:model APICNetworkRel
type APICNetworkRel struct {

	// connector of APICNetworkRel.
	Connector *string `json:"connector,omitempty"`

	// rel_key of APICNetworkRel.
	RelKey *string `json:"rel_key,omitempty"`

	// target_network of APICNetworkRel.
	TargetNetwork *string `json:"target_network,omitempty"`
}
