package ffmpeg

import (
	"os"
	"os/exec"
)

type ffmpegExecutor struct {
	Builder QueryBuilder
}

func NewExecutor(builder QueryBuilder) *ffmpegExecutor {
	return &ffmpegExecutor{Builder: builder}
}

func (f *ffmpegExecutor) Run() error {
	cmd := exec.Command(f.Builder.Program(), f.Builder.Args()...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
