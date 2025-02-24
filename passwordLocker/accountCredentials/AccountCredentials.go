package accountcredentials

import (
	"fmt"
)

type Accountcredentials struct {

	AccountName string

	AccountUserName string

	AccountPassword string
}

var AccountcredentialsDb map[string]Accountcredentials

func init() {
	AccountcredentialsDb = make(map[string]Accountcredentials)
}

func AddUserCredentials(name string, username string, password string) (Accountcredentials, error) {

	if _, exists := AccountcredentialsDb[name]; exists {
		return Accountcredentials{ },
			fmt.Errorf(" Account Credential with this name already exists")
	}

	accountCred := Accountcredentials {name, username, password}
	if AccountcredentialsDb == nil {
		AccountcredentialsDb = make(map[string]Accountcredentials)
	}
	AccountcredentialsDb[name] = accountCred
	return accountCred, nil
}

func GetAccount(key string) (Accountcredentials, error) {
	if account, exists := AccountcredentialsDb[key]; exists {
		return account, nil
	}

	return Accountcredentials{},
		fmt.Errorf(" Account Not found")
}


func (acc *Accountcredentials) EditAccountCredentials(name string, key string,
	value string) (Accountcredentials, error) {

	if account, exists := AccountcredentialsDb[name]; exists {

		switch key {
			case "AccountName":
				account.AccountName = value
			case "AccountPassword":
				account.AccountPassword = value
			default:
				return Accountcredentials{},
					fmt.Errorf("Key field to change not found !!!")
		}
		return account, nil
	}
	return Accountcredentials{},
		fmt.Errorf("Acount Not Found !!!")
}

func ViewCredentials() {

	fmt.Println("| Account Name | Account UserName | Account Passsword")
	for _, account := range AccountcredentialsDb {

		fmt.Println("| " + account.AccountName +
			" | " + account.AccountUserName +
			" | " + account.AccountPassword + "")
	}
}