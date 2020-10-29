// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build js,wasm

package syscall

// Provided by package runtime.
func now() (sec int64, nsec int32)

func Open(path string, openmode int, perm uint32) (int, error) {
	return 0, ENOSYS
}

func Close(fd int) error {
	return ENOSYS
}

func CloseOnExec(fd int) {
	// nothing to do - no exec
}

func Mkdir(path string, perm uint32) error {
	return ENOSYS
}

func ReadDirent(fd int, buf []byte) (int, error) {
	return 0, ENOSYS
}

func Stat(path string, st *Stat_t) error {
	return ENOSYS
}

func Lstat(path string, st *Stat_t) error {
	return ENOSYS
}

func Fstat(fd int, st *Stat_t) error {
	return ENOSYS
}

func Unlink(path string) error {
	return ENOSYS
}

func Rmdir(path string) error {
	return ENOSYS
}

func Chmod(path string, mode uint32) error {
	return ENOSYS
}

func Fchmod(fd int, mode uint32) error {
	return ENOSYS
}

func Chown(path string, uid, gid int) error {
	return ENOSYS
}

func Fchown(fd int, uid, gid int) error {
	return ENOSYS
}

func Lchown(path string, uid, gid int) error {
	return ENOSYS
}

func UtimesNano(path string, ts []Timespec) error {
	return ENOSYS
}

func Rename(from, to string) error {
	return ENOSYS
}

func Truncate(path string, length int64) error {
	return ENOSYS
}

func Ftruncate(fd int, length int64) error {
	return ENOSYS
}

func Getcwd(buf []byte) (n int, err error) {
	return 0, ENOSYS
}

func Chdir(path string) (err error) {
	return ENOSYS
}

func Fchdir(fd int) error {
	return ENOSYS
}

func Readlink(path string, buf []byte) (n int, err error) {
	return 0, ENOSYS
}

func Link(path, link string) error {
	return ENOSYS
}

func Symlink(path, link string) error {
	return ENOSYS
}

func Fsync(fd int) error {
	return ENOSYS
}

func Read(fd int, b []byte) (int, error) {
	return 0, ENOSYS
}

func Write(fd int, b []byte) (int, error) {
	return 0, ENOSYS
}

func Pread(fd int, b []byte, offset int64) (int, error) {
	return 0, ENOSYS
}

func Pwrite(fd int, b []byte, offset int64) (int, error) {
	return 0, ENOSYS
}

func Seek(fd int, offset int64, whence int) (int64, error) {
	return 0, ENOSYS
}

func Dup(fd int) (int, error) {
	return 0, ENOSYS
}

func Dup2(fd, newfd int) error {
	return ENOSYS
}

func Pipe(fd []int) error {
	return ENOSYS
}

// checkPath checks that the path is not empty and that it contains no null characters.
func checkPath(path string) error {
	if path == "" {
		return EINVAL
	}
	for i := 0; i < len(path); i++ {
		if path[i] == '\x00' {
			return EINVAL
		}
	}
	return nil
}
