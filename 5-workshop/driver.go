package main

import (
	"github.com/docker/go-plugins-helpers/volume"
)

type ipfsDriver struct {
	// […]
}

func newIPFSDriver(ipfsMountPoint, ipnsMountPoint string) ipfsDriver {
	return ipfsDriver{}
}

func (d ipfsDriver) Create(r volume.Request) volume.Response {
	// […]
	return volume.Response{}
}

func (d ipfsDriver) Path(r volume.Request) volume.Response {
	// […]
	return volume.Response{}
}

func (d ipfsDriver) Remove(r volume.Request) volume.Response {
	// […]
	return volume.Response{}
}

func (d ipfsDriver) Mount(r volume.Request) volume.Response {
	// […]
	return volume.Response{}
}

func (d ipfsDriver) Unmount(r volume.Request) volume.Response {
	// […]
	return volume.Response{}
}

func (d ipfsDriver) Get(r volume.Request) volume.Response {
	// […]
	return volume.Response{}
}

func (d ipfsDriver) List(r volume.Request) volume.Response {
	// […]
	return volume.Response{}
}
