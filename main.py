import csv
from collections import defaultdict
import numpy as np
import sys

# Try importing bitarray (Each distributor uses 96 Bytes of RAM at maximum) else "numpy" is used with 77kB info per distributor
np_flag = 0
try:
	from bitarray import bitarray
except Exception as e:
	np_flag = 1

np_flag = 1

cityList = list()
indexMap = defaultdict(list)

DistributorMap = dict()

def init_maps():

	global cityList
	with open("cities.csv","r") as dataf:
		data = csv.reader(dataf)
		next(data) # skip header row

		for row in data:
			cityList.append("".join(row[5].upper().split()) + "_" + "".join(row[4].upper().split()) +"_" + "".join(row[3].upper().split()))

	# Form Index Map for bitarray creation and updation
	cityList = sorted(cityList)
	indJ = 0
	for city in cityList:
		indexMap[city].append(indJ)
		indexMap["_".join(city.split("_")[:2])].append(indJ)
		indexMap[city.split("_")[0]].append(indJ)
		indJ += 1

def createbitarray(rE, rI):
	if np_flag == 1:
		temp = np.zeros(len(cityList), dtype=bool)
		if rI == []: # include all
			temp = np.ones(len(cityList), dtype=bool)
		else:
			for ri in rI:
				temp[min(indexMap["_".join(ri.split("_")[::-1])]):max(indexMap["_".join(ri.split("_")[::-1])])+1] = True

		if rE == []: # nothing to exclude
			pass
		else:
			for re in rE:
				temp[min(indexMap["_".join(re.split("_")[::-1])]):max(indexMap["_".join(re.split("_")[::-1])])+1] = False

	elif np_flag == 0:
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
	modifiableIndex = list(np.nonzero(temp)[0])

	if np_flag == 1:
		if rI == []:
			pass
		else:
			temp = np.zeros(len(cityList), dtype=bool)
			for ri in rI:
				idx = list(set(indexMap["_".join(ri.split("_")[::-1])]).intersection(set(modifiableIndex)))
				temp[idx] = True

		if rE == []:
			pass
		else:
			for re in rE:
				idx = list(set(indexMap["_".join(re.split("_")[::-1])]).intersection(set(modifiableIndex)))
				temp[idx] = False

		return temp

	elif np_flag == 0:
		if rI == []:
			pass
		else:
			temp = bitarray(len(cityList))
			temp.setall(False)
			temp = np.array(temp)
			for ri in rI:
				idx = list(set(indexMap["_".join(ri.split("_")[::-1])]).intersection(set(modifiableIndex)))
				temp[idx] = True
			t = bitarray()
			t.pack(temp.tostring())

		if rE == []:
			pass
		else:
			temp = np.array(temp)
			for re in rE:
				idx = list(set(indexMap["_".join(re.split("_")[::-1])]).intersection(set(modifiableIndex)))
				temp[idx] = False
			t = bitarray()
			t.pack(temp.tostring())

		return t

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

			if len(DistributorPath) > 1:
				DistributorNew, DistributorBase = DistributorPath[:2]
				setbitarray(DistributorNew, DistributorBase, rulesE, rulesI)
			else:
				DistributorNew = DistributorPath[0]
				setbitarray(DistributorNew, "", rulesE, rulesI)
			l = nextLine()

def check_permission(DName,Location):
	A = DistributorMap[DName].copy()

	if np_flag == 1:
		B = np.zeros(len(cityList), dtype=bool)
		B[min(indexMap["_".join(Location.split("_")[::-1])]):max(indexMap["_".join(Location.split("_")[::-1])])+1] = True

		if np.array_equal((A & B),B):
			return "YES"
		else:
			return "NO"

	elif np_flag == 0:
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
	# 	# print()

	# For command line arguments
	print(sys.argv[2] + " has permission over " + sys.argv[3] + " ?")
	print(check_permission(sys.argv[2],sys.argv[3]))
