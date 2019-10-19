package mash

import (
	"fmt"
	"io"
	"os"
)

type FSHandler struct {
}

func NewFSHandler() *FSHandler {
	return &FSHandler{}
}

func (h *FSHandler) Pwd() Result {
	wd, err := os.Getwd()
	return NewResult("fs:pwd", wd, err)
}

func (h *FSHandler) Mkdir(d string) Result {
	err := os.Mkdir(d, os.ModePerm)
	return NewResult(fmt.Sprintf("fs:mkdir (%s)", d), "", err)
}

func (h *FSHandler) CopyFile(src, target string) Result {
	ctx := fmt.Sprintf("copy file  (%s) → (%s)", src, target)
	srcF, err := os.Open(src)
	if err != nil {
		return NewResult(ctx, "", err)
	}
	defer srcF.Close()
	targetF, err := os.Create(target)
	if err != nil {
		return NewResult(ctx, "", err)
	}
	defer targetF.Close()
	_, err = io.Copy(targetF, srcF)
	if err != nil {
		return NewResult(ctx, "", err)
	}
	return NewResult(ctx, "", nil)
}
