package utils

func FormatError(rawCmd string, msg string) string {
	return rawCmd + " command failed: " + msg
}
