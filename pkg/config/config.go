package config

import(
  "sync"
)

var (
  A int = 0
  B int = 0
  C int = 0

 Lock sync.Mutex = sync.Mutex{}
)
