package selinux

import (
	"fmt"
	"github.com/ctrsploit/ctrsploit/internal"
	"github.com/ctrsploit/ctrsploit/pkg/selinux"
	"github.com/ctrsploit/sploit-spec/pkg/result"
	"github.com/ctrsploit/sploit-spec/pkg/result/item"
)

type Result struct {
	Name       result.Title
	Enabled    item.Bool  `json:"enabled"`
	Mode       item.Short `json:"mode"`
	MountPoint item.Short `json:"mount_point"`
}

func (r Result) String() (s string) {
	s += internal.Print(r.Name, r.Enabled)
	if r.Enabled.Result {
		s += internal.Print(r.Mode, r.MountPoint)
	}
	return
}

func Selinux() (err error) {
	r := Result{
		Name: result.Title{
			Name: "SELinux",
		},
		Enabled: item.Bool{
			Name:        "Enabled",
			Description: "",
			Result:      selinux.IsEnabled(),
		},
	}
	if r.Enabled.Result {
		r.Mode = item.Short{
			Name:        "Mode",
			Description: "",
			Result:      selinux.Mode().String(),
		}
		r.MountPoint = item.Short{
			Name:        "Mount point",
			Description: "",
			Result:      selinux.GetSelinuxMountPoint(),
		}
	}
	fmt.Println(r)
	return
}
