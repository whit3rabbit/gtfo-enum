package main

import (
	"fmt"
	"os/exec"
	"os/user"
	"regexp"
	"syscall"
	"os"
)

func CheckUidForFileInfo(fi os.FileInfo) int {
	return int(fi.Sys().(*syscall.Stat_t).Uid)
}

func CheckGidForFileInfo(fi os.FileInfo) int {
	return int(fi.Sys().(*syscall.Stat_t).Uid)
}

func checkIsExecutable(info os.FileInfo) bool {

	mode := info.Mode()

	uid:= CheckUidForFileInfo(info)
	gid := CheckGidForFileInfo(info)

	userInfo, err := user.Current()
	if err != nil {
		fmt.Println("Unable to get current information user")
	}


	// Return true or false for file permission bitmask
	isExecOwner := mode&0100 != 0  // Only owner (user) can run
	isExecGroup := mode&0010 != 0  // Only user in group can run
	isExecOther := mode&0001 != 0  // Anyone can run
	isExecAny :=  mode&0111 != 0   // Anyone can run

	if isExecOwner { // Check if owner value is set
		if fmt.Sprint(uid) != userInfo.Uid { // Check if the uid is NOT the same as current user ID
			isExecOwner = false
		}
	}

	if isExecGroup {
		if fmt.Sprint(gid) != userInfo.Gid { // Check if the guid is NOT the same as current group ID
			isExecGroup = false
		}
	}
	switch {
		case isExecOther:
			return isExecOther
		case isExecAny:
			return isExecAny
		case isExecOwner:
			return isExecOwner
		case isExecGroup:
			return isExecGroup
		default:
			return false
	}
}

func checkIsStickyBitSet(fileMode os.FileMode) bool {

	fileString := fmt.Sprintf("%s", fileMode)

	patternRegex := "^u" // This checks if sticky bit is set as first character
	isStickyBitSet, _ := regexp.MatchString(patternRegex, fileString)  // Return true if it is set

	if isStickyBitSet {
		return true
	}

	return false

}

// Checks if binary is in env path
func checkInPath(path string) (string, bool, error) {

        binaryPath, err := exec.LookPath(path)
        if err != nil {
                return path, false, err
        }
        return binaryPath, true, err
}


