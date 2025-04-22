package auth

func GetSession() string {
	return extractSession()
}
func extractSession() string{
	return "login"
} 