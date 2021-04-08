/*
 * @Author: ph4ntom
 * @Date: 2021-03-23 19:01:26
 * @LastEditors: ph4ntom
 * @LastEditTime: 2021-04-02 17:24:21
 */
package manager

import (
	"Stowaway/share"
)

type Manager struct {
	ConsoleManager *consoleManager
	FileManager    *fileManager
	SocksManager   *socksManager
	ForwardManager *forwardManager
	SSHManager     *sshManager
	ShellManager   *shellManager
	InfoManager    *infoManager
}

func NewManager(file *share.MyFile) *Manager {
	manager := new(Manager)
	manager.ConsoleManager = newConsoleManager()
	manager.FileManager = newFileManager(file)
	manager.SocksManager = newSocksManager()
	manager.ForwardManager = newForwardManager()
	manager.SSHManager = newSSHManager()
	manager.ShellManager = newShellManager()
	manager.InfoManager = newInfoManager()
	return manager
}

func (manager *Manager) Run() {
	go manager.SocksManager.run()
	go manager.ForwardManager.run()
}
