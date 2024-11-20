package test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/farseer-go/collections"
	"github.com/stretchr/testify/assert"
)

func Test_dictionary_Values(t *testing.T) {
	maps := make(map[string]string)
	maps["name"] = "steden"
	maps["age"] = "18"
	dic := collections.NewDictionaryFromMap[string, string](maps)
	array := dic.Values().ToArray()
	assert.True(t, true, array[0] == "steden" || array[0] == "18")
	assert.True(t, true, array[1] == "steden" || array[0] == "18")

	assert.Equal(t, "steden", dic.ToDictionary().GetValue("name"))
	assert.Equal(t, "18", dic.ToDictionary().GetValue("age"))
}

func Test_dictionary_Keys(t *testing.T) {
	maps := make(map[string]string)
	maps["name"] = "steden"
	maps["age"] = "18"
	dic := collections.NewDictionaryFromMap[string, string](maps)
	lst := dic.Keys()
	assert.Equal(t, 2, lst.Count())
	assert.True(t, lst.Contains("name"))
	assert.True(t, lst.Contains("age"))
}

func Test_dictionary_Count(t *testing.T) {

	maps := make(map[string]string)
	maps["name"] = "steden"
	maps["age"] = "18"
	dic := collections.NewDictionaryFromMap[string, string](maps)
	assert.Equal(t, dic.Count(), 2)
}

func Test_dictionary_Add(t *testing.T) {
	maps := make(map[string]string)
	maps["name"] = "steden"
	maps["age"] = "18"
	dic := collections.NewDictionaryFromMap[string, string](maps)
	assert.Equal(t, 2, dic.Count())
	dic.Add("sex", "boy")
	assert.Equal(t, 3, dic.Count())

	assert.Equal(t, "steden", dic.GetValue("name"))
	assert.Equal(t, "18", dic.GetValue("age"))
	assert.Equal(t, "boy", dic.GetValue("sex"))
}

func Test_dictionary_Update(t *testing.T) {
	maps := make(map[string]string)
	maps["name"] = "steden"
	maps["age"] = "18"
	dic := collections.NewDictionaryFromMap[string, string](maps)
	dic.Update("name", func(value *string) {
		*value = "tao"
	})
	assert.Equal(t, "tao", dic.GetValue("name"))
}

func TestDictionary_AddMap(t *testing.T) {
	dic := collections.NewDictionary[string, int]()
	maps := map[string]int{"age": 18}
	dic.AddMap(maps)
	assert.Equal(t, 1, dic.Count())
	assert.Equal(t, 18, dic.GetValue("age"))
}

func Test_dictionary_Clear(t *testing.T) {
	maps := make(map[string]string)
	maps["name"] = "steden"
	maps["age"] = "18"
	dic := collections.NewDictionaryFromMap[string, string](maps)
	assert.Equal(t, 2, dic.Count())
	dic.Clear()
	assert.Equal(t, 0, dic.Count())
}

func Test_dictionary_Remove(t *testing.T) {
	maps := make(map[string]string)
	maps["name"] = "steden"
	maps["age"] = "18"
	dic := collections.NewDictionaryFromMap(maps)
	assert.Equal(t, "steden", dic.GetValue("name"))
	dic.Remove("name")
	assert.Equal(t, "", dic.GetValue("name"))
}

func Test_dictionary_ContainsKey(t *testing.T) {
	dic := collections.NewDictionary[string, string]()
	assert.Equal(t, false, dic.ContainsKey("name"))
	dic.Add("name", "steden")
	assert.Equal(t, true, dic.ContainsKey("name"))
}

func Test_dictionary_ContainsValue(t *testing.T) {
	dic := collections.NewDictionary[string, string]()
	assert.Equal(t, false, dic.ContainsValue("steden"))
	dic.Add("name", "steden")
	assert.Equal(t, true, dic.ContainsValue("steden"))
}

