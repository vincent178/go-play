package main

import "syscall"

func main() {

}

func isNotExist(err error) bool {
	return err == syscall.ENOENT
}
