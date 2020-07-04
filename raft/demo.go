package main

import (
	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
	"io"
	"net"
	"os"
	"path/filepath"
	"time"
)

type MyFsm struct {
}

func (myFsm MyFsm) Apply(log *raft.Log) interface{} {

}

func (myFsm MyFsm) Snapshot() (raft.FSMSnapshot, error) {

}

func (myFsm MyFsm) Restore(closer io.ReadCloser) error {

}

func main() {
	boltStore, _ := raftboltdb.NewBoltStore(filepath.Join())
	stableStore, _ := raftboltdb.NewBoltStore("/raft-stable.db")
	snapshots, _ := raft.NewFileSnapshotStore("", 3, os.Stderr)
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1")
	transport, _ := raft.NewTCPTransport("127.0.0.1", addr, 3, time.Second, os.Stderr)
	raft, _ := raft.NewRaft(
		raft.DefaultConfig(), // 节点配置信息
		MyFsm{},              // 有限状态机
		boltStore,            // 存储 Raft 日志
		stableStore,          // 稳定存储, 用来存储 Raft 节点信息
		snapshots,            // 快照存储, 用来存储节点的快照信息
		transport,            // Raft 节点间消息通道
	)
	
}
