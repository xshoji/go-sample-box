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

	user := &entity.User{
		Id:               1,
		Name:             `a`,
		NicknameOptional: &entity.User_Nickname{Nickname: "a"},
		AgeOptional:      &entity.User_Age{Age: 16},
		Birth: &entity.Birth{
			Datetime:         getTimestamp(),
			WeightOptional:   &entity.Birth_Weight{Weight: 12},
			HospitalOptional: &entity.Birth_Hospital{Hospital: "hugeHospital"},
		},
		Addresss: &entity.Address{
			Country:         entity.Country_JP,
			ZipCodeOptional: &entity.Address_ZipCode{ZipCode: 1},
			StateOptional:   &entity.Address_State{State: "test"},
			CityOptional:    &entity.Address_City{City: "city"},
		},
	}

	m := jsonpb.Marshaler{
		EmitDefaults: true,
		Indent:       "    ",
		OrigName:     true,
	}
	m.Marshal(os.Stdout, user)

}
