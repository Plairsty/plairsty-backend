package utils

import (
	"regexp"
)

func EmailValidator(email string) bool {
	const emailRegx = `^([a-zA-Z0-9_\-\.]+)@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.)|(([a-zA-Z0-9\-]+\.)+))([a-zA-Z]{2,4}|[0-9]{1,3})(\]?)$`
	return regexp.MustCompile(emailRegx).MatchString(email)
}
