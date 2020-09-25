package sfu

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVP8Helper_Unmarshal(t *testing.T) {
	type args struct {
		payload []byte
	}
	tests := []struct {
		name            string
		args            args
		wantErr         bool
		checkTemporal   bool
		temporalSupport bool
		checkKeyFrame   bool
		keyFrame        bool
		checkPictureID  bool
		pictureID       uint16
		checkTlzIdx     bool
		tlzIdx          uint8
		checkTempID     bool
		temporalID      uint8
	}{
		{
			name:    "Empty or nil payload must return error",
			args:    args{payload: []byte{}},
			wantErr: true,
		},
		{
			name:    "Small payloads must return errors",
			args:    args{payload: []byte{0x0, 0x1, 0x2}},
			wantErr: true,
		},
		{
			name:            "Temporal must be supported by setting T bit to 1",
			args:            args{payload: []byte{0xff, 0x20, 0x1, 0x2, 0x3, 0x4}},
			checkTemporal:   true,
			temporalSupport: true,
		},
		{
			name:           "Picture must be ID 7 bits by setting M bit to 0 and present by I bit set to 1",
			args:           args{payload: []byte{0xff, 0xff, 0x11, 0x2, 0x3, 0x4}},
			checkPictureID: true,
			pictureID:      17,
		},
		{
			name:           "Picture ID must be 15 bits by setting M bit to 1 and present by I bit set to 1",
			args:           args{payload: []byte{0xff, 0xff, 0xff, 0xfd, 0x3, 0x4, 0x5}},
			checkPictureID: true,
			pictureID:      32765,
		},
		{
			name:        "Temporal level zero index must be present if L set to 1",
			args:        args{payload: []byte{0xff, 0xff, 0xff, 0xfd, 0xb4, 0x4, 0x5}},
			checkTlzIdx: true,
			tlzIdx:      180,
		},
		{
			name:        "Temporal index must be present and used if T bit set to 1",
			args:        args{payload: []byte{0xff, 0xff, 0xff, 0xfd, 0xb4, 0x9f, 0x5, 0x6}},
			checkTempID: true,
			temporalID:  2,
		},
		{
			name:          "Check if packet is a keyframe by looking at P bit set to 0",
			args:          args{payload: []byte{0xff, 0xff, 0xff, 0xfd, 0xb4, 0x9f, 0x94, 0x1}},
			checkKeyFrame: true,
			keyFrame:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &VP8Helper{}
			if err := p.Unmarshal(tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.checkTemporal {
				assert.Equal(t, tt.temporalSupport, p.TemporalSupported)
			}
			if tt.checkKeyFrame {
				assert.Equal(t, tt.keyFrame, p.IsKeyFrame)
			}
			if tt.checkPictureID {
				assert.Equal(t, tt.pictureID, p.PictureID)
			}
			if tt.checkTlzIdx {
				assert.Equal(t, tt.tlzIdx, p.TL0PICIDX)
			}
			if tt.checkTempID {
				assert.Equal(t, tt.temporalID, p.TID)
			}
		})
	}
}

func Test_setVP8TemporalLayer(t *testing.T) {
	type args struct {
		pl []byte
		s  *WebRTCSimulcastSender
	}
	tests := []struct {
		name        string
		args        args
		wantPayload []byte
		wantSkip    bool
	}{
		{
			name: "Must skip when current temporal is bigger than wanted",
			args: args{
				s: &WebRTCSimulcastSender{
					currentTempLayer: 2,
					refPicId:         0,
					lastPicId:        0,
					refTlzi:          0,
					lastTlzi:         0,
				},
				pl: []byte{0xff, 0xff, 0xff, 0xfd, 0xb4, 0xdf, 0x5, 0x6},
			},
			wantPayload: nil,
			wantSkip:    true,
		},
		{
			name: "Must return modified payload",
			args: args{
				s: &WebRTCSimulcastSender{
					currentTempLayer: 3,
					refPicId:         32764,
					refTlzi:          179,
				},
				pl: []byte{0xff, 0xff, 0xff, 0xfd, 0xb4, 0xdf, 0x5, 0x6},
			},
			wantPayload: []byte{0xff, 0xff, 0x80, 0x01, 0x1, 0xdf, 0x5, 0x6},
			wantSkip:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPayload, gotSkip := setVP8TemporalLayer(tt.args.pl, tt.args.s)
			if !reflect.DeepEqual(gotPayload, tt.wantPayload) {
				t.Errorf("setVP8TemporalLayer() gotPayload = %v, want %v", gotPayload, tt.wantPayload)
			}
			if gotSkip != tt.wantSkip {
				t.Errorf("setVP8TemporalLayer() gotSkip = %v, want %v", gotSkip, tt.wantSkip)
			}
		})
	}
}
