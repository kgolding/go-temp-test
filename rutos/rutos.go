package rutos

import (
	"encoding/json"
	"log"
	"os/exec"
	"strings"
)

type SendSMSPayload struct {
	Number   string `json:"number"`
	Test     string `json:"text"`
	Validate bool   `json:"validate,omitempty"`
	Async    bool   `json:"async,omitempty"`
}

func SendSMS(tel string, text string) error {
	err := ubus_call("send_sms", map[string]any{
		"number": tel,
		"text":   text,
	})
	return err
}

func ubus_call(command string, data map[string]any) error {
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}
	args := []string{"ubus", "call", "gsm.modem0", command, "'" + string(payload) + "'"}
	log.Printf("Execute: %s\n", strings.Join(args, " "))
	_, err = exec.Command(args[0], args[1:]...).CombinedOutput()
	if err != nil {
		return err
	}
	// @todo parse output

	return nil
}
