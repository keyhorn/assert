package assert

import "testing"

func TestBuildMessage0(t *testing.T) {
	fixedMessage := ""
	expected := ""
	actual := ""
	message0 := ""
	expectedMessage := ": \n" +
		"Expected: \n" +
		"     got: \n"
	actualMessage := buildMessage(fixedMessage, expected, actual, message0)
	if expectedMessage != actualMessage {
		t.Errorf("Expected \"%+v\" but \"%+v\"", expectedMessage, actualMessage)
	}
}

func TestBuildMessage1(t *testing.T) {
	fixedMessage := "This is a message"
	expected := "a"
	actual := "b"
	message0 := "This is additional message"
	expectedMessage := "This is a message: This is additional message\n" +
		"Expected: a\n" +
		"     got: b\n"
	actualMessage := buildMessage(fixedMessage, expected, actual, message0)
	if expectedMessage != actualMessage {
		t.Errorf("Expected \"%+v\" but \"%+v\"", expectedMessage, actualMessage)
	}
}

func TestBuildMessage2(t *testing.T) {
	fixedMessage := "This is a message"
	expected := "a"
	actual := "b"
	message := "This is additional message %+v"
	param1 := "param1"
	expectedMessage := "This is a message: This is additional message param1\n" +
		"Expected: a\n" +
		"     got: b\n"
	actualMessage := buildMessage(fixedMessage, expected, actual, message, param1)
	if expectedMessage != actualMessage {
		t.Errorf("Expected \"%+v\" but \"%+v\"", expectedMessage, actualMessage)
	}
}

func TestBuildMessage3(t *testing.T) {
	fixedMessage := "This is a message"
	expected := "a"
	actual := "b"
	message := "This is additional message %+v %+v"
	param1 := "param1"
	param2 := "param2"
	expectedMessage := "This is a message: This is additional message param1 param2\n" +
		"Expected: a\n" +
		"     got: b\n"
	actualMessage := buildMessage(fixedMessage, expected, actual, message, param1, param2)
	if expectedMessage != actualMessage {
		t.Errorf("Expected \"%+v\" but \"%+v\"", expectedMessage, actualMessage)
	}
}

func TestBuildMessageFromMsgAndArgs0(t *testing.T) {
	expected := ""
	actual := buildMessageFromMsgAndArgs("")
	if expected != actual {
		t.Errorf("Expected \"%+v\" but \"%+v\"", expected, actual)
	}
}

func TestBuildMessageFromMsgAndArgs1(t *testing.T) {
	expected := "This is a message"
	actual := buildMessageFromMsgAndArgs("This is a message")
	if expected != actual {
		t.Errorf("Expected \"%+v\" but \"%+v\"", expected, actual)
	}
}

func TestBuildMessageFromMsgAndArgs2(t *testing.T) {
	expected := "This is a message: param1"
	actual := buildMessageFromMsgAndArgs("This is a message: %+v", "param1")
	if expected != actual {
		t.Errorf("Expected \"%+v\" but \"%+v\"", expected, actual)
	}
}

func TestBuildMessageFromMsgAndArgs3(t *testing.T) {
	expected := "This is a message: param1 param2"
	actual := buildMessageFromMsgAndArgs("This is a message: %+v %+v", "param1", "param2")
	if expected != actual {
		t.Errorf("Expected \"%+v\" but \"%+v\"", expected, actual)
	}
}
