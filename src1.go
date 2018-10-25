package main

import (
	"fmt"
	"strings"
)

const data = `CANIO	2	AD
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
BETSH	BDS	AF`

const dist1 = `CANIO	2	AD
ETATR	2	AD
ENCAM	3	AD
QACIG	BAL	AF
QRGHT	BAL	AF
PLCSA	3	AD
ARINL	4	AD`

const dist2 = `CANIO	2	AD
ETATR	2	AD
ENCAM	3	AD
QACIG	BAL	AF
QRGHT	BAL	AF`

const dist3 = `CANIO	2	AD
ETATR	2	AD
ENCAM	3	AD`

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

	return b, a, strcountry, strprovince, strcity
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

func checkForSubset(s1 []string, s2 []string) bool {
	fmt.Println("s1: ", s1, "\n")
	fmt.Println("s2: ", s2, "\n")
	if len(s2) > len(s1) {
		return false
	}
	encountered := make([]bool, len(s2))
	fmt.Println("encountered: ", len(encountered))

	for i := 0; i < len(s2); i++ {
		for j := 0; j < len(s1); j++ {
			if s2[i] == s1[j] {
				encountered[i] = true
			}
		}
	}
	fmt.Println("encountered: ", encountered)
	for i := 0; i < len(encountered); i++ {
		if encountered[i] == false {
			return false
		}
	}
	return true
}
func main() {
	var n int
	fmt.Println("enter no. of distributors")
	fmt.Scanln(n)
//for i=0:i<n;i++ {
	//b[i], a[i], strcountry[i], strprovince[i], strcity[i] = F(data[i])
	b, a, strcountry, strprovince, strcity = F(data)
	b1, a1, strcountry1, strprovince1, strcity1 = F(dist1)
	b2, a2, strcountry2, strprovince2, strcity2 = F(dist2)
	b3, a3, strcountry3, strprovince3, strcity3 = F(dist3)
	//}
	for i:=0;i<
	
}
