import (
	"strconv"
	"strings"
)

func haveConflict(event1 []string, event2 []string) bool {
	return !compareTime(event1[1], event2[0]) && !compareTime(event2[1], event1[0])
}

// t1 < t2 true
func compareTime(t1, t2 string) bool {
	t1d := strings.Split(t1, ":")
	t2d := strings.Split(t2, ":")

	th1, _ := strconv.Atoi(t1d[0])
	ts1, _ := strconv.Atoi(t1d[1])

	th2, _ := strconv.Atoi(t2d[0])
	ts2, _ := strconv.Atoi(t2d[1])
	if th1 < th2 {
		return true
	}

	if th1 == th2 && ts1 < ts2 {
		return true
	}
	return false
}