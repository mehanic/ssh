package territory

func removeByUsername(s []User, n string) []User {
	var index int
	for i, u := range s {
		if u.Session.User() == n {
			index = i
			break
		}
	}
	return append(s[:index], s[index+1:]...)
}

func send(u User, m Message) {
	raw := m.From + "> " + m.Message + "\n"
	u.Terminal.Write([]byte(raw))
}
