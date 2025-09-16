package oservice

var OsServiceReference = MapOsService{
	"apparmor": {Name: "apparmor", CName: "apparmor.service"},
	"crio":     {Name: "crio", CName: "crio.service"},
	"kubelet":  {Name: "kubelet", CName: "kubelet.service"},
}
