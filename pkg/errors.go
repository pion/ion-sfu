package sfu

import "errors"

var (
	errPeerConnectionInitFailed = errors.New("pc init failed")
	errPtNotSupported           = errors.New("payload type not supported")
	// Helpers errors
	errShortPacket = errors.New("packet is not large enough")
	errNilPacket   = errors.New("invalid nil packet")
)
