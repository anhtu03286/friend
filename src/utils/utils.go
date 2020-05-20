package utils

import (
	"regexp"
)

const EMAIL_REGEX = "[_A-Za-z0-9-\\+]+(\\.[_A-Za-z0-9-]+)*@[A-Za-z0-9-]+(\\.[A-Za-z0-9]+)*(\\.[A-Za-z]{2,})"

func IsFormatEmail(email string) bool {
	re, _ := regexp.Compile(EMAIL_REGEX)
	if re.MatchString(email) {
		return true
	}
	return false
}

func GetEmailFromText(text string) []string {
	keys := make(map[string]bool)
	setEmails := []string{}

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
	newList := []int64{}
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