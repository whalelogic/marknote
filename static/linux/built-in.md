# Linux Built-in Commands Cheatsheet

Core commands every Linux user relies on daily: filesystem navigation, text
processing, process management, permissions, and system inspection.


| Category | Command | Description / Usage |
| --- | --- | --- |
| **Navigation & Filesystem** | `pwd` | print working directory |
| **Navigation & Filesystem** | `cd /path/to/dir` | change directory |
| **Navigation & Filesystem** | `cd -` | go to previous directory |
| **Navigation & Filesystem** | `cd ~` | go to home directory |
| **Navigation & Filesystem** | `ls -la` | list all files, long format |
| **Navigation & Filesystem** | `ls -lh` | human-readable sizes |
| **Navigation & Filesystem** | `ls -lt` | sort by modification time |
| **Navigation & Filesystem** | `ls -lS` | sort by size |
| **Navigation & Filesystem** | `tree -L 2` | show directory tree, 2 levels deep |
| **Navigation & Filesystem** | `find / -name "*.log"` | find files by name |
| **Navigation & Filesystem** | `find . -type f -mtime -7` | files modified in last 7 days |
| **Navigation & Filesystem** | `find . -type d -empty` | empty directories |
| **Navigation & Filesystem** | `find . -size +100M` | files larger than 100MB |
| **File Operations** | `cp file1 file2` | copy file |
| **File Operations** | `cp -r dir1 dir2` | copy directory recursively |
| **File Operations** | `cp -a src dst` | archive copy (preserves attributes) |
| **File Operations** | `mv old new` | move/rename |
| **File Operations** | `rm file` | remove file |
| **File Operations** | `rm -rf dir` | remove directory recursively, force |
| **File Operations** | `mkdir -p a/b/c` | create nested directories |
| **File Operations** | `rmdir dir` | remove empty directory |
| **File Operations** | `touch file` | create empty file / update timestamp |
| **File Operations** | `ln -s target linkname` | create symbolic link |
| **File Operations** | `ln target linkname` | create hard link |
| **File Operations** | `stat file` | detailed file metadata |
| **File Operations** | `file file` | detect file type |
| **File Operations** | `readlink -f symlink` | resolve symlink to absolute path |
| **Viewing File Contents** | `cat file` | print entire file |
| **Viewing File Contents** | `cat -n file` | print with line numbers |
| **Viewing File Contents** | `tac file` | print file in reverse |
| **Viewing File Contents** | `less file` | paginated viewer (q to quit) |
| **Viewing File Contents** | `more file` | simpler pager |
| **Viewing File Contents** | `head file` | first 10 lines |
| **Viewing File Contents** | `head -n 20 file` | first 20 lines |
| **Viewing File Contents** | `tail file` | last 10 lines |
| **Viewing File Contents** | `tail -n 50 file` | last 50 lines |
| **Viewing File Contents** | `tail -f file` | follow file as it grows (logs) |
| **Viewing File Contents** | `tail -F file` | follow, re-attach if file rotates |
| **Viewing File Contents** | `wc -l file` | count lines |
| **Viewing File Contents** | `wc -w file` | count words |
| **Viewing File Contents** | `wc -c file` | count bytes |
| **Text Processing** | `grep "pattern" file` | search for pattern |
| **Text Processing** | `grep -i "pattern" file` | case-insensitive |
| **Text Processing** | `grep -r "pattern" dir/` | recursive search |
| **Text Processing** | `grep -v "pattern" file` | invert match (exclude) |
| **Text Processing** | `grep -n "pattern" file` | show line numbers |
| **Text Processing** | `grep -c "pattern" file` | count matches |
| **Text Processing** | `grep -E "pat1|pat2" file` | extended regex (OR) |
| **Text Processing** | `grep -A 3 -B 3 "pattern" file` | 3 lines after/before match |
| **Text Processing** | `sed 's/foo/bar/' file` | replace first occurrence per line |
| **Text Processing** | `sed 's/foo/bar/g' file` | replace all occurrences |
| **Text Processing** | `sed -i 's/foo/bar/g' file` | edit file in place |
| **Text Processing** | `sed -n '5,10p' file` | print lines 5-10 |
| **Text Processing** | `sed '/pattern/d' file` | delete matching lines |
| **Text Processing** | `sort file` | sort lines alphabetically |
| **Text Processing** | `sort -n file` | numeric sort |
| **Text Processing** | `sort -r file` | reverse sort |
| **Text Processing** | `sort -k2 file` | sort by 2nd field |
| **Text Processing** | `sort -u file` | sort and remove duplicates |
| **Text Processing** | `uniq file` | remove adjacent duplicate lines |
| **Text Processing** | `uniq -c file` | count occurrences |
| **Text Processing** | `cut -d',' -f1,3 file` | extract fields 1 and 3 (CSV) |
| **Text Processing** | `cut -c1-10 file` | extract characters 1-10 |
| **Text Processing** | `tr 'a-z' 'A-Z' < file` | translate lowercase to uppercase |
| **Text Processing** | `tr -d '\n' < file` | delete newlines |
| **Text Processing** | `paste file1 file2` | merge lines side by side |
| **Text Processing** | `diff file1 file2` | show differences |
| **Text Processing** | `diff -u file1 file2` | unified diff format |
| **Permissions & Ownership** | `chmod 755 file` | rwxr-xr-x |
| **Permissions & Ownership** | `chmod +x script.sh` | add execute permission |
| **Permissions & Ownership** | `chmod -R 644 dir/` | recursive permission change |
| **Permissions & Ownership** | `chown user:group file` | change owner and group |
| **Permissions & Ownership** | `chown -R user dir/` | recursive ownership change |
| **Permissions & Ownership** | `chgrp group file` | change group only |
| **Permissions & Ownership** | `umask` | show default permission mask |
| **Permissions & Ownership** | `umask 022` | set default permission mask |
| **Permissions & Ownership** | `r / w / x` | Symbol Reference: Read (4) / Write (2) / Execute (1) |
| **Permissions & Ownership** | `rwx / rw- / r-x` | Octal Reference: All (7) / Read + Write (6) / Read + Execute (5) |
| **Process Management** | `ps aux` | list all running processes |
| **Process Management** | `ps -ef` | alternate format |
| **Process Management** | `top` | live process viewer |
| **Process Management** | `htop` | improved interactive process viewer |
| **Process Management** | `kill PID` | terminate process (SIGTERM) |
| **Process Management** | `kill -9 PID` | force kill (SIGKILL) |
| **Process Management** | `killall processname` | kill by name |
| **Process Management** | `pkill -f "pattern"` | kill by matching command line |
| **Process Management** | `jobs` | list background jobs |
| **Process Management** | `bg` | resume job in background |
| **Process Management** | `fg` | bring job to foreground |
| **Process Management** | `nohup command &` | run immune to hangups |
| **Process Management** | `disown` | detach job from shell |
| **Process Management** | `nice -n 10 command` | run with lower priority |
| **Process Management** | `renice -n 5 -p PID` | change priority of running process |
| **System Information** | `uname -a` | kernel and system info |
| **System Information** | `hostname` | show hostname |
| **System Information** | `uptime` | system uptime and load average |
| **System Information** | `free -h` | memory usage, human-readable |
| **System Information** | `df -h` | disk space usage |
| **System Information** | `du -sh dir/` | directory size summary |
| **System Information** | `du -sh * | sort -h` | size of each item, sorted |
| **System Information** | `lscpu` | CPU information |
| **System Information** | `lsblk` | list block devices |
| **System Information** | `mount` | show mounted filesystems |
| **System Information** | `who` | logged in users |
| **System Information** | `w` | who + what they're doing |
| **System Information** | `last` | login history |
| **System Information** | `whoami` | current user |
| **System Information** | `id` | user and group IDs |
| **Networking** | `ip a` | show network interfaces |
| **Networking** | `ip route` | show routing table |
| **Networking** | `ping host` | test connectivity |
| **Networking** | `curl -I [https://example.com](https://example.com)` | fetch HTTP headers |
| **Networking** | `netstat -tulpn` | listening ports (older systems) |
| **Networking** | `ss -tulpn` | listening ports (modern replacement) |
| **Networking** | `dig example.com` | DNS lookup |
| **Networking** | `nslookup example.com` | DNS lookup, alternate tool |
| **Networking** | `traceroute host` | trace network path |
| **Networking** | `hostname -I` | local IP addresses |
| **Archiving & Compression** | `tar -cvf archive.tar dir/` | create tar archive |
| **Archiving & Compression** | `tar -xvf archive.tar` | extract tar archive |
| **Archiving & Compression** | `tar -czvf archive.tar.gz dir/` | create gzip-compressed archive |
| **Archiving & Compression** | `tar -xzvf archive.tar.gz` | extract gzip archive |
| **Archiving & Compression** | `tar -cjvf archive.tar.bz2 dir/` | create bzip2 archive |
| **Archiving & Compression** | `tar -tvf archive.tar` | list contents without extracting |
| **Archiving & Compression** | `zip -r archive.zip dir/` | create zip archive |
| **Archiving & Compression** | `unzip archive.zip` | extract zip archive |
| **Archiving & Compression** | `gzip file` | compress single file |
| **Archiving & Compression** | `gunzip file.gz` | decompress |
| **Environment & Shell** | `echo $PATH` | show PATH variable |
| **Environment & Shell** | `export VAR=value` | set environment variable |
| **Environment & Shell** | `env` | list all environment variables |
| **Environment & Shell** | `alias ll='ls -la'` | create alias |
| **Environment & Shell** | `unalias ll` | remove alias |
| **Environment & Shell** | `which command` | show path to executable |
| **Environment & Shell** | `type command` | show how shell interprets command |
| **Environment & Shell** | `history` | show command history |
| **Environment & Shell** | `!!` | rerun last command |
| **Environment & Shell** | `!123` | rerun history item 123 |
| **Environment & Shell** | `source ~/.bashrc` | reload shell config |
| **Searching & Locating** | `locate filename` | fast file search (needs updatedb) |
| **Searching & Locating** | `updatedb` | update locate database |
| **Searching & Locating** | `whereis command` | locate binary, source, man page |
| **Searching & Locating** | `find / -perm -4000` | find setuid files |
| **Disk & Filesystem** | `fdisk -l` | list disk partitions |
| **Disk & Filesystem** | `mkfs.ext4 /dev/sdX1` | format partition |
| **Disk & Filesystem** | `mount /dev/sdX1 /mnt` | mount a filesystem |
| **Disk & Filesystem** | `umount /mnt` | unmount |
| **Disk & Filesystem** | `fsck /dev/sdX1` | check filesystem |
| **Useful Combos** | `ps aux | grep nginx` | find process by name |
| **Useful Combos** | `du -sh * | sort -rh | head -10` | top 10 largest items |
| **Useful Combos** | `find . -name "*.tmp" -delete` | delete matching files |
| **Useful Combos** | `history | grep ssh` | find past ssh commands |
| **Useful Combos** | `cat /etc/passwd | cut -d: -f1` | list all usernames |
| **Useful Combos** | `df -h | awk '$5+0 > 80 {print}'` | partitions over 80% full |
