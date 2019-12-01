package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"gopkg.in/yaml.v2"
	"reflect"
	"strconv"
	"strings"
)

// option details: https://github.com/golang/glog/blob/master/glog.go#L38-L70
func main() {

	flag.Parse()

	//stupidHonestMethod()
	//littleSmartMethod()
	smartMethod()
}

func smartMethod() {
	var yamlString = getSimpleYaml()

	node := &yaml.MapSlice{}
	yaml.Unmarshal([]byte(yamlString), node)
	orderedMapSlice := NewOrderedMapSlice(node)
	dumpNode("result - orderedMapSlice", orderedMapSlice.GetContents())
	fmt.Printf("\n\n\n")
	developmentTeams := orderedMapSlice.Get(`development-teams`)
	dumpNode("result - developmentTeams", developmentTeams.GetContents())
	fmt.Printf("\n\n\n")
	teamA := developmentTeams.Get(`team-a`)
	dumpNode("result - teamA", teamA.GetContents())
	fmt.Printf("\n\n\n")
	id := teamA.Get(`pc-app-name1`).Get(`id`)
	dumpNode("result - id", id.GetContents())
	fmt.Printf("\n\n\n")
	ranks := teamA.Get(`ranks`)
	dumpNode("result - ranks", ranks.GetContents())
	fmt.Printf("\n\n\n")
	ranksFirst := ranks.Get("0")
	dumpNode("result - ranksFirst", ranksFirst.GetContents())
	fmt.Printf("\n\n\n")
	ranksSecond := ranks.Get("1")
	dumpNode("result - ranksSecond", ranksSecond.GetContents())
	fmt.Printf("\n\n\n")
	ranksUndefined := ranks.Get("2")
	dumpNode("result - ranksUndefined", ranksUndefined.GetContents())
	fmt.Printf("\n\n\n")
	undefinedValue := orderedMapSlice.Get(`hogehoge`)
	dumpNode("undefinedValue", undefinedValue.GetContents())
	fmt.Printf("\n\n\n")

	//orderedMapSlice.Set(`development-teams`, nil)
	//fmt.Printf("-- orderedMapSlice set development-teams %v\n", orderedMapSlice.GetContents())
	//dumpNode("result - developmentTeams#set", orderedMapSlice.GetContents())

	//orderedMapSlice.Get(`development-teams`).Set(`team-a`, nil)
	//fmt.Printf("-- orderedMapSlice set development-teams.team-a %v\n", orderedMapSlice.GetContents())
	//dumpNode("result - developmentTeams#set", orderedMapSlice.GetContents())

}

type OrderedMapSlice struct {
	parent      *OrderedMapSlice
	contents    interface{}
	currentItem *yaml.MapItem
	isRoot      bool
}

func NewOrderedMapSlice(yamlData interface{}) OrderedMapSlice {
	rootMapItem := yaml.MapItem{Key: `root`, Value: yamlData}
	rootMapSlice := yaml.MapSlice{rootMapItem}
	orderedMapSlice := OrderedMapSlice{
		parent: &OrderedMapSlice{
			parent:      nil,
			contents:    rootMapSlice,
			currentItem: nil,
			isRoot:      true,
		},
		contents:    yamlData,
		currentItem: &rootMapItem,
		isRoot:      false,
	}
	return orderedMapSlice
}

func createOrderedMapSlice(parent *OrderedMapSlice, item *yaml.MapItem, yamlData interface{}) *OrderedMapSlice {
	return &OrderedMapSlice{
		parent:      parent,
		contents:    yamlData,
		currentItem: item,
		isRoot:      false,
	}
}

func (o *OrderedMapSlice) GetContents() interface{} {
	if o.contents == nil {
		return nil
	}
	v, ok := o.contents.(yaml.MapSlice)
	if ok {
		return v
	}
	return o.contents.(interface{})
}

func (o *OrderedMapSlice) Key() interface{} {
	if o.currentItem != nil {
		return o.currentItem.Key
	}
	return nil
}

func (o *OrderedMapSlice) Value() interface{} {
	if o.currentItem != nil {
		return o.currentItem.Value
	}
	return nil
}

func (o *OrderedMapSlice) Parent() *OrderedMapSlice {
	if o.parent == nil {
		panic("Parent is nul.")
	}
	return o.parent
}

