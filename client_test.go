package gocent

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	c := New(Config{})
	if c == nil {
		t.Errorf("New returned nil client")
	}
}

func TestClient_GenerateHashSign(t *testing.T) {
	c := New(Config{Secret: "8942b28b-f429-4306-91ac-dfb91479f191"})
	hashStr := c.GenerateHashSign("6164c0d9-5601-43ef-832b-bbed6ead01d7", "$514149", "")

	if hashStr != "75f86bd0c860e416489b7f8664305903c976e9aaf49e790606f4f571fba4813c" {
		t.Fatal("Hash is not equals")
	}
}
