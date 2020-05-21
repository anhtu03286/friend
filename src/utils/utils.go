package utils

import (
	"regexp"
)

const EMAIL_REGEX = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

func IsFormatEmail(email string) bool {
	re := regexp.MustCompile(EMAIL_REGEX)
	if re.MatchString(email) {
		return true
	}
	return false
}

func GetEmailFromText(text string) []string {
	keys := make(map[string]bool)
	var setEmails []string

	re := regexp.MustCompile(EMAIL_REGEX)
	submatchall := re.FindAllString(text, -1)

	for _, element := range submatchall {
		if _, ok := keys[element]; !ok {
			keys[element] = true
			setEmails = append(setEmails, element)
		}
	}
	return setEmails
}

func RemoveItemFromList(list []int64, item int64) []int64 {
	var newList []int64
	for _, i := range list {
		if i != item {
			newList = append(newList, i)
		}
	}
	return newList
}

func Contains(list []int64, item int64) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}
	return false
}