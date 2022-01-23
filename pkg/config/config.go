package config

import(
  "sync"
)

const MAX = 12200160415121876738 // 93rd term of fibonnaci after this there is over flow from uint64

var (
  A uint64 = 0
  B uint64 = 0
  C uint64 = 0

 Lock sync.Mutex = sync.Mutex{}
)
