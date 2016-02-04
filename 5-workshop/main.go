package main

import (
	"path/filepath"

	"github.com/docker/go-plugins-helpers/volume"
	"github.com/vdemeester/docker-volume-ipfs/ipfs"
)

const ipfsID = "_ipfs"

var (
	defaultDir = filepath.Join(volume.DefaultDockerRootDirectory, ipfsID)
)

func main() {
	var daemon *ipfs.Daemon

	// […]
}
