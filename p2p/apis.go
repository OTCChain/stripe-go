package p2p

import (
	"encoding/json"
	pbs "github.com/otcChain/chord-go/pbs/rpc"
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
func (nt *NetworkV1) ApiPeesList(request *pbs.RpcMsgItem) *pbs.RpcResponse {
	peerStr := nt.DebugTopicPeers(string(request.Parameter))
	return pbs.RpcOk([]byte(peerStr))
}

func (nt *NetworkV1) HostID(_ *pbs.RpcMsgItem) *pbs.RpcResponse {
	return pbs.RpcOk([]byte(nt.p2pHost.ID()))
}

func (nt *NetworkV1) ApiPushMsg(request *pbs.RpcMsgItem) *pbs.RpcResponse {
	param := &RpcPushTopic{}
	if err := json.Unmarshal(request.Parameter, param); err != nil {
		return pbs.RpcError(err.Error())
	}
	res := nt.DebugTopicMsg(param.Topics, param.Message)
	return pbs.RpcOk([]byte(res))
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