func (o *OrderedMapSlice) Get(key string) *OrderedMapSlice {
	dumpNode("o.GetContents", o.GetContents())
	fmt.Printf(">> o.contents : %T, %v\n", o.contents, o.contents)
	if o.contents == nil {
		return createOrderedMapSlice(o, nil, nil)
	}
	mapSlice, ok := o.contents.(*yaml.MapSlice)
	fmt.Printf("-- o.contents.(yaml.MapSlice)\n")
	fmt.Printf("---- mapSlice	: %T, %p, %v\n", mapSlice, mapSlice, mapSlice)
	fmt.Printf("---- ok : %v\n", ok)
	if ok {
		for index, item := range *mapSlice {
			referencedItem := &(*mapSlice)[index]
			fmt.Printf("---- item.Value: %T, item: %p, key: %v, value: %v, value-pointer: %p, value-pointers pointer: %v\n", referencedItem.Value, referencedItem, referencedItem.Key, referencedItem.Value, referencedItem.Value, &referencedItem.Value)
			if referencedItem.Key == key {
				v, ok := referencedItem.Value.(yaml.MapSlice)
				if ok {
					return createOrderedMapSlice(o, referencedItem, &v)
				}
				return createOrderedMapSlice(o, referencedItem, &item.Value)
			}
		}
	}
	slice, ok := o.contents.(*interface{})
	fmt.Printf("--o.contents.(*interface{})\n")
	fmt.Printf("---- slice : %T, %v\n", slice, slice)
	fmt.Printf("---- ok : %v\n", ok)
	if ok {
		// > go - range over interface{} which stores a slice - Stack Overflow
		// > https://stackoverflow.com/questions/14025833/range-over-interface-which-stores-a-slice?answertab=active#tab-top
		switch reflect.TypeOf(*slice).Kind() {
		case reflect.Slice:
			fmt.Printf(">> slice is slice!\n")
			s := reflect.ValueOf(*slice)
			for i := 0; i < s.Len(); i++ {
				index := strconv.FormatInt(int64(i), 10)
				fmt.Printf("---- index : %v, key : %v\n", index, key)
				fmt.Printf("---- index == key : %v\n", index == key)
				if index == key {
					fmt.Printf("---- item: %T\n", s.Index(i).Interface())
					return createOrderedMapSlice(o, nil, s.Index(i).Interface())
				}
			}
		}

	}
	return createOrderedMapSlice(o, nil, nil)
}

