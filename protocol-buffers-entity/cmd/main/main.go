package main

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/xshoji/go-sample-box/protocol-buffers-entity/entity/proto"
	"os"
	"time"
)

func main() {

	getTimestamp := func() *timestamp.Timestamp {
		t, _ := ptypes.TimestampProto(time.Now())
		return t
	}

	// どうやらnullはjson化した際設定されない思想らしい。
	// > why protobuf optional field does not take null - Google グループ
	// > https://groups.google.com/forum/#!topic/protobuf/KcLoxlJGVMY
	// Protocol buffers has no concept of null.  Fields cannot be set to null.  You can *clear* a field, like:
	user := &entity.User{
		Id:               1,
		Name:             `a`,
		NicknameOptional: &entity.User_Nickname{Nickname: "a"},
		AgeOptional:      &entity.User_Age{Age: 16},
		Birth: &entity.Birth{
			Datetime:         getTimestamp(),
			WeightOptional:   &entity.Birth_Weight{Weight: 12},
			HospitalOptional: nil,
		},
		Addresss: &entity.Address{
			Country:         entity.Country_JP,
			ZipCodeOptional: &entity.Address_ZipCode{ZipCode: 1},
			StateOptional:   &entity.Address_State{State: "test"},
			CityOptional:    nil,
		},
	}

	m := jsonpb.Marshaler{
		EmitDefaults: true,
		Indent:       "  ",
		OrigName:     true,
	}
	m.Marshal(os.Stdout, user)

}
