package main

func DebugLog(args ...string) []string {
	another := []string{"[DEBUG]"}
	args = append(another, args...)
	return args

}

func InfoLog(args ...string) []string {
	another := []string{"[INFO]"}
	args = append(another, args...)
	return args
}

func ErrorLog(args ...string) []string {
	another := []string{"[ERROR]"}
	args = append(another, args...)
	return args
}
