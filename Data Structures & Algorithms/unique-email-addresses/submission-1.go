func numUniqueEmails(emails []string) int {
    var result int
	uniqueEmails := make(map[string]bool)

	for _, email := range emails {
		parts := strings.Split(email, "@")
		localName, domain := parts[0], parts[1]
		sanitizedLocalName := strings.ReplaceAll(localName, ".", "")
		suffixSplit := strings.Split(sanitizedLocalName, "+")
		uniqueEmail := sanitizedLocalName

		if len(suffixSplit) > 0 {
			uniqueEmail = suffixSplit[0]
		}
		uniqueEmail += "@" + domain
		if uniqueEmails[uniqueEmail] {
			continue
		}
		uniqueEmails[uniqueEmail] = true
		result++
	}

	return result
}