package main

func Auth(login string, password string) (bool, error) {
	if login == "11" && password == "22" {
		return true, nil
	}

	return false, nil
}
