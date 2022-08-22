package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_dictionary_Values(t *testing.T) {
	maps := make(map[string]string)
	maps["name"] = "steden"
	maps["age"] = "18"
	dic := NewDictionaryFromMap[string, string](maps)
	array := dic.Values().ToArray()
	assert.Equal(t, array[0], "steden")
	assert.Equal(t, array[1], "18")
}

func Test_dictionary_Keys(t *testing.T) {
	maps := make(map[string]string)
	maps["name"] = "steden"
	maps["age"] = "18"
	dic := NewDictionaryFromMap[string, string](maps)
	lst := dic.Keys()
	assert.Equal(t, 2, lst.Count())
	assert.True(t, lst.Contains("name"))
	assert.True(t, lst.Contains("age"))
}

func Test_dictionary_Count(t *testing.T) {

	maps := make(map[string]string)
	maps["name"] = "steden"
	maps["age"] = "18"
	dic := NewDictionaryFromMap[string, string](maps)
	assert.Equal(t, dic.Count(), 2)
}

func Test_dictionary_Add(t *testing.T) {
	maps := make(map[string]string)
	maps["name"] = "steden"
	maps["age"] = "18"
	dic := NewDictionaryFromMap[string, string](maps)
	assert.Equal(t, dic.Count(), 2)
	dic.Add("sex", "boy")
	assert.Equal(t, dic.Count(), 3)

	assert.Equal(t, dic.GetValue("name"), "steden")
	assert.Equal(t, dic.GetValue("age"), "18")
	assert.Equal(t, dic.GetValue("sex"), "boy")
}

func Test_dictionary_Clear(t *testing.T) {
	maps := make(map[string]string)
	maps["name"] = "steden"
	maps["age"] = "18"
	dic := NewDictionaryFromMap[string, string](maps)
	assert.Equal(t, 2, dic.Count())
	dic.Clear()
	assert.Equal(t, 0, dic.Count())
}

func Test_dictionary_Remove(t *testing.T) {
	maps := make(map[string]string)
	maps["name"] = "steden"
	maps["age"] = "18"
	dic := NewDictionaryFromMap(maps)
	assert.Equal(t, "steden", dic.GetValue("name"))
	dic.Remove("name")
	assert.Equal(t, "", dic.GetValue("name"))
}

func Test_dictionary_ContainsKey(t *testing.T) {
	dic := NewDictionary[string, string]()
	assert.Equal(t, false, dic.ContainsKey("name"))
	dic.Add("name", "steden")
	assert.Equal(t, true, dic.ContainsKey("name"))
}

func Test_dictionary_ContainsValue(t *testing.T) {
	dic := NewDictionary[string, string]()
	assert.Equal(t, false, dic.ContainsValue("steden"))
	dic.Add("name", "steden")
	assert.Equal(t, true, dic.ContainsValue("steden"))
}

func TestDictionary_ToMap(t *testing.T) {
	maps := make(map[string]string)
	maps["name"] = "steden"
	maps["age"] = "18"
	dic := NewDictionaryFromMap[string, string](maps)
	tomap := dic.ToMap()

	assert.Equal(t, len(maps), len(tomap))
	assert.Equal(t, maps["name"], tomap["name"])
	assert.Equal(t, maps["age"], tomap["age"])
}