func (o *OrderedMapSlice) Set(key string, value interface{}) {
	//success := false
	orderedMapSlice := o.Get(key)
	if orderedMapSlice == nil {
		return
	}
	fmt.Printf(">>>> item.Value: %T, item: %p, key: %v, value: %v, value-pointer: %p, value-pointers pointer: %v\n", orderedMapSlice.currentItem.Value, orderedMapSlice.currentItem, orderedMapSlice.currentItem.Key, orderedMapSlice.currentItem.Value, orderedMapSlice.currentItem.Value, &orderedMapSlice.currentItem.Value)
	printPointer("@@@@@ pointer orderedMapSlice.currentItem ", orderedMapSlice.currentItem, orderedMapSlice.currentItem)
	orderedMapSlice.currentItem.Value = value
	fmt.Printf(">>>> item.Value: %T, item: %p, key: %v, value: %v, value-pointer: %p, value-pointers pointer: %v\n", orderedMapSlice.currentItem.Value, orderedMapSlice.currentItem, orderedMapSlice.currentItem.Key, orderedMapSlice.currentItem.Value, orderedMapSlice.currentItem.Value, &orderedMapSlice.currentItem.Value)
	printPointer("@@@@@ pointer orderedMapSlice.currentItem ", orderedMapSlice.currentItem, orderedMapSlice.currentItem)

	//printPointer("@@@@@ pointer o          ", o, o)
	//printPointer("@@@@@ pointer reGetParent", reGetParent, reGetParent)

	//newCurrent := o.Get(key).contents
	//newOrdered := o.Get(key)
	//reGetCurrent := orderedMapSlice.Parent().Get(key).GetContents()
	//reGetOrderedMapSliceFromO := o.Get(key)
	//reGetOrdered := orderedMapSlice.Parent().Get(key)
	//
	//
	//
	//fmt.Printf(">>>> reGetOrderedMapSliceFromO item.Value: %T, item: %p, key: %v, value: %v, value-pointer: %p, value-pointers pointer: %v\n", reGetOrderedMapSliceFromO.currentItem.Value, reGetOrderedMapSliceFromO.currentItem, reGetOrderedMapSliceFromO.currentItem.Key, reGetOrderedMapSliceFromO.currentItem.Value, reGetOrderedMapSliceFromO.currentItem.Value, &reGetOrderedMapSliceFromO.currentItem.Value)
	//printPointer("@@@@@ pointer reGetOrderedMapSliceFromO.currentItem ", reGetOrderedMapSliceFromO.currentItem, reGetOrderedMapSliceFromO.currentItem)
	//fmt.Printf(">>>> reGetOrdered item.Value             : %T, item: %p, key: %v, value: %v, value-pointer: %p, value-pointers pointer: %v\n", reGetOrdered.currentItem.Value, reGetOrdered.currentItem, reGetOrdered.currentItem.Key, reGetOrdered.currentItem.Value, reGetOrdered.currentItem.Value, &reGetOrdered.currentItem.Value)
	//printPointer("@@@@@ pointer reGetOrdered.currentItem              ", reGetOrdered.currentItem, reGetOrdered.currentItem)

	//reGetParent := o.Get(key).Parent()
	//reGetOrderedMapSlice := reGetParent.Get(key)
	//printPointer("@@@@@ pointer orderedMapSlice             ", orderedMapSlice, orderedMapSlice)
	//printPointer("@@@@@ pointer orderedMapSlice.contents    ", orderedMapSlice.contents, orderedMapSlice.contents)
	//printPointer("@@@@@ pointer orderedMapSlice.currentItem ", *orderedMapSlice.currentItem, orderedMapSlice.currentItem)
	//printPointer("@@@@@ pointer reGetOrderedMapSlice.contents     ", reGetOrderedMapSlice.contents, reGetOrderedMapSlice.contents)
	//printPointer("@@@@@ pointer reGetOrderedMapSliceFromO.contents", reGetOrderedMapSliceFromO.contents, reGetOrderedMapSliceFromO.contents)
	//
	//printPointer("@@@@@ pointer orderedMapSlice                  ", orderedMapSlice, orderedMapSlice)
	//printPointer("@@@@@ pointer orderedMapSlice.contents          ", orderedMapSlice.contents, orderedMapSlice.contents)
	//printPointer("@@@@@ pointer reGetOrderedMapSlice.contents     ", reGetOrderedMapSlice.contents, reGetOrderedMapSlice.contents)
	//printPointer("@@@@@ pointer reGetOrderedMapSliceFromO.contents", reGetOrderedMapSliceFromO.contents, reGetOrderedMapSliceFromO.contents)
	//
	//fmt.Printf("@@@@@ set orderedMapSlice.GetContents() : %v\n", orderedMapSlice.GetContents())
	//fmt.Printf("@@@@@ set orderedMapSlice.contents : %v\n", orderedMapSlice.contents)
	////fmt.Printf("@@@@@ set o.Get(key).contents : %v\n", newCurrent)
	//printPointer("@@@@@ pointer orderedMapSlice.GetContents()", orderedMapSlice.GetContents(), orderedMapSlice.GetContents())
	//printPointer("@@@@@ pointer orderedMapSlice.contents", orderedMapSlice.contents, orderedMapSlice.contents)
	//printPointer("@@@@@ pointer reGetCurrent", reGetCurrent, reGetCurrent)
	//printPointer("@@@@@ pointer reGetOrdered", reGetOrdered, reGetOrdered)
	//printPointer("@@@@@ pointer orderedMapSlice", orderedMapSlice, orderedMapSlice)
	//printPointer("@@@@@ pointer o.Get(key).contents", newCurrent, newCurrent)
	//printPointer("@@@@@ pointer orderedMapSlice.Parent()", orderedMapSlice.Parent(), orderedMapSlice.Parent())
	//printPointer("@@@@@ pointer newOrdered.Parent()", newOrdered.Parent(), newOrdered.Parent())

	//mapSlice, ok := orderedMapSlice.contents.(*yaml.MapSlice)
	//fmt.Printf("-- orderedMapSlice.contents.(yaml.MapSlice)\n")
	//fmt.Printf("---- mapSlice	: %T, %v\n", mapSlice, mapSlice)
	//fmt.Printf("---- ok : %v\n", ok)
	//if ok {
	//	for i, item := range *mapSlice {
	//		fmt.Printf("---- item: %T, item.Key: %v, key: %v\n", item.Value, item.Key, key)
	//		if item.Key == key {
	//			(*mapSlice)[i].Value = value
	//		}
	//	}
	//}
	//fmt.Printf("---- mapSlice after : %T, %v\n", mapSlice, mapSlice)
}

func dumpNode(name string, node interface{}) {
	bytes, _ := yaml.Marshal(node)
	fmt.Printf("=======\n%v:\n%v=======\n", name, string(bytes))
}

func printPointer(name string, v interface{}, p interface{}) {
	fmt.Printf("-- %v => v %T: &v=%p v=&i=%p p=%p\n", name, v, &v, v, p)
}

type mapSliceUtil struct{}

var MapSliceUtil = mapSliceUtil{}

