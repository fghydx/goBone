package Windows

import (
	"github.com/kardianos/service"
	"log"
	"os"
)

/*
例子:
func main() {
	CreateServicesF("TestService", "测试服务", "这是一个测试服务", func(){服务执行逻辑}, func(){服务停止时逻辑})
}

*/

type SrvIntf interface {
	Onstart()
	Onstop()
}

type WinServices struct {
	srvIntf SrvIntf
	start   func()
	stop    func()
}

func (Srv *WinServices) Start(s service.Service) error {
	if Srv.srvIntf != nil {
		go Srv.srvIntf.Onstart()
	}
	if Srv.start != nil {
		go Srv.start()
	}
	return nil
}

func (Srv *WinServices) Stop(s service.Service) error {
	if Srv.srvIntf != nil {
		Srv.srvIntf.Onstop()
	}
	if Srv.stop != nil {
		Srv.stop()
	}
	return nil
}

func CreateServicesI(name, displayname, description string, srvintf SrvIntf) {
	createNewServices(name, displayname, description, srvintf, nil, nil)
}

func CreateServicesF(name, displayname, description string, start func(), stop func()) {
	createNewServices(name, displayname, description, nil, start, stop)
}

func createNewServices(name, displayname, description string, srvintf SrvIntf, start func(), stop func()) {
	svcConfig := &service.Config{
		Name:        name,
		DisplayName: displayname,
		Description: description,
	}

	s := &WinServices{nil, nil, nil}
	if srvintf != nil {
		s.srvIntf = srvintf
	} else {
		s.start = start
		s.stop = stop
	}

	sys := service.ChosenSystem()
	srv, err := sys.New(s, svcConfig)

	if err != nil {
		log.Printf("Set logger error:%s\n", err.Error())
	}
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "install":
			err := srv.Install()
			if err != nil {
				log.Fatalf("Install service error:%s\n", err.Error())
			}
		case "uninstall":
			err := srv.Uninstall()
			if err != nil {
				log.Fatalf("Uninstall service error:%s\n", err.Error())
			}
		}
		return
	}
	err = srv.Run()
	if err != nil {
		log.Fatalf("Run programe error:%s\n", err.Error())
	}
}
