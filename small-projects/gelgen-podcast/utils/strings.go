package utils

// CheckNullString 用來檢查 string 是否為空字串
func CheckNullString(str string) *string {
	if len(str) == 0 {
		return nil
	}
	return &str
}
