package assert

import (
	"fmt"
)

func buildMessage(fixedMessage string, expected, actual interface{}, msgAndArgs ...interface{}) string {
	message := buildMessageFromMsgAndArgs(msgAndArgs...)
	return fmt.Sprintf("%+v: %+v\n"+
		"Expected: %+v\n"+
		"     got: %+v\n", fixedMessage, message, expected, actual)
}

func buildMessageFromMsgAndArgs(msgAndArgs ...interface{}) string {
	// empty message.
	if len(msgAndArgs) == 0 || msgAndArgs == nil {
		return ""
	}
	// message only.
	if len(msgAndArgs) == 1 {
		msg := msgAndArgs[0]
		if msgAsStr, ok := msg.(string); ok {
			return msgAsStr
		}
		return fmt.Sprintf("%+v", msg)
	}
	// message and arguments.
	msg := msgAndArgs[0]
	if msgAsStr, ok := msg.(string); ok {
		return fmt.Sprintf(msgAsStr, msgAndArgs[1:]...)
	}

	return fmt.Sprintf("%+v", msg)
}
