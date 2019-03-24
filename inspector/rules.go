package inspector

func rules() map[string]string {
	// Create the rules Map
	r := make(map[string]string)

	// Borders
	r["border: 1px solid #dee2e6"] = "border"
	r["border-top: 1px solid #dee2e6"] = "border-top"
	r["border-bottom: 1px solid #dee2e6"] = "border-bottom"
	r["border-right: 1px solid #dee2e6"] = "border-right"
	r["border-left: 1px solid #dee2e6"] = "border-left"

	// Border radius
	r["border-radius: .25rem"] = "border"

	// Typography
	r["font-weight: bold"] = "text-bold"
	return r
}

func SearchRule(rule string) string {
	for k, v := range rules() {
		if rule == k {
			return v
		}
	}
	return ""
}
