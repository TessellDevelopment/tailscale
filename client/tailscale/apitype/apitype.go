// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

// Package apitype contains types for the Tailscale LocalAPI and control plane API.
package apitype

import "tailscale.com/tailcfg"

// LocalAPIHost is the Host header value used by the LocalAPI.
const LocalAPIHost = "local-tailscaled.sock"

// WhoIsResponse is the JSON type returned by tailscaled debug server's /whois?ip=$IP handler.
type WhoIsResponse struct {
	Node        *tailcfg.Node
	UserProfile *tailcfg.UserProfile

	// CapMap is a map of capabilities to their values.
	// See tailcfg.PeerCapMap and tailcfg.PeerCapability for details.
	CapMap tailcfg.PeerCapMap
}

// FileTarget is a node to which files can be sent, and the PeerAPI
// URL base to do so via.
type FileTarget struct {
	Node *tailcfg.Node

	// PeerAPI is the http://ip:port URL base of the node's PeerAPI,
	// without any path (not even a single slash).
	PeerAPIURL string
}

type WaitingFile struct {
	Name string
	Size int64
}

// SetPushDeviceTokenRequest is the body POSTed to the LocalAPI endpoint /set-device-token.
type SetPushDeviceTokenRequest struct {
	// PushDeviceToken is the iOS/macOS APNs device token (and any future Android equivalent).
	PushDeviceToken string
}
