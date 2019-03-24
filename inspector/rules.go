package inspector

func rules() map[string]string {
	r := make(map[string]string)
	r["font-weight:bold"] = "text-bold"
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
