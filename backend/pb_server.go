package main

import (
	"log"
	"sync"

	"github.com/pocketbase/pocketbase"
)

var (
	pbInstance *pocketbase.PocketBase
	pbOnce     sync.Once
)

// StartPocketBase creates and starts a PocketBase instance in a goroutine.
// It returns the PocketBase instance once it's created.
func StartPocketBase() *pocketbase.PocketBase {
	pbOnce.Do(func() {
		pb := pocketbase.New()
		pbInstance = pb

		// start pocketbase in a separate goroutine so it doesn't block Gin
		go func() {
			if err := pb.Start(); err != nil {
				log.Fatalf("pocketbase start error: %v", err)
			}
		}()
	})

	return pbInstance
}

// GetPocketBase returns the running PocketBase instance (may be nil until started)
func GetPocketBase() *pocketbase.PocketBase {
	return pbInstance
}
