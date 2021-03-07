package p2p

import (
	"encoding/json"
	"github.com/otcChain/chord-go/rpc"
	"github.com/otcChain/chord-go/utils"
)

type RpcPushTopic struct {
	Topics  string `json:"topic"`
	Message string `json:"msg"`
}

func (nt *NetworkV1) initRpcApis() {
	rpc.HttpRpcApis["/p2p/PeerList"] = nt.ApiPeesList
	rpc.HttpRpcApis["/p2p/PushMsg"] = nt.ApiPushMsg
	rpc.HttpRpcApis["/p2p/nid"] = nt.HostID
}

//--->public rpc apis
func (nt *NetworkV1) ApiPeesList(msg *rpc.JsonRpcMessageItem) (json.RawMessage, *rpc.JsonError) {
	peerStr := nt.DebugTopicPeers(string(msg.Params))
	return []byte(peerStr), nil
}

func (nt *NetworkV1) HostID(msg *rpc.JsonRpcMessageItem) (json.RawMessage, *rpc.JsonError) {
	return []byte(nt.p2pHost.ID()), nil
}

func (nt *NetworkV1) ApiPushMsg(msg *rpc.JsonRpcMessageItem) (json.RawMessage, *rpc.JsonError) {
	param := &RpcPushTopic{}
	if err := json.Unmarshal(msg.Params, param); err != nil {
		return nil, &rpc.JsonError{
			Code:    -1,
			Message: err.Error(),
		}
	}
	res := nt.DebugTopicMsg(param.Topics, param.Message)
	return []byte(res), nil
}

//---rpc debug
func (nt *NetworkV1) DebugTopicMsg(topic, msg string) string {
	if err := nt.msgManager.SendMsg(topic, []byte(msg)); err != nil {
		return err.Error()
	}
	return "publish success!"
}

func (nt *NetworkV1) DebugTopicPeers(topic string) string {
	utils.LogInst().Debug().Msgf("p2p cmd rpc query for topic[%s]", topic)
	peers := nt.msgManager.PeersOfTopic(topic)
	bts, _ := json.Marshal(peers)
	return string(bts)
}
