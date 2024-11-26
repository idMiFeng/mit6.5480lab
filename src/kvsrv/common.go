package kvsrv

const (
	Pending = iota
	Committed
)

// Put or Append
type PutAppendArgs struct {
	Key      string
	Value    string
	Type     int
	ClientID int64
	// You'll have to add definitions here.
	// Field names must start with capital letters,
	// otherwise RPC will break.
}

type PutAppendReply struct {
	Value string
}

type GetArgs struct {
	Key      string
	Type     int
	ClientID int64
	// You'll have to add definitions here.
}

type GetReply struct {
	Value string
}
