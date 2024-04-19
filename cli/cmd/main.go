package main

import (
	"github.com/GalichAnton/chat-server/cli/cmd/root"
	"github.com/GalichAnton/platform_common/pkg/closer"
)

func main() {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	root.Execute()
}
