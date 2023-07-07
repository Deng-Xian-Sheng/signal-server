//Copyright (c) [2023] [JinCanQi]
//[make_data_set_so-vits-svc] is licensed under Mulan PubL v2.
//You can use this software according to the terms and conditions of the Mulan PubL v2.
//You may obtain a copy of Mulan PubL v2 at:
//         http://license.coscl.org.cn/MulanPubL-2.0
//THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
//EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
//MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
//See the Mulan PubL v2 for more details.

package offer

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"

	webrtc "github.com/pion/webrtc/v3"
)

func aaa() { //nolint:gocognit
	offerAddr := flag.String("offer-address", ":50000", "offer HTTP 服务器的地址。")
	answerAddr := flag.String("answer-address", "127.0.0.1:60000", "answer HTTP 服务器的地址。")
	flag.Parse()

	var candidatesMux sync.Mutex
	pendingCandidates := make([]*webrtc.ICECandidate, 0)

	// 下面的所有内容都是Pion WebRTC API！感谢您使用它❤️.

	// 准备配置
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}

	// 创建新的 RTCPeerConnection
	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if cErr := peerConnection.Close(); cErr != nil {
			fmt.Printf("无法关闭对等连接: %v\n", cErr)
		}
	}()

	// 当 ICE 候选项可用时，发送到其他 Pion 实例
	// 另一个 Pion 实例将通过调用 AddICECandidate 来添加此候选项。
	peerConnection.OnICECandidate(func(c *webrtc.ICECandidate) {
		if c == nil {
			return
		}

		candidatesMux.Lock()
		defer candidatesMux.Unlock()

		desc := peerConnection.RemoteDescription()
		if desc == nil {
			pendingCandidates = append(pendingCandidates, c)
		} else if onICECandidateErr := signalCandidate(*answerAddr, c); onICECandidateErr != nil {
			panic(onICECandidateErr)
		}
	})

	// 一个 HTTP 处理程序，允许其他 Pion 实例向我们发送 ICE 候选
	// 这使我们能够更快地添加 ICE 候选者，我们不必等待 STUN 或 TURN
	// 可能较慢的候选人
	http.HandleFunc("/candidate", func(w http.ResponseWriter, r *http.Request) {
		candidate, candidateErr := ioutil.ReadAll(r.Body)
		if candidateErr != nil {
			panic(candidateErr)
		}
		if candidateErr := peerConnection.AddICECandidate(webrtc.ICECandidateInit{Candidate: string(candidate)}); candidateErr != nil {
			panic(candidateErr)
		}
	})

	// 一个HTTP处理程序，用于处理从其他Pion进程提供给我们的会话描述
	http.HandleFunc("/sdp", func(w http.ResponseWriter, r *http.Request) {
		sdp := webrtc.SessionDescription{}
		if sdpErr := json.NewDecoder(r.Body).Decode(&sdp); sdpErr != nil {
			panic(sdpErr)
		}

		if sdpErr := peerConnection.SetRemoteDescription(sdp); sdpErr != nil {
			panic(sdpErr)
		}

		candidatesMux.Lock()
		defer candidatesMux.Unlock()

		for _, c := range pendingCandidates {
			if onICECandidateErr := signalCandidate(*answerAddr, c); onICECandidateErr != nil {
				panic(onICECandidateErr)
			}
		}
	})
	// 启动接受来自应答进程的请求的 HTTP 服务器
	// nolint: gosec
	go func() { panic(http.ListenAndServe(*offerAddr, nil)) }()

	// 创建带标签的数据通道 'data'
	dataChannel, err := peerConnection.CreateDataChannel("data", nil)
	if err != nil {
		panic(err)
	}

	// 为对等连接状态设置处理程序
	// 这将在 连接/断开连接 时通知您
	peerConnection.OnConnectionStateChange(func(s webrtc.PeerConnectionState) {
		fmt.Printf("连接状态已更改: %s\n", s.String())

		if s == webrtc.PeerConnectionStateFailed {
			// 等待连接在 30 秒内没有网络活动或再次失败。可以使用 ICE 重新启动重新连接它。
			// 使用 webrtc.PeerConnectionStateDisconnected 断开连接，如果您有兴趣检测更快的超时。
			// 请注意，PeerConnection 可能会从 PeerConnectionStateDisconnected 返回。
			fmt.Println("连接已进入退出失败状态")
			os.Exit(0)
		}
	})

	// 注册通道打开处理
	dataChannel.OnOpen(func() {
		fmt.Printf("数据通道 '%s'-'%d' 打开。数据可发送到任何连接的数据通道\n", dataChannel.Label(), dataChannel.ID())

		for range time.NewTicker(5 * time.Second).C {
			message := signal.RandSeq(15)
			fmt.Printf("发送 '%s'\n", message)

			// 以文本形式发送消息
			sendTextErr := dataChannel.SendText(message)
			if sendTextErr != nil {
				panic(sendTextErr)
			}
		}
	})

	// 注册文本消息处理
	dataChannel.OnMessage(func(msg webrtc.DataChannelMessage) {
		fmt.Printf("来自数据通道的消息 '%s': '%s'\n", dataChannel.Label(), string(msg.Data))
	})

	// 创建要发送到其他流程的offer。
	offer, err := peerConnection.CreateOffer(nil)
	if err != nil {
		panic(err)
	}

	// 设置本地描述，并启动我们的 UDP 侦听器
	// 注意：这将开始收集ICE候选人
	if err = peerConnection.SetLocalDescription(offer); err != nil {
		panic(err)
	}

	// 将我们的报价发送到在另一个进程中侦听的 HTTP 服务器
	payload, err := json.Marshal(offer)
	if err != nil {
		panic(err)
	}
	resp, err := http.Post(fmt.Sprintf("http://%s/sdp", *answerAddr), "application/json; charset=utf-8", bytes.NewReader(payload)) // nolint:noctx
	if err != nil {
		panic(err)
	} else if err := resp.Body.Close(); err != nil {
		panic(err)
	}

	// 永远阻止
	select {}
}
