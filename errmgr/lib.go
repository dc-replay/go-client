package errmgr

import "log"

func Report(err error) bool {
	if err != nil {
		log.Printf(err.Error())
		return true
	}
	return false
}
