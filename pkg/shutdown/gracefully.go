package shutdown

import (
	"os"
	"os/signal"
	"syscall"
)

func Gracefully() {
  sigCh := make(chan os.Signal, 1)
  defer close(sigCh)
  signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
  <-sigCh
}
