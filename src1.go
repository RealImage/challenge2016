//will give the first set of wrong data for all the distributor, because it may lead to inconsistency
// when we try to follow the underlying distributor based upon the distributor above
//whenever the program gives the first record of wrong data for a particular distributor, remove it to find the
//next record of wrong data
package main

import (
	"fmt"
	"strings"
	//"reflect"
)



func removeDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

func F(data string) (map[string][]string, map[string][]string, []string, []string, []string) {

	strLine := strings.Split(data, "\n")
	strWord := make([]string, 3)
	strcountry := make([]string, len(strLine))
	strprovince := make([]string, len(strLine))
	strcity := make([]string, len(strLine))
	for i := 0; i < len(strLine); i++ {
		strWord = strings.Split(strLine[i], "\t")
		strcountry[i] = strWord[2]
		strprovince[i] = strWord[1]
		strcity[i] = strWord[0]
	}

	strcountry = removeDuplicates(strcountry)
	strprovince = removeDuplicates(strprovince)
	strcity = removeDuplicates(strcity)
/*
	a := make(map[string][]string, len(strprovince))
	for j := 0; j < len(strprovince); j++ {
		for i := 0; i < len(strLine); i++ {
			strWord = strings.Split(strLine[i], "\t")
			if strings.EqualFold(strprovince[j], strWord[1]) == true {
				a[strprovince[j]] = append(a[strprovince[j]], strWord[0])
			}
		}
	}

	b := make(map[string][]string, len(strcountry))
	for j := 0; j < len(strcountry); j++ {
		for i := 0; i < len(strLine); i++ {
			strWord = strings.Split(strLine[i], "\t")
			if strings.EqualFold(strcountry[j], strWord[2]) == true {
				b[strcountry[j]] = append(b[strcountry[j]], strWord[1])
			}
		}
		b[strcountry[j]] = removeDuplicates(b[strcountry[j]])
	}
	*/
//here b and a are not used but can be used to query, replace nil with a and b to send them to main function
	return nil, nil, strcountry, strprovince, strcity
}
func isAvailable(x map[string][]string, str1 string, str2 string) int {
	flag := 0
	for i := 0; i < len(x[str1]); i++ {
		if str2 == x[str1][i] {
			flag = 1
			break
		}
	}
	return flag
}

func checkForSubset(s1 []string, s2 []string) (bool,int) {
	//fmt.Println("s1: ", s1, "\n")
	//fmt.Println("s2: ", s2, "\n")
	if len(s2) > len(s1) {
		return false,-1
	}
	encountered := make([]bool, len(s2))
	//fmt.Println("encountered: ", len(encountered))

	for i := 0; i < len(s2); i++ {
		for j := 0; j < len(s1); j++ {
			if s2[i] == s1[j] {
				encountered[i] = true
			}
		}
	}
	//fmt.Println("encountered: ", encountered)
	for i := 0; i < len(encountered); i++ {
		if encountered[i] == false {
			return false,i
		}
	}
	return true,-1
}
func main() {
	data :=[]string {`CANIO	2	AD
ETATR	2	AD
ENCAM	3	AD
PLCSA	3	AD
ARINL	4	AD
LMSNA	4	AD
ORDIN	5	AD
SJUAL	6	AD
DFUIR	FU	AE
JZRAH	RK	AE
AKUPR	BAL	AF
CHIMT	BAL	AF
DEHAI	BAL	AF
DOWAB	BAL	AF
LABSA	BAL	AF
QACIG	BAL	AF
QRGHT	BAL	AF
AEKMR	BDG	AF
BMURG	BDG	AF
JWAND	BDG	AF
QADIS	BDG	AF
QALIN	BDG	AF
SATEH	BDG	AF
ADKWM	BDS	AF
ASKHM	BDS	AF
BETSH	BDS	AF`,

`CANIO	2	AD
ETATR	2	AD
ENCAM	3	AD
QACIG	BAL	AF
QRGHT	BAL	AF
PLCSA	3	AD
ARINL	4	AD`,

`CANIO	2	AD
ETATR	2	AD
ENCAM	3	AD
QACIG	BAL	AF
QRGHT	BAL	AF`,

`CANIO	2	AD
ETATR	2	AD
QADIS	BDG	AF
SATEH	BDG	AF
ENCAM	3	AD`}
	
//for simplicity and being not allowed to use any database, I have included the data in the program itself as 
//a variable
	n:=len(data)
	fmt.Println("no. of distributors is ",n-1)
	fmt.Println("data[0] is the original data, data[1],data[2]...data[n-1] denotes distributors")
//	b:=make([]map[string][]string,n)
//	 a:=make([]map[string][]string,n)
	 strcountry:=make([][]string,n)
	strprovince:=make([][]string,n)
	 strcity:=make([][]string,n)
for i:=0;i<n;i++ {
	_,_, strcountry[i], strprovince[i], strcity[i] = F(data[i])
// here intentionally made _,_,  coz., when b[0]["AD"] 	is queried, we can get list of provinces under country "AD"
//similarly for a:province to city mapping	
	}
	
	
				fun:=func(strTemp [][]string) (int,int,string) {
				flag:=0
				var j,i int
				var x bool
				var position int
				for j=0;j<n-1;j++ {
				for i=j+1;i<n;i++ {
						x,position =checkForSubset(strTemp[j],strTemp[i])
						//fmt.Println("x:",x,"\n")
						if x==false {
							flag=1;
							goto l;
						}
				}
				}
				l:
				if flag==1 {
					
					return j,i,strTemp[i][position]
				}else {
					
					return -1,-1,""
				}
				
				}
	
	j,i,problem:=fun(strcountry)
	if j!=-1 && i!=-1 && problem!=""{
		if j==0 {
			fmt.Println("with distributor : ",i,", in country ",problem)
		}else {
		fmt.Println("between distributor ",j," and ",i," and country is",problem)
		}
	}else {
		fmt.Println("country-working fine")
	}
	
	j,i,problem=fun(strprovince)
	if j!=-1 && i!=-1 && problem!=""{
		if j==0 {
			fmt.Println("with distributor : ",i,", in province ",problem)
		}else {
		fmt.Println("between distributor ",j," and ",i," and province is",problem)
		}
	}else {
		fmt.Println("province-working fine")
	}
	
	j,i,problem=fun(strcity)
	if j!=-1 && i!=-1 && problem!=""{
		if j==0 {
			fmt.Println("with distributor : ",i,", in city ",problem)
		}else {
		fmt.Println("between distributor ",j," and ",i," and city is",problem)
		}
	}else {
		fmt.Println("city-working fine")
	}
	for i=0;i<len(strprovince);i++ {
	fmt.Println("\n",strprovince[i])
	}
	
	
}