func main() {

	gtfobins := map[string]string{
		"apt" : "https://gtfobins.github.io/gtfobins/apt/",
		"apt-get" : "https://gtfobins.github.io/gtfobins/apt-get/",
		"aria2c" : "https://gtfobins.github.io/gtfobins/aria2c/",
		"arp" : "https://gtfobins.github.io/gtfobins/arp/",
		"ash" : "https://gtfobins.github.io/gtfobins/ash/",
		"awk" : "https://gtfobins.github.io/gtfobins/awk/",
		"base32" : "https://gtfobins.github.io/gtfobins/base32/",
		"base64" : "https://gtfobins.github.io/gtfobins/base64/",
		"bash" : "https://gtfobins.github.io/gtfobins/bash/",
		"bpftrace" : "https://gtfobins.github.io/gtfobins/bpftrace/",
		"bundler" : "https://gtfobins.github.io/gtfobins/bundler/",
		"busctl" : "https://gtfobins.github.io/gtfobins/busctl/",
		"busybox" : "https://gtfobins.github.io/gtfobins/busybox/",
		"byebug" : "https://gtfobins.github.io/gtfobins/byebug/",
		"cancel" : "https://gtfobins.github.io/gtfobins/cancel/",
		"cat" : "https://gtfobins.github.io/gtfobins/cat/",
		"chmod" : "https://gtfobins.github.io/gtfobins/chmod/",
		"chown" : "https://gtfobins.github.io/gtfobins/chown/",
		"chroot" : "https://gtfobins.github.io/gtfobins/chroot/",
		"cobc" : "https://gtfobins.github.io/gtfobins/cobc/",
		"cp" : "https://gtfobins.github.io/gtfobins/cp/",
		"cpan" : "https://gtfobins.github.io/gtfobins/cpan/",
		"cpulimit" : "https://gtfobins.github.io/gtfobins/cpulimit/",
		"crash" : "https://gtfobins.github.io/gtfobins/crash/",
		"crontab" : "https://gtfobins.github.io/gtfobins/crontab/",
		"csh" : "https://gtfobins.github.io/gtfobins/csh/",
		"curl" : "https://gtfobins.github.io/gtfobins/curl/",
		"cut" : "https://gtfobins.github.io/gtfobins/cut/",
		"dash" : "https://gtfobins.github.io/gtfobins/dash/",
		"date" : "https://gtfobins.github.io/gtfobins/date/",
		"dd" : "https://gtfobins.github.io/gtfobins/dd/",
		"dialog" : "https://gtfobins.github.io/gtfobins/dialog/",
		"diff" : "https://gtfobins.github.io/gtfobins/diff/",
		"dmesg" : "https://gtfobins.github.io/gtfobins/dmesg/",
		"dmsetup" : "https://gtfobins.github.io/gtfobins/dmsetup/",
		"dnf" : "https://gtfobins.github.io/gtfobins/dnf/",
		"docker" : "https://gtfobins.github.io/gtfobins/docker/",
		"dpkg" : "https://gtfobins.github.io/gtfobins/dpkg/",
		"easy_install" : "https://gtfobins.github.io/gtfobins/easy_install/",
		"eb" : "https://gtfobins.github.io/gtfobins/eb/",
		"ed" : "https://gtfobins.github.io/gtfobins/ed/",
		"emacs" : "https://gtfobins.github.io/gtfobins/emacs/",
		"env" : "https://gtfobins.github.io/gtfobins/env/",
		"eqn" : "https://gtfobins.github.io/gtfobins/eqn/",
		"expand" : "https://gtfobins.github.io/gtfobins/expand/",
		"expect" : "https://gtfobins.github.io/gtfobins/expect/",
		"facter" : "https://gtfobins.github.io/gtfobins/facter/",
		"file" : "https://gtfobins.github.io/gtfobins/file/",
		"find" : "https://gtfobins.github.io/gtfobins/find/",
		"finger" : "https://gtfobins.github.io/gtfobins/finger/",
		"flock" : "https://gtfobins.github.io/gtfobins/flock/",
		"fmt" : "https://gtfobins.github.io/gtfobins/fmt/",
		"fold" : "https://gtfobins.github.io/gtfobins/fold/",
		"ftp" : "https://gtfobins.github.io/gtfobins/ftp/",
		"gawk" : "https://gtfobins.github.io/gtfobins/gawk/",
		"gcc" : "https://gtfobins.github.io/gtfobins/gcc/",
		"gdb" : "https://gtfobins.github.io/gtfobins/gdb/",
		"gem" : "https://gtfobins.github.io/gtfobins/gem/",
		"genisoimage" : "https://gtfobins.github.io/gtfobins/genisoimage/",
		"gimp" : "https://gtfobins.github.io/gtfobins/gimp/",
		"git" : "https://gtfobins.github.io/gtfobins/git/",
		"grep" : "https://gtfobins.github.io/gtfobins/grep/",
		"gtester" : "https://gtfobins.github.io/gtfobins/gtester/",
		"hd" : "https://gtfobins.github.io/gtfobins/hd/",
		"head" : "https://gtfobins.github.io/gtfobins/head/",
		"hexdump" : "https://gtfobins.github.io/gtfobins/hexdump/",
		"highlight" : "https://gtfobins.github.io/gtfobins/highlight/",
		"iconv" : "https://gtfobins.github.io/gtfobins/iconv/",
		"iftop" : "https://gtfobins.github.io/gtfobins/iftop/",
		"ionice" : "https://gtfobins.github.io/gtfobins/ionice/",
		"ip" : "https://gtfobins.github.io/gtfobins/ip/",
		"irb" : "https://gtfobins.github.io/gtfobins/irb/",
		"jjs" : "https://gtfobins.github.io/gtfobins/jjs/",
		"journalctl" : "https://gtfobins.github.io/gtfobins/journalctl/",
		"jq" : "https://gtfobins.github.io/gtfobins/jq/",
		"jrunscript" : "https://gtfobins.github.io/gtfobins/jrunscript/",
		"ksh" : "https://gtfobins.github.io/gtfobins/ksh/",
		"ksshell" : "https://gtfobins.github.io/gtfobins/ksshell/",
		"ld.so" : "https://gtfobins.github.io/gtfobins/ld.so/",
		"ldconfig" : "https://gtfobins.github.io/gtfobins/ldconfig/",
		"less" : "https://gtfobins.github.io/gtfobins/less/",
		"logsave" : "https://gtfobins.github.io/gtfobins/logsave/",
		"look" : "https://gtfobins.github.io/gtfobins/look/",
		"ltrace" : "https://gtfobins.github.io/gtfobins/ltrace/",
		"lua" : "https://gtfobins.github.io/gtfobins/lua/",
		"lwp-download" : "https://gtfobins.github.io/gtfobins/lwp-download/",
		"lwp-request" : "https://gtfobins.github.io/gtfobins/lwp-request/",
		"mail" : "https://gtfobins.github.io/gtfobins/mail/",
		"make" : "https://gtfobins.github.io/gtfobins/make/",
		"man" : "https://gtfobins.github.io/gtfobins/man/",
		"mawk" : "https://gtfobins.github.io/gtfobins/mawk/",
		"more" : "https://gtfobins.github.io/gtfobins/more/",
		"mount" : "https://gtfobins.github.io/gtfobins/mount/",
		"mtr" : "https://gtfobins.github.io/gtfobins/mtr/",
		"mv" : "https://gtfobins.github.io/gtfobins/mv/",
		"mysql" : "https://gtfobins.github.io/gtfobins/mysql/",
		"nano" : "https://gtfobins.github.io/gtfobins/nano/",
		"nawk" : "https://gtfobins.github.io/gtfobins/nawk/",
		"nc" : "https://gtfobins.github.io/gtfobins/nc/",
		"nice" : "https://gtfobins.github.io/gtfobins/nice/",
		"nl" : "https://gtfobins.github.io/gtfobins/nl/",
		"nmap" : "https://gtfobins.github.io/gtfobins/nmap/",
		"node" : "https://gtfobins.github.io/gtfobins/node/",
		"nohup" : "https://gtfobins.github.io/gtfobins/nohup/",
		"nroff" : "https://gtfobins.github.io/gtfobins/nroff/",
		"nsenter" : "https://gtfobins.github.io/gtfobins/nsenter/",
		"od" : "https://gtfobins.github.io/gtfobins/od/",
		"openssl" : "https://gtfobins.github.io/gtfobins/openssl/",
		"pdb" : "https://gtfobins.github.io/gtfobins/pdb/",
		"perl" : "https://gtfobins.github.io/gtfobins/perl/",
		"pg" : "https://gtfobins.github.io/gtfobins/pg/",
		"php" : "https://gtfobins.github.io/gtfobins/php/",
		"pic" : "https://gtfobins.github.io/gtfobins/pic/",
		"pico" : "https://gtfobins.github.io/gtfobins/pico/",
		"pip" : "https://gtfobins.github.io/gtfobins/pip/",
		"pry" : "https://gtfobins.github.io/gtfobins/pry/",
		"puppet" : "https://gtfobins.github.io/gtfobins/puppet/",
		"python" : "https://gtfobins.github.io/gtfobins/python/",
		"rake" : "https://gtfobins.github.io/gtfobins/rake/",
		"readelf" : "https://gtfobins.github.io/gtfobins/readelf/",
		"red" : "https://gtfobins.github.io/gtfobins/red/",
		"redcarpet" : "https://gtfobins.github.io/gtfobins/redcarpet/",
		"restic" : "https://gtfobins.github.io/gtfobins/restic/",
		"rlogin" : "https://gtfobins.github.io/gtfobins/rlogin/",
		"rlwrap" : "https://gtfobins.github.io/gtfobins/rlwrap/",
		"rpm" : "https://gtfobins.github.io/gtfobins/rpm/",
		"rpmquery" : "https://gtfobins.github.io/gtfobins/rpmquery/",
		"rsync" : "https://gtfobins.github.io/gtfobins/rsync/",
		"ruby" : "https://gtfobins.github.io/gtfobins/ruby/",
		"run-mailcap" : "https://gtfobins.github.io/gtfobins/run-mailcap/",
		"run-parts" : "https://gtfobins.github.io/gtfobins/run-parts/",
		"rvim" : "https://gtfobins.github.io/gtfobins/rvim/",
		"scp" : "https://gtfobins.github.io/gtfobins/scp/",
		"screen" : "https://gtfobins.github.io/gtfobins/screen/",
		"script" : "https://gtfobins.github.io/gtfobins/script/",
		"sed" : "https://gtfobins.github.io/gtfobins/sed/",
		"service" : "https://gtfobins.github.io/gtfobins/service/",
		"setarch" : "https://gtfobins.github.io/gtfobins/setarch/",
		"sftp" : "https://gtfobins.github.io/gtfobins/sftp/",
		"shuf" : "https://gtfobins.github.io/gtfobins/shuf/",
		"smbclient" : "https://gtfobins.github.io/gtfobins/smbclient/",
		"socat" : "https://gtfobins.github.io/gtfobins/socat/",
		"soelim" : "https://gtfobins.github.io/gtfobins/soelim/",
		"sort" : "https://gtfobins.github.io/gtfobins/sort/",
		"sqlite3" : "https://gtfobins.github.io/gtfobins/sqlite3/",
		"ssh" : "https://gtfobins.github.io/gtfobins/ssh/",
		"start-stop-daemon" : "https://gtfobins.github.io/gtfobins/start-stop-daemon/",
		"stdbuf" : "https://gtfobins.github.io/gtfobins/stdbuf/",
		"strace" : "https://gtfobins.github.io/gtfobins/strace/",
		"strings" : "https://gtfobins.github.io/gtfobins/strings/",
		"systemctl" : "https://gtfobins.github.io/gtfobins/systemctl/",
		"tac" : "https://gtfobins.github.io/gtfobins/tac/",
		"tail" : "https://gtfobins.github.io/gtfobins/tail/",
		"tar" : "https://gtfobins.github.io/gtfobins/tar/",
		"taskset" : "https://gtfobins.github.io/gtfobins/taskset/",
		"tclsh" : "https://gtfobins.github.io/gtfobins/tclsh/",
		"tcpdump" : "https://gtfobins.github.io/gtfobins/tcpdump/",
		"tee" : "https://gtfobins.github.io/gtfobins/tee/",
		"telnet" : "https://gtfobins.github.io/gtfobins/telnet/",
		"tftp" : "https://gtfobins.github.io/gtfobins/tftp/",
		"time" : "https://gtfobins.github.io/gtfobins/time/",
		"timeout" : "https://gtfobins.github.io/gtfobins/timeout/",
		"tmux" : "https://gtfobins.github.io/gtfobins/tmux/",
		"top" : "https://gtfobins.github.io/gtfobins/top/",
		"ul" : "https://gtfobins.github.io/gtfobins/ul/",
		"unexpand" : "https://gtfobins.github.io/gtfobins/unexpand/",
		"uniq" : "https://gtfobins.github.io/gtfobins/uniq/",
		"unshare" : "https://gtfobins.github.io/gtfobins/unshare/",
		"uudecode" : "https://gtfobins.github.io/gtfobins/uudecode/",
		"uuencode" : "https://gtfobins.github.io/gtfobins/uuencode/",
		"valgrind" : "https://gtfobins.github.io/gtfobins/valgrind/",
		"vi" : "https://gtfobins.github.io/gtfobins/vi/",
		"vim" : "https://gtfobins.github.io/gtfobins/vim/",
		"watch" : "https://gtfobins.github.io/gtfobins/watch/",
		"wget" : "https://gtfobins.github.io/gtfobins/wget/",
		"whois" : "https://gtfobins.github.io/gtfobins/whois/",
		"wish" : "https://gtfobins.github.io/gtfobins/wish/",
		"xargs" : "https://gtfobins.github.io/gtfobins/xargs/",
		"xxd" : "https://gtfobins.github.io/gtfobins/xxd/",
		"yelp" : "https://gtfobins.github.io/gtfobins/yelp/",
		"yum" : "https://gtfobins.github.io/gtfobins/yum/",
		"zip" : "https://gtfobins.github.io/gtfobins/zip/",
		"zsh" : "https://gtfobins.github.io/gtfobins/zsh/",
		"zsoelim" : "https://gtfobins.github.io/gtfobins/zsoelim/",
		"zypper" : "https://gtfobins.github.io/gtfobins/zypper/",
	}
	
	for key, value := range gtfobins {
		path, exists, err := checkInPath(key)
		if err != nil {
			continue
		}

		stickybit := false
		executable := false

		if exists {
			info, err := os.Stat(path)
			if err != nil {
				fmt.Println(err)
			}
			
			if checkIsStickyBitSet(info.Mode()) {
				stickybit = true
			}
			
			if checkIsExecutable(info) {
				executable = true
			}
		
			
			if stickybit && executable {
				fmt.Println("SUID and executable: ", path, value)
			}
			
			if executable {
        fmt.Println("Executable: ", path, value)
			}
			
			//fmt.Println(path, value) // Debug or show all
		}
	}
}
