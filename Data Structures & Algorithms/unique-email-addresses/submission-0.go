func numUniqueEmails(emails []string) int {
    result := 0
    existingEmails := make(map[string]bool)

    for _, email := range emails {
        name, domain, _ := strings.Cut(email, "@")
        if !isValidDomain(domain) {
            continue
        }
        sanitizedName := sanitizeName(name)
        sanitizedEmail := sanitizedName + "@" + domain

        if _, exists := existingEmails[sanitizedEmail]; !exists {
            existingEmails[sanitizedEmail] = true
            result++
        }
    }

    return result
}

func isValidDomain(domain string) bool {
    return len(domain) > 4
}

func sanitizeName(name string) string {
    sanitizedName, _, _ := strings.Cut(name, "+")
    
    return strings.ReplaceAll(sanitizedName, ".", "")
}
