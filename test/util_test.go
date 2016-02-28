package nope_test

func formatError(rawCmd string, msg string) string {
	return rawCmd + " command failed: " + msg
}
