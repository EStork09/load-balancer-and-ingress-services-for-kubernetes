// Copyright 2021 VMware, Inc.
// SPDX-License-Identifier: Apache License 2.0
package models

// This file is auto-generated.

// VirtualserviceFaults virtualservice faults
// swagger:model VirtualserviceFaults
type VirtualserviceFaults struct {

	// Enable debug faults. Field introduced in 20.1.6.
	DebugFaults *bool `json:"debug_faults,omitempty"`

	// Enable pool server faults. Field introduced in 20.1.6.
	PoolServerFaults *bool `json:"pool_server_faults,omitempty"`

	// Enable VS scaleout and scalein faults. Field introduced in 20.1.6.
	ScaleoutFaults *bool `json:"scaleout_faults,omitempty"`

	// Enable shared vip faults. Field introduced in 20.1.6.
	SharedVipFaults *bool `json:"shared_vip_faults,omitempty"`

	// Enable SSL certificate expiry faults. Field introduced in 20.1.6.
	SslCertExpiryFaults *bool `json:"ssl_cert_expiry_faults,omitempty"`

	// Enable SSL certificate status faults. Field introduced in 20.1.6.
	SslCertStatusFaults *bool `json:"ssl_cert_status_faults,omitempty"`
}
