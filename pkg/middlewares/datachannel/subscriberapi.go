package datachannel

import (
	"encoding/json"

	"github.com/pion/ion-sfu/pkg/sfu"
	"github.com/pion/webrtc/v3"
)

const (
	videoHighQuality   = "high"
	videoMediumQuality = "medium"
	videoLowQuality    = "low"
	videoMuted         = "none"
)

type setRemoteMedia struct {
	StreamID string `json:"streamId"`
	Video    string `json:"video"`
	Audio    bool   `json:"audio"`
}

func SubscriberAPI(next sfu.MessageProcessor) sfu.MessageProcessor {
	return sfu.ProcessFunc(func(peer *sfu.Peer, dc *webrtc.DataChannel, msg webrtc.DataChannelMessage) {
		srm := &setRemoteMedia{}
		if err := json.Unmarshal(msg.Data, srm); err != nil {
			return
		}
		downTracks := peer.Subscriber.GetDownTracks(srm.StreamID)

		for _, dt := range downTracks {
			switch dt.Kind() {
			case webrtc.RTPCodecTypeAudio:
				dt.Mute(!srm.Audio)
			case webrtc.RTPCodecTypeVideo:
				switch srm.Video {
				case videoHighQuality:
					dt.Mute(false)
					dt.SwitchSpatialLayer(2)
				case videoMediumQuality:
					dt.Mute(false)
					dt.SwitchSpatialLayer(1)
				case videoLowQuality:
					dt.Mute(false)
					dt.SwitchSpatialLayer(0)
				case videoMuted:
					dt.Mute(true)
				}
			}
		}
		next.Process(peer, dc, msg)
	})
}
