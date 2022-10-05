package logrtlim

type Logger struct {
	msgsToAllowedTS map[string]int
}

func Constructor() Logger {
	m := make(map[string]int)
	return Logger{m}
}

func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	ts, ok := this.msgsToAllowedTS[message]
	if !ok || timestamp >= ts {
		this.msgsToAllowedTS[message] = timestamp + 10
		return true
	}
	return false
}
