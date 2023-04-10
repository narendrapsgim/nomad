// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package structs

import "errors"

// An SIToken is the important bits of a Service Identity token generated by Consul.
type SIToken struct {
	ConsulNamespace string
	TaskName        string // the nomad task backing the consul service (native or sidecar)
	AccessorID      string
	SecretID        string
}

// An SITokenAccessor is a reference to a created Consul Service Identity token on
// behalf of an allocation's task.
type SITokenAccessor struct {
	ConsulNamespace string
	NodeID          string
	AllocID         string
	AccessorID      string
	TaskName        string

	// Raft index
	CreateIndex uint64
}

// SITokenAccessorsRequest is used to operate on a set of SITokenAccessor, like
// recording a set of accessors for an alloc into raft.
type SITokenAccessorsRequest struct {
	Accessors []*SITokenAccessor
}

// DeriveSITokenRequest is used to request Consul Service Identity tokens from
// the Nomad Server for the named tasks in the given allocation.
type DeriveSITokenRequest struct {
	NodeID   string
	SecretID string
	AllocID  string
	Tasks    []string
	QueryOptions
}

func (r *DeriveSITokenRequest) Validate() error {
	switch {
	case r.NodeID == "":
		return errors.New("missing node ID")
	case r.SecretID == "":
		return errors.New("missing node SecretID")
	case r.AllocID == "":
		return errors.New("missing allocation ID")
	case len(r.Tasks) == 0:
		return errors.New("no tasks specified")
	default:
		return nil
	}
}

type DeriveSITokenResponse struct {
	// Tokens maps from Task Name to its associated SI token
	Tokens map[string]string

	// Error stores any error that occurred. Errors are stored here so we can
	// communicate whether it is retryable
	Error *RecoverableError

	QueryMeta
}
