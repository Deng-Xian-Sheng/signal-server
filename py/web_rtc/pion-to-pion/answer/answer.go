//Copyright (c) [2023] [JinCanQi]
//[make_data_set_so-vits-svc] is licensed under Mulan PubL v2.
//You can use this software according to the terms and conditions of the Mulan PubL v2.
//You may obtain a copy of Mulan PubL v2 at:
//         http://license.coscl.org.cn/MulanPubL-2.0
//THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
//EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
//MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
//See the Mulan PubL v2 for more details.

package answer

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

func aaa() { // nolint:gocognit
	offerAddr := flag.String("offer-address", "localhost:50000", "offer HTTP 服务器的地址。")
	answerAddr := flag.String("answer-address", ":60000", "answer HTTP 服务器的地址。")
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
		if err := peerConnection.Close(); err != nil {
			fmt.Printf("无法关闭对等连接: %v\n", err)
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
		} else if onICECandidateErr := signalCandidate(*offerAddr, c); onICECandidateErr != nil {
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
		if err := json.NewDecoder(r.Body).Decode(&sdp); err != nil {
			panic(err)
		}

		if err := peerConnection.SetRemoteDescription(sdp); err != nil {
			panic(err)
		}

		// 创建要发送到其他进程的答案
		answer, err := peerConnection.CreateAnswer(nil)
		if err != nil {
			panic(err)
		}

		// 将我们的答案发送到在另一个进程中侦听的 HTTP 服务器
		payload, err := json.Marshal(answer)
		if err != nil {
			panic(err)
		}
		resp, err := http.Post(fmt.Sprintf("http://%s/sdp", *offerAddr), "application/json; charset=utf-8", bytes.NewReader(payload)) // nolint:noctx
		if err != nil {
			panic(err)
		} else if closeErr := resp.Body.Close(); closeErr != nil {
			panic(closeErr)
		}

		// 设置本地描述，并启动我们的 UDP 侦听器
		err = peerConnection.SetLocalDescription(answer)
		if err != nil {
			panic(err)
		}

		candidatesMux.Lock()
		for _, c := range pendingCandidates {
			onICECandidateErr := signalCandidate(*offerAddr, c)
			if onICECandidateErr != nil {
				panic(onICECandidateErr)
			}
		}
		candidatesMux.Unlock()
	})

	// 为对等连接状态设置处理程序
	// 这将在 连接/断开连接 时通知您
	peerConnection.OnConnectionStateChange(func(s webrtc.PeerConnectionState) {
		fmt.Printf("对等连接状态已更改: %s\n", s.String())

		if s == webrtc.PeerConnectionStateFailed {
			// 等待连接在 30 秒内没有网络活动或再次失败。可以使用 ICE 重新启动重新连接它。
			// 使用 webrtc.PeerConnectionStateDisconnected 断开连接，如果您有兴趣检测更快的超时。
			// 请注意，PeerConnection 可能会从 PeerConnectionStateDisconnected 返回。
			fmt.Println("连接已进入退出失败状态")
			os.Exit(0)
		}
	})

	// 注册数据通道
	peerConnection.OnDataChannel(func(d *webrtc.DataChannel) {
		fmt.Printf("新建数据通道 %s %d\n", d.Label(), d.ID())

		// 数据通道打开
		d.OnOpen(func() {
			fmt.Printf("数据通道 '%s'-'%d' 打开。数据可发送到任何连接的数据通道\n", d.Label(), d.ID())

			for range time.NewTicker(5 * time.Second).C {
				message := signal.RandSeq(15)
				fmt.Printf("发送 '%s'\n", message)

				// 以文本形式发送消息
				sendTextErr := d.SendText(message)
				if sendTextErr != nil {
					panic(sendTextErr)
				}
			}
		})

		// 注册文本消息处理
		d.OnMessage(func(msg webrtc.DataChannelMessage) {
			fmt.Printf("来自数据通道的消息 '%s': '%s'\n", d.Label(), string(msg.Data))
		})
	})

	// 启动 HTTP 服务器，该服务器接受来自 offer 流程的请求以交换 SDP 和 Candidates
	// nolint: gosec
	panic(http.ListenAndServe(*answerAddr, nil))
}
