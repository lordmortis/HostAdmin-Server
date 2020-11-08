package datasource

type Login struct {
	Username string
	Password string
}

func (login *Login) Validate() map[string]interface{} {
	errorMap := make(map[string]interface{})

	if len(login.Username) == 0 {
		errorMap["username"] = []string{"must be present"}
	}

	if len(login.Password) == 0 {
		errorMap["password"] = []string{"must be present"}
	}

	return errorMap
}
