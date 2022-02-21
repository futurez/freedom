package futils

import (
	"os"
	"os/signal"
	"syscall"
)

func Graceful() {
	exitCh :=  make(chan os.Signal)
	signal.Notify(exitCh, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-exitCh
	signal.Stop(exitCh)
	close(exitCh)
}
