import csv
from collections import defaultdict
from bitarray import bitarray
import sys

treeMapCities = defaultdict(list)
# mapCodeCity = dict()
# mapCodeProvince = dict()
# mapCodeCountry = dict()

cityList = list()
indexMap = defaultdict(list)

DistributorMap = dict()

def init_maps():
	'''
	Map Country, Province, City and CODE
	Test
	print(treeMapCities["IN"]["TN"])
	print(mapCodeCity["CHENNAI_TAMILNADU_INDIA"])
	print(mapCodeProvince["TAMILNADU_INDIA"])
	print(mapCodeCountry["INDIA"])
	'''
	global cityList
	with open("cities.csv","r") as dataf:
		data = csv.reader(dataf)
		next(data) # skip header row

		for row in data:
			if treeMapCities[row[2]] == []:
				treeMapCities[row[2]] = defaultdict(list)

			treeMapCities[row[2]][row[1]].append(row[0])
			cityList.append("".join(row[5].upper().split()) + "_" + "".join(row[4].upper().split()) +"_" + "".join(row[3].upper().split()))

			# Map Codes (Not needed as of now, needed to map names to code in future if needed)
			# mapCodeCity["".join(row[3].upper().split()) + "_" + "".join(row[4].upper().split()) +"_" + "".join(row[5].upper().split())] = row[0]+ "_" + row[1] +"_" +row[2]
			# mapCodeProvince["".join(row[4].upper().split()) +"_" + "".join(row[5].upper().split())] = row[1] +"_" +row[2]
			# mapCodeCountry["".join(row[5].upper().split())] = row[2]

	# Form Index Map for bitarray creation and updation
	cityList = sorted(cityList)
	indJ = 0
	for city in cityList:
		indexMap[city].append(indJ)
		indexMap["_".join(city.split("_")[:2])].append(indJ)
		indexMap[city.split("_")[0]].append(indJ)
		indJ += 1

def createbitarray(rE, rI):
	temp = bitarray()
	if rI == []: # include all
		temp = bitarray(len(cityList))
		temp.setall(True)
	else:
		temp = bitarray(len(cityList))
		temp.setall(False)
		for ri in rI:
			temp[min(indexMap["_".join(ri.split("_")[::-1])]):max(indexMap["_".join(ri.split("_")[::-1])])+1] = True

	if rE == []: # nothing to exclude
		pass
	else:
		for re in rE:
			temp[min(indexMap["_".join(re.split("_")[::-1])]):max(indexMap["_".join(re.split("_")[::-1])])+1] = False

	return temp

def updatebitarray(DBaseArray, rE, rI):
	temp = DBaseArray.copy()
	modifiableIndex = []
	j=0
	for i in temp:
		if i == True:
			modifiableIndex.append(j)
		j+=1

	if rI == []:
		pass
	else:
		temp = bitarray(len(cityList))
		temp.setall(False)
		for ri in rI:
			for ind in indexMap["_".join(ri.split("_")[::-1])]:
				if ind in modifiableIndex:
					temp[ind] = True
	if rE == []:
		pass
	else:
		for re in rE:
			for ind in indexMap["_".join(re.split("_")[::-1])]:
				if ind in modifiableIndex:
					temp[ind] = False
	return temp

def setbitarray(DNew, DBase, rE, rI):
	if DBase == "": # new distributor
		DistributorMap[DNew] = createbitarray(rE, rI)
	else: # derived distributor
		DistributorMap[DNew] = updatebitarray(DistributorMap[DBase], rE, rI)

line = -1
def get_persmissions():
	'''
	Process the permissions.txt file and build bitmaps
	'''
	data = [i.strip() for i in open(sys.argv[1],"r").readlines()]
	def nextLine():
		global line
		line+=1
		if line < len(data):
			return data[line]
		else:
			return "END"

	l = nextLine()
	while "END" not in l:
		if "DISTRIBUTOR" in l:
			DistributorPath = [i.strip() for i in l.split("<")]
			rulesE = []
			rulesI = []
			l = nextLine()
			while "NEXT" not in l and "END" not in l:
				if "INCLUDE" in l:
					rulesI.append(l.split()[1].strip())
				elif "EXCLUDE" in l:
					rulesE.append(l.split()[1].strip())
				l = nextLine()

			# print(DistributorPath)
			if len(DistributorPath) > 1:
				DistributorNew, DistributorBase = DistributorPath[:2]
				setbitarray(DistributorNew, DistributorBase, rulesE, rulesI)
			else:
				DistributorNew = DistributorPath[0]
				setbitarray(DistributorNew, "", rulesE, rulesI)
			l = nextLine()

def check_permission(DName,Location):
	A = DistributorMap[DName].copy()
	B = bitarray(len(cityList))
	B.setall(False)
	B[min(indexMap["_".join(Location.split("_")[::-1])]):max(indexMap["_".join(Location.split("_")[::-1])])+1] = True

	if A & B == B:
		return "YES"
	else:
		return "NO"

if __name__ == '__main__':
	init_maps()
	print("Done Mapping.")
	get_persmissions()
	print("Done Formulating Permissions.")
	print()

	# Uncomment for continuous input
	# yn = "Y"
	# while yn == "Y" or yn.upper() == "YES":
	# 	D = input("Distributor Name: ")
	# 	L = input("Location: ")
	# 	print(D + " has permission over " + L + " ?")
	# 	print(check_permission(D,L))
	# 	yn = input("Do you want to continue? (Y/N): ")

	# Uncomment to use input.txt
	# for line in open("input.txt").readlines():
	# 	D, L, Ans = line.strip().split(",")
	# 	print(D + " has permission over " + L + " ?")
	# 	print("Ground Truth: " + Ans + " Programs Answer: " + check_permission(D,L))
	# 	print()

	# For command line arguments
	print(sys.argv[2] + " has permission over " + sys.argv[3] + " ?")
	print(check_permission(sys.argv[2],sys.argv[3]))
