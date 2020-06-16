package html

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_AttributeSizing(t *testing.T) {
	kv := &KeyVal{
		Key:   "key",
		Value: "value",
	}

	size := kv.GetRenderSize()
	assert.Equal(t, 11, size)
}

func Test_ElementSizing(t *testing.T) {
	e := &Element{}
	e.AddAttribute("key", "value")

	actual := e.GetRenderSize()
	buf := new(bytes.Buffer)
	_, err := e.RenderAttr(buf)
	assert.NoError(t, err)
	assert.Equal(t, actual, buf.Len())

	buf.Reset()
	e.AddAttribute("anotherkey", "another value")
	actual = e.GetRenderSize()
	_, err = e.RenderAttr(buf)
	assert.NoError(t, err)
	assert.Equal(t, actual, buf.Len())

	buf.Reset()
	e.AddAttribute("thirdkey", "value")
	actual = e.GetRenderSize()
	_, err = e.RenderAttr(buf)
	assert.NoError(t, err)
	assert.Equal(t, actual, buf.Len())
}
