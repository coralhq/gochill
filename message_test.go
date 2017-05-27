package gochill

import (
	"errors"
	"os"
	"testing"
)

func TestMsgEmptyFullMsg(t *testing.T) {
	message := Msg("test")

	if message.Short != "test" {
		t.Fail()
		t.Log("Expected result short variable is equal with test")
		t.Log("Given result : ", message.Short)
	}

	if message.Full != "" {
		t.Fail()
		t.Log("Expected result full variable is empty string")
		t.Log("Given result : ", message.Full)
	}
}

func TestMsgFullMsgNotEmpty(t *testing.T) {
	message := Msg("test short", "test full")

	if message.Short != "test short" {
		t.Fail()
		t.Log("Expected result short variable is equal with test sort")
		t.Log("Given result : ", message.Short)
	}

	if message.Full != "test full" {
		t.Fail()
		t.Log("Expected result full variable is equal with test full")
		t.Log("Given result : ", message.Full)
	}

}

func TestNewMessage(t *testing.T) {
	os.Setenv(EnvServiceKeyName, "testing_gochill_service")

	m := NewMessage(LevelAlert)
	if m.Level != LevelAlert {
		t.Fail()
		t.Log("Unexpected Level : ", m.Level)
	}

	if m.Host == "" {
		t.Fail()
		t.Log("Host should not be empty.")
	}

	if m.Service != os.Getenv(EnvServiceKeyName) {
		t.Fail()
		t.Log("Unexpected Service Name : ", m.Service)
	}
}

func TestMessageHostnameError(t *testing.T) {
	fakeHostname := func() (string, error) {
		return "", errors.New("error")
	}

	m := NewMessage(LevelAlert)
	m.ReplaceOSEngine(fakeHostname)

	if m.Host != "" {
		t.Fail()
		t.Log("Expected hostname : nil")
		t.Log("Given hostname value : ", m.Host)
	}
}

func TestParseBool(t *testing.T) {

	ok := getOutput("true")
	if !ok {
		t.Fail()
		t.Log("Expected status : true")
		t.Log("Given status : ", ok)
	}

	no := getOutput("false")
	if no {
		t.Fail()
		t.Log("Expected status : false")
		t.Log("Given status : ", no)
	}

	auto := getOutput("auto")
	if !auto {
		t.Fail()
		t.Log("Expected status : true")
		t.Log("Given status : ", auto)
	}
}
