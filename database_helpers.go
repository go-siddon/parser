package parser

import "slices"

func checkTag(fieldTags []string, tag string) bool {
	return slices.Contains(fieldTags, tag)
}

func getFieldname(fieldTags []string) string {
	for _, tags := range fieldTags {
		switch tags {
		case propertyIndex, propertyRequired, propertyUnique:
			continue
		default:
			return tags
		}
	}
	return ""
}
