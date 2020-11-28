package cmd

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainCommandIsCorrectType(t *testing.T) {
	assert.ObjectsAreEqual("*cobra.Command", reflect.TypeOf(MainCommand))
}