func TestDictionary_ToMap(t *testing.T) {
	maps := map[string]string{"name": "steden", "age": "18"}
	dic := collections.NewDictionaryFromMap[string, string](maps)
	tomap := dic.ToMap()

	assert.Equal(t, len(maps), len(tomap))
	assert.Equal(t, maps["name"], tomap["name"])
	assert.Equal(t, maps["age"], tomap["age"])
}

func TestDictionary_ToReadonlyDictionary(t *testing.T) {
	maps := map[string]string{"name": "steden", "age": "18"}
	dic := collections.NewDictionaryFromMap[string, string](maps)
	readonly := dic.ToReadonlyDictionary()
	tomap := readonly.ToMap()
	assert.Equal(t, len(maps), len(tomap))
	assert.Equal(t, maps["name"], tomap["name"])
	assert.Equal(t, maps["age"], tomap["age"])
}

func TestDictionary_Value(t *testing.T) {
	maps := make(map[string]string)
	maps["name"] = "steden"
	maps["age"] = "18"
	dic := collections.NewDictionaryFromMap[string, string](maps)
	value, err := dic.Value()
	assert.Equal(t, value, "{\"age\":\"18\",\"name\":\"steden\"}")
	assert.Equal(t, err, nil)
	dic = collections.NewDictionaryFromMap[string, string](nil)
	value, err = dic.Value()
	assert.Equal(t, "{}", value)
	assert.Equal(t, nil, err)
}

func TestDictionary_Scan(t *testing.T) {
	var val any
	val = nil
	dic := collections.NewDictionary[string, string]()
	rtn := dic.Scan(val)
	assert.Equal(t, rtn, nil)
	val = 123
	rtn = dic.Scan(val)
	assert.Equal(t, rtn, errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", 123)))
	val = []byte(`{
      "name":"sam",
      "birthday":"1995-06-03"
   }`)
	rtn = dic.Scan(val)
	assert.Equal(t, rtn, nil)
	val = `{
      "name":"sam",
      "birthday":"1995-06-03"
   }`
	rtn = dic.Scan(val)
	assert.Equal(t, rtn, nil)
}

func TestDictionary_MarshalJSON(t *testing.T) {
	dic := collections.NewDictionary[string, string]()
	dic.Add("a", "1")
	json, err := dic.MarshalJSON()
	assert.Equal(t, []byte("{\"a\":\"1\"}"), json)
	assert.Equal(t, err, nil)
}
func TestDictionary_UnmarshalJSON(t *testing.T) {
	val := []byte(`{
      "name":"sam",
      "birthday":"1995-06-03"
   }`)
	dic := collections.NewDictionary[string, string]()
	err := dic.UnmarshalJSON(val)
	maps := dic.ToMap()
	assert.Equal(t, err, nil)
	assert.Equal(t, maps["name"], "sam")
	assert.Equal(t, maps["birthday"], "1995-06-03")
}

func TestDictionary_GormDataType(t *testing.T) {
	dic := collections.NewDictionary[string, string]()
	val := dic.GormDataType()
	assert.Equal(t, "JSON", val)
}

func TestDictionary_IsNil(t *testing.T) {
	dic := collections.NewDictionaryFromMap[string, string](nil)
	val := dic.IsNil()
	assert.Equal(t, false, val)
	maps := make(map[string]string)
	maps["name"] = "steden"
	maps["age"] = "18"
	dic2 := collections.NewDictionaryFromMap[string, string](maps)
	val = dic2.IsNil()
	assert.Equal(t, val, false)

	var dicNil collections.Dictionary[string, string]
	assert.Equal(t, 0, dicNil.Count())

}

func TestDictionaryJson(t *testing.T) {
	var dic collections.Dictionary[int, int]
	assert.True(t, dic.IsNil())
	marshal, _ := sonic.Marshal(dic)
	assert.Equal(t, "{}", string(marshal))

	_ = sonic.Unmarshal([]byte("{}"), &dic)
	assert.False(t, dic.IsNil())
}

func TestReadonlyDictionary(t *testing.T) {
	var dic collections.Dictionary[int, int]
	value, _ := dic.Value()
	assert.Equal(t, nil, value)
}