func (p *mapSliceUtil) Get(mapSlice *yaml.MapSlice, key string) *yaml.MapSlice {
	if mapSlice == nil {
		return nil
	}
	for _, item := range *mapSlice {
		if item.Key == key {
			v := item.Value.(yaml.MapSlice)
			return &v
		}
	}
	return nil
}

func (p *mapSliceUtil) Set(mapSlicePointer *yaml.MapSlice, key string, value interface{}) {
	success := false
	mapSlice := *mapSlicePointer
	// overwrite value
	//	printPointer("for mapSlicePointer before", *mapSlicePointer, mapSlicePointer)
	printPointer("for mapSlice before       ", mapSlice, &mapSlice)
	for i, item := range mapSlice {
		if item.Key == key {
			mapSlice[i].Value = value
			success = true
		}
	}
	//	printPointer("for mapSlicePointer after ", *mapSlicePointer, mapSlicePointer)
	printPointer("for mapSlice after        ", mapSlice, &mapSlice)
	// add new value
	if !success {
		//a := append(mapSlice, yaml.MapItem{Key: key, Value:value})
		//		printPointer("add mapSlicePointer before", *mapSlicePointer, mapSlicePointer)
		printPointer("add mapSlice before       ", mapSlice, &mapSlice)
		//		&mapSlice = &a
		//		printPointer("add mapSlicePointer after ", *mapSlicePointer, mapSlicePointer)
		printPointer("add mapSlice after        ", mapSlice, &mapSlice)
	}
}

func (p *mapSliceUtil) Delete(mapSlice yaml.MapSlice, key string) yaml.MapSlice {
	newMapSlice := yaml.MapSlice{}
	for _, item := range mapSlice {
		if item.Key != key {
			newMapSlice = append(newMapSlice, item)
		}
	}
	return newMapSlice
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

//func littleSmartMethod() {
//	var yamlString = getYaml()
//
//	node := &yaml.MapSlice{}
//	yaml.Unmarshal([]byte(yamlString), node)
//	// get value
//	developmentTeamsNode := MapSliceUtil.Get(node, `development-teams`)
//	fmt.Printf("%v\n", ">>> get value")
//	dumpNode("developmentTeamsNode", *developmentTeamsNode)
//	teamNode := MapSliceUtil.Get(developmentTeamsNode, `team-c`)
//	dumpNode("teamNode", *teamNode)
//
//	// set value (overwrite)
//	fmt.Printf("%v\n", ">>> set value (overwrite)")
//	printPointer("teamNode", *teamNode, teamNode)
//	MapSliceUtil.Set(teamNode, `members`, []string{"a", "b", "c"})
//	printPointer("teamNode", *teamNode, teamNode)
//	dumpNode("teamNode", *teamNode)
//
//	// set value (append)
//	fmt.Printf("%v\n", ">>> set value (append)")
//	MapSliceUtil.Set(teamNode, `location`, "tokyo")
//	MapSliceUtil.Set(teamNode, `skills`, []string{"java","php","golang"})
//	printPointer("teamNode", *teamNode, teamNode)
//	dumpNode("teamNode", *teamNode)
//
//	// delete value
//	products := MapSliceUtil.Get(teamNode, `products`)
//	MapSliceUtil.Set(teamNode, `products`, MapSliceUtil.Delete(*products, `ios-app-name10`))
//
//	// delete value
//	//MapSliceUtil.Set(developmentTeamsNode, `team-a`, nil)
//	//MapSliceUtil.Set(developmentTeamsNode, `team-d`, "new value")
//	//MapSliceUtil.Set(developmentTeamsNode, `team-e`, []string{"a", "b", "c"})
//	//MapSliceUtil.Set(&node, `development-teams`, MapSliceUtil.Delete(*developmentTeamsNode, `team-b`))
//	//teamNode = MapSliceUtil.Get(*developmentTeamsNode, `team-c`)
//	//products = MapSliceUtil.Get(*teamNode, `products`)
//	//MapSliceUtil.Set(teamNode, `products`, MapSliceUtil.Delete(*products, `ios-app-name10`))
//	//bytes, _ := yaml.Marshal(products)
//	//fmt.Printf("\n\n\n--- result:\n%v\n\n", string(bytes))
//
//	//	*developmentTeamsNode = yaml.MapSlice{}
//
//
////	dumpNode("teamNode", node)
//
//}

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

func getSimpleYaml() string {
	return `
# Development teams records
development-teams:
  team-a:
    pc-app-name1:
      id: 1001
    pc-app-name2:
      id: 1002
    ranks:
    - 100
    - 1000
`
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
  team-d: null
`
}
