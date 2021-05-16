package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	content := []byte("WRITE a hello\nREAD a\nSTART\nWRITE a hello-again\nREAD a\nABORT\nREAD a\nQUIT\n")
	var stdin bytes.Buffer
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	stdin.Write(content)
	err := run(nil, &stdin, &stdout, &stderr)
	assert.Nil(t, err)
	assert.Equal(t, "hello\nhello-again\nhello\n", stdout.String())
}
func TestRunUnknownCommand(t *testing.T) {
	content := []byte("EXECUTE Order 66\nQUIT\n")
	var stdin bytes.Buffer
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	stdin.Write(content)
	err := run(nil, &stdin, &stdout, &stderr)
	assert.Nil(t, err)
	assert.Equal(t, "Unknown command: EXECUTE\n", stderr.String())
}
func TestRunKeyNotFound(t *testing.T) {
	content := []byte("READ b\nQUIT\n")
	var stdin bytes.Buffer
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	stdin.Write(content)
	err := run(nil, &stdin, &stdout, &stderr)
	assert.Nil(t, err)
	assert.Equal(t, "Key not found: b\n", stderr.String())
}

func TestCheckArgs(t *testing.T) {
	err := checkArgs([]string{"a", "b"}, 1, 1)
	assert.NotNil(t, err)

	err = checkArgs([]string{"a", "b"}, 1, 2)
	assert.Nil(t, err)

	err = checkArgs([]string{"a"}, 1, 1)
	assert.Nil(t, err)

	err = checkArgs([]string{"a", "b", "c"}, 1, -1)
	assert.Nil(t, err)
}
