package p2p

import "encoding/json"

func (nt *NetworkV1) DebugTopicMsg(topic, msg string) string {
	if err := nt.msgManager.SendMsg(topic, []byte(msg)); err != nil {
		return err.Error()
	}
	return "publish success!"
}

func (nt *NetworkV1) DebugTopicPeers(topic string) string {
	peers := nt.msgManager.PeersOfTopic(topic)
	bts, _ := json.Marshal(peers)
	return string(bts)
}
