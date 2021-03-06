package cellvalue

import (
	"fmt"
	"strconv"
)

type valueStoreInt64 struct {
	*valueStore

	value int
}

func (vs *valueStoreInt64) Store(data string, vt ValueType) error {
	i64, err := strconv.ParseInt(data, 10, 0)
	if err != nil {
		return fmt.Errorf("ValueStore[%v] Invalid int data[%v]", vs.GoType(), data)
	}

	vs.value = int(i64)
	vs.valueStore.value = vs.value
	return nil
}

func init() {
	basicValueStoreCreator[vtInt64] = func(vt ValueType) ValueStore {
		return &valueStoreInt64{
			valueStore: &valueStore{
				vt: vt,
			},
		}
	}
}
