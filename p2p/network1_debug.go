package p2p

import "encoding/json"

func (nt *NetworkV1) DebugTopicMsg(topic, msg string) string {

	topics := nt.msgManager.topics
	nt.msgManager.lock.RLock()
	defer nt.msgManager.lock.Unlock()

	t, ok := topics[MessageChannel(topic)]
	if !ok {
		return "no such topic"
	}

	if err := t.Publish(nt.ctx, []byte(msg)); err != nil {
		return err.Error()
	}
	return "publish success!"
}

func (nt *NetworkV1) DebugTopicPeers(topic string) string {
	topics := nt.msgManager.topics
	nt.msgManager.lock.RLock()
	defer nt.msgManager.lock.Unlock()
	t, ok := topics[MessageChannel(topic)]
	if !ok {
		return "no such topic"
	}
	result := t.ListPeers()
	bts, _ := json.Marshal(result)
	return string(bts)
}
