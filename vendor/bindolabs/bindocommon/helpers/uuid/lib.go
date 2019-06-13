package uuid

import (
	//"bindolabs/bindocommon/env"
	"context"
	"strconv"

	"github.com/smallnest/rpcx/core"
	
	service2 "bindolabs/bindocommon/helpers/service"
	"bindolabs/golib/id"
)

type UUID uint64

var IDGenr *id.SnowFlake

func init() {
	IDGenr, _ = id.NewSnowFlake(0, 0)
}

func uuidFromService() UUID {
	service, err := service2.NewClient("uuid")
	if err != nil {
		panic(err)
	}
	var reply int64
	for i := 0; i < 4; i++ {
		if err := service.Call(core.NewContext(context.Background(), core.NewHeader()), "New", 1, &reply); err != nil {
			if err == core.ErrShutdown {
				if i == 3 {
					panic(err)
				}
				continue
			}
			panic(err)
		} else {
			break
		}
	}

	if reply < 1 {
		panic("Get UUID failed")
	}
	return UUID(reply)
}

func uuidFromLocal() UUID {
	i, _ := IDGenr.Next()
	return UUID(i)
}

func New() UUID {
	//if env.IsTest() || env.IsDev() {
	//	return uuidFromLocal()
	//}
	return uuidFromService()
}

func (c UUID) String() string {
	return strconv.FormatUint(uint64(c), 10)
	// return string(c)
}

//
//func (c UUID) Value() (driver.Value, error) {
//	return int64(c), nil
//}
//
//func (c *UUID) Scan(input interface{}) error {
//	switch value := input.(type) {
//	case string:
//		a, err := strconv.ParseUint(value, 10, 64)
//		if err != nil {
//			return err
//		}
//		tmp := UUID(a)
//		c = &tmp
//		return nil
//	case []byte:
//		a, err := strconv.ParseUint(string(value), 10, 64)
//		if err != nil {
//			return err
//		}
//		tmp := UUID(a)
//		c = &tmp
//		return nil
//	case uint, uint8, uint16, uint32, uint64:
//		if i, ok := value.(uint64); ok {
//			tmp := UUID(i)
//			c = &tmp
//			return nil
//		}
//		return fmt.Errorf("not support type:%T", value)
//	case int, int64, int8, int16, int32:
//		if i, ok := value.(int64); ok {
//			tmp := UUID(i)
//			c = &tmp
//			return nil
//		}
//		return fmt.Errorf("not support type:%T", value)
//	default:
//		return fmt.Errorf("%T is not a uint64", value)
//	}
//}
