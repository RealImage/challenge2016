package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Name string
	Code string
	Values map[string]*Node
}

func main(){
	//root node which has all the distributor details
	root:=initroot()
	if root==nil{
		fmt.Println("failed to start please check the csv file")
		return
	}
	fmt.Println("welcome")

	for {
		var first string
		fmt.Println("1.To check access")
		fmt.Println("2.To grant access")
		fmt.Scanln(&first)
		if first == "1" {
			check := checkAccess(root)
			fmt.Println(check)

		} else if first=="2" {
			newUser(root)
			check := checkAccess(root)
			fmt.Println(check)

		}else{
			break
		}
	}

}

func checkAccess(root *Node)bool{
	access:=""
	distributor:=""
	fmt.Println("enter distributor name")
	fmt.Scanln(&distributor)
	fmt.Println("enter the value in the format of country name/provience name/city name")

	fmt.Scanln(&access)

	accessArr:=strings.Split(access,"/")
	temp:=root.Values[distributor].Values

	for index:= range accessArr{
		if _,check:=temp[accessArr[index]];check{
			temp=temp[accessArr[index]].Values
		}else{
			return false
		}
	}

	return true
}

func newUser(root *Node){

	distributor1:=""
	distributor2:=""
	access:=""

	fmt.Println("enter distributor1 name")
	fmt.Scanln(&distributor1)

	fmt.Println("enter distributor1 name")
	fmt.Scanln(&distributor2)

	fmt.Println("enter the value in the format of country name/provience name/city name")
	fmt.Scanln(&access)

	accessArr:=strings.Split(access,"/")
	temp:=root.Values
    rootTemp:=root.Values[distributor1]

	for index:= range accessArr{

		if val,check:=rootTemp.Values[accessArr[index]];!check || val==nil{
			fmt.Println("distributor1 doest have access in ",accessArr[index])
			root.Values[distributor2]=nil
			return
		}

		if val,check:=temp[accessArr[index]];check && val==nil{
			temp=temp[accessArr[index]].Values
			rootTemp=rootTemp.Values[accessArr[index]]
			continue
		}

		locationMap:=make(map[string]*Node)
		tempnode:=Node{
			Name: accessArr[index],
			Values: locationMap,
		}

		temp[accessArr[index]]=&tempnode
		temp=tempnode.Values
	    rootTemp=rootTemp.Values[accessArr[index]]
	}

	fmt.Println("access given",access)
}

func initroot()*Node {
	location:=make(map[string]*Node)
	distributor1Map:=make(map[string]*Node)
	distributor1:=Node{
		Name: "distributor1",
		Values:distributor1Map,
	}

	location["distributor1"]=&distributor1

	root:=Node{
		Name: "root",
		Values:location,
	}

	fd, error := os.Open("cities.csv")

	if error != nil {
		fmt.Println(error)
		return nil
	}

	fmt.Println("Successfully opened the CSV file")
	fileReader:= csv.NewReader(fd)
	records, error := fileReader.ReadAll()

	if error != nil {
		fmt.Println(error)
		return  nil
	}



	for _,record:= range records{

	if 	_,check:=distributor1.Values[record[5]];!check{
		tempMap:=make(map[string]*Node)
		temp:=Node{
			Name: record[5],
			Code: record[2],
			Values: tempMap,
		}
		distributor1.Values[record[5]]=&temp
	continue
	}
	node:=distributor1.Values[record[5]]
	if _,check:=node.Values[record[4]];!check{
			tempMap:=make(map[string]*Node)
			temp:=Node{
				Name: record[5],
				Code: record[1],
				Values: tempMap,
			}
			node.Values[record[4]]=&temp
	continue
	}

	node=distributor1.Values[record[5]].Values[record[4]]
	if _,check:=node.Values[record[3]];!check{
			tempMap:=make(map[string]*Node)
			temp:=Node{
				Name: record[5],
				Values: tempMap,
				Code: record[0],
			}
			node.Values[record[3]]=&temp
			continue
	}

}

return &root
}