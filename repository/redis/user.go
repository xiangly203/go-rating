package redis

func LoginCheck(phone string, code string) bool {
	val, err := GetRedisVal(phone)
	if err != nil || val != code {
		return false
	}
	return true
}

func GetCode() {

}
