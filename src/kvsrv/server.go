package kvsrv

import (
	"log"
	"sync"
)

const Debug = false

func DPrintf(format string, a ...interface{}) (n int, err error) {
	if Debug {
		log.Printf(format, a...)
	}
	return
}

type KVServer struct {
	mu    sync.Mutex
	kvs   map[string]string
	exist map[int64]string
}

func (kv *KVServer) Get(args *GetArgs, reply *GetReply) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	if args.Type == Committed {
		delete(kv.exist, args.ClientID)
		return
	}
	if val, ok := kv.exist[args.ClientID]; ok {
		//重复请求
		reply.Value = val
		return
	}
	reply.Value = kv.kvs[args.Key]
	kv.exist[args.ClientID] = reply.Value
}

func (kv *KVServer) Put(args *PutAppendArgs, reply *PutAppendReply) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	if args.Type == Committed {
		delete(kv.exist, args.ClientID)
		return
	}
	if _, ok := kv.exist[args.ClientID]; ok {
		//重复请求
		return
	}
	kv.kvs[args.Key] = args.Value
	kv.exist[args.ClientID] = args.Value
}

func (kv *KVServer) Append(args *PutAppendArgs, reply *PutAppendReply) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	if args.Type == Committed {
		delete(kv.exist, args.ClientID)
		return
	}
	if val, ok := kv.exist[args.ClientID]; ok {
		//重复请求
		reply.Value = val
		return
	}
	old := kv.kvs[args.Key]
	kv.kvs[args.Key] += args.Value
	reply.Value = old
	kv.exist[args.ClientID] = old
}

func StartKVServer() *KVServer {
	kv := new(KVServer)
	kv.kvs = make(map[string]string)
	kv.exist = make(map[int64]string)
	return kv
}
