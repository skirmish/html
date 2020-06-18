package html

/*func Test_AttributeSizing(t *testing.T) {
	kv := &keyVal{
		Key:   "key",
		Value: "value",
	}

	size := kv.getRenderSize()
	assert.Equal(t, 11, size)
}*/

/*func Test_ElementSizing(t *testing.T) {
	e := &Element{}
	e.AddAttribute("key", "value")

	actual := e.getRenderSize()
	buf := new(bytes.Buffer)
	_, err := e.RenderAttr(buf)
	assert.NoError(t, err)
	assert.Equal(t, actual, buf.Len())

	buf.Reset()
	e.AddAttribute("anotherkey", "another value")
	actual = e.getRenderSize()
	_, err = e.RenderAttr(buf)
	assert.NoError(t, err)
	assert.Equal(t, actual, buf.Len())

	buf.Reset()
	e.AddAttribute("thirdkey", "value")
	actual = e.getRenderSize()
	_, err = e.RenderAttr(buf)
	assert.NoError(t, err)
	assert.Equal(t, actual, buf.Len())
}
*/
