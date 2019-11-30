package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"gopkg.in/yaml.v2"
	"strings"
)

// option details: https://github.com/golang/glog/blob/master/glog.go#L38-L70
func main() {

	flag.Parse()

	stupidHonestMethod()
	//	littleSmartMethod()
}

func stupidHonestMethod() {
	var yamlString = getYaml()

	// new node
	node := yaml.MapSlice{}
	newNode := yaml.MapSlice{}
	yaml.Unmarshal([]byte(yamlString), &node)
	fmt.Printf("--- input:\n%v\n\n", node)

	glog.Infof(">>> START >>>\n")
	for _, topLevelNode := range node {
		glog.Infof("topLevelNode.Key : %v\n", topLevelNode.Key)
		if topLevelNode.Key != `development-teams` {
			// next
			glog.Infof("    -> added, next\n")
			newNode = append(newNode, topLevelNode)
			continue
		}

		newTeams := make(yaml.MapSlice, 0)
		// cast
		glog.Infof("    topLevelNode.Value : %v\n", topLevelNode.Value)
		teams := topLevelNode.Value.(yaml.MapSlice)
		glog.Infof("    teams : %v\n", teams)
		for _, team := range teams {

			newTeamDetails := make(yaml.MapSlice, 0)
			teamDetails := team.Value.(yaml.MapSlice)
			glog.Infof("        team.Value : %v\n", team.Value)
			for _, teamDetail := range teamDetails {
				if teamDetail.Key != `products` {
					newTeamDetails = append(newTeamDetails, teamDetail)
					continue
				}

				newProducts := make(yaml.MapSlice, 0)
				// cast
				products := teamDetail.Value.(yaml.MapSlice)
				glog.Infof("            teamDetail.Value : %v\n", teamDetail.Value)
				glog.Infof("            products : %v\n", products)
				for _, product := range products {
					// if ios app ...
					if strings.Contains(product.Key.(string), `ios`) {
						// add
						newProducts = append(newProducts, product)
					}
				}
				if len(newProducts) == 0 {
					continue
				}
				// set products
				teamDetail.Value = newProducts
				newTeamDetails = append(newTeamDetails, teamDetail)
			}
			// set teamDetail
			team.Value = newTeamDetails
			newTeams = append(newTeams, team)
		}
		// set team
		topLevelNode.Value = newTeams
		newNode = append(newNode, topLevelNode)
	}
	glog.Infof("<<< END <<<\n")

	bytes, _ := yaml.Marshal(newNode)
	fmt.Printf("\n\n\n--- result:\n%v\n\n", string(bytes))
}

type mapSliceUtil struct{}

var MapSliceUtil = mapSliceUtil{}

func (p *mapSliceUtil) Get(mapSlice yaml.MapSlice, key string) *yaml.MapSlice {
	for _, item := range mapSlice {
		if item.Key == key {
			v := item.Value.(yaml.MapSlice)
			return &v
		}
	}
	return nil
}

func (p *mapSliceUtil) Set(mapSlice yaml.MapSlice, key string, value yaml.MapItem) {
	for i, item := range mapSlice {
		if item.Key == key {
			mapSlice[i].Value = value
		}
	}
}

func (p *mapSliceUtil) Delete(mapSlice *yaml.MapSlice, key string) {
	newMapSlice := yaml.MapSlice{}
	for _, item := range *mapSlice {
		fmt.Printf("%v\n", item.Key)
		if item.Key != key {
			newMapSlice = append(newMapSlice, item)
		}
	}
	fmt.Printf("%v\n", newMapSlice)
	*mapSlice = newMapSlice
}

func (p *mapSliceUtil) MapFunc(mapSlice yaml.MapSlice, mapFunc func(item yaml.MapItem)) {
	for _, item := range mapSlice {
		mapFunc(item)
	}
}

//func (p *mapSliceUtil) MapFuncByKey(mapSlice *yaml.MapSlice, key string, mapFunc func (mapSlice yaml.MapSlice) yaml.MapSlice) {
//	internalMapSlice := MapSliceUtil.Get(*mapSlice, key)
//	mapFunc(internalMapSlice)
//	for _, item := range mapSlice {
//		mapFunc(item)
//	}
//}

//func (p *mapSliceUtil) Delete(mapSlice yaml.MapSlice, isTarget func (item yaml.MapItem) bool) yaml.MapSlice {
//	for i, item := range mapSlice {
//		if isTarget(item) {
//
//		}
//		copy(mapSlice[i:], mapSlice[i+1:])
//		mapSlice[len(mapSlice)-1] = yaml.MapItem{}
//		mapSlice = mapSlice[:len(mapSlice)-1]
//	}
//}

func littleSmartMethod() {
	var yamlString = getYaml()

	node := yaml.MapSlice{}
	yaml.Unmarshal([]byte(yamlString), &node)
	//developmentTeams := MapSliceUtil.Get(node, `development-teams`)
	//MapSliceUtil.MapFunc(developmentTeams, func(developmentTeam yaml.MapItem){
	//	products := MapSliceUtil.Get(developmentTeam.Value.(yaml.MapSlice), `products`)
	//	MapSliceUtil.MapFunc(products, func(product yaml.MapItem){
	//		if strings.Contains(product.Key.(string), `ios`) {
	//			fmt.Printf("--- product:\n%v\n\n", product)
	//		}
	//	})
	//})
	devteamsNode := MapSliceUtil.Get(node, `development-teams`)
	MapSliceUtil.Set(*devteamsNode, `team-a`, yaml.MapItem{})
	MapSliceUtil.Delete(devteamsNode, `team-b`)
	//	*devteamsNode = yaml.MapSlice{}

	bytes, _ := yaml.Marshal(node)
	fmt.Printf("\n\n\n--- result:\n%v\n\n", string(bytes))

}

func getYaml() string {
	return `
# Employee records
employees:
-  martin:
    name: Martin D'vloper
    job: Developer
    skills:
      - python
      - perl
      - pascal
-  tabitha:
    name: Tabitha Bitumen
    job: Developer
    skills:
      - lisp
      - fortran
      - erlang
# Development teams records
development-teams:
  team-a:
    products:
      pc-app-name1:
        id: 1001
        price: 10000
        platform: windows
      ios-app-name2:
        id: 1002
        price: 20000
        platform: ios
      android-app-name3:
        id: 1003
        price: 30000
        platform: android
    members:
      - taro
      - jiro
      - hanako
  team-b:
    products:
      pc-app-name6:
        id: 1006
        price: 10000
        platform: windows
      android-app-name5:
        id: 1005
        price: 20000
        platform: android
      android-app-name4:
        id: 1004
        price: 30000
        platform: android
    members:
      - yamada
      - tanaka
      - sato
  team-c:
    products:
      ios-app-name10:
        id: 1010
        price: 20000
        platform: ios
      ios-app-name9:
        id: 1009
        price: 30000
        platform: ios
    members:
      - john
      - nick
      - bob
`
}
