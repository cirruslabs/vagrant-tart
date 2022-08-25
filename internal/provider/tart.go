package provider

import (
	plugincore "github.com/hashicorp/vagrant-plugin-sdk/core"
)

type TartProvider struct {
}

func (t TartProvider) Capability(name string, args ...interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (t TartProvider) HasCapability(name string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (t TartProvider) Usable() (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (t TartProvider) Installed() (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (t TartProvider) Action(name string, args ...interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (t TartProvider) MachineIdChanged() error {
	//TODO implement me
	panic("implement me")
}

func (t TartProvider) SshInfo() (*plugincore.SshInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (t TartProvider) State() (*plugincore.MachineState, error) {
	//TODO implement me
	panic("implement me")
}
