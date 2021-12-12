import pandas as pd
import json

cityData = pd.read_csv("cities.csv")
    

data = cityData[['City Name','Country Name','Province Name']]
data = data.to_json(orient="records")
data = json.loads(data)

permissions = {}

def updateDistributorsData(dataset,distributor,Country,state='',city=''):
    if city != '':
        x = [d for d in dataset if d['City Name'] == city]
    elif state != '':
        x = [d for d in dataset if d['Province Name'] == state]
    else:
        x = [d for d in dataset if d['Country Name'] == Country]

    if len(x) < 1:
        print('Invalid Access Request')
        return 0
    
    if distributor in permissions.keys(): 
        for i in x:
            if i not in permissions[distributor]:
                permissions[distributor].append(i)
    else:
        permissions[distributor] = x
        

def removePermission(distributor,Country,state='',city=''):
    resultset = permissions[distributor]
    removalcount = 0
    if city != '':
        for i in resultset:
            if city == i['City Name']:
                resultset.remove(i)
                removalcount+=1
    elif state != '':
        for i in resultset:
            if state == i['Province Name']:
                resultset.remove(i)
                removalcount+=1
    else:
        for i in resultset:
            if Country == i['Country Name']:
                resultset.remove(i)
                removalcount+=1
    
    if removalcount < 1:
        print('The Distributor doesnt have the access already')
        return 0
    else:
        permissions[distributor] = resultset


def accessChoice(distributor,parentDistributor,InclrorExcl):
    accesslvl = int(input('The Level of access you want to give for the distributor : \n 1 -- > Country Level \n 2 -- > State/Province Level \n 3 -- > City Level\n Your Choice : '))
    
    if parentDistributor != '':
        dataset = permissions[parentDistributor]
    else:
        dataset = data

    if accesslvl in [1,2,3]:
        Country = input('Country : ')
        if accesslvl > 1:
            state = input('State/Province : ')
            if accesslvl > 2:
                city = input('City : ')
                return  updateDistributorsData(dataset,distributor,Country,state,city) if InclrorExcl else removePermission(distributor,Country,state,city)
            return updateDistributorsData(dataset,distributor,Country,state) if InclrorExcl else removePermission(distributor,Country,state)
        updateDistributorsData(dataset,distributor,Country) if InclrorExcl else removePermission(distributor,Country)
    else:
        print('OOps!!!Enter a Valid choice')

def addAccess(distributor,parentDistributor = ''):
    if not accessChoice(distributor,parentDistributor,True):
        print('Included Access for:',distributor)

def removeAccess(distributor,parentDistributor = ''):
    if not accessChoice(distributor,parentDistributor,False):
        print('Excluded Access for:',distributor)

def getdatafor(distributor,Country,state,city):
    hasAccess = False
    if city != '':
        for i in permissions[distributor]:
            if city == i['City Name']:
                hasAccess = True
                break
    elif state != '':
        for i in permissions[distributor]:
            if state == i['Province Name']:
                hasAccess = True
                break
    else:
        for i in permissions[distributor]:
            if Country == i['Country Name']:
                hasAccess = True
                break

    if hasAccess:
        print(distributor + ' has access to ' + Country + '-' + state + '-' + city)
    else:
        print(distributor + ' doesnt have access to ' + Country + '-' + state + '-' + city)

        

def checkDistributorhaspermtocity(name):
    accesslvl = int(input('Check permission based on : \n 1 -- > Country Level \n 2 -- > State/Province Level \n 3 -- > City Level\n Your Choice : '))

    if accesslvl in [1,2,3]:
        Country = input('Country : ')
        if accesslvl > 1:
            state = input('State/Province : ')
            if accesslvl > 2:
                city = input('City : ')
                return getdatafor(name,Country,state,city)
            return getdatafor(name,Country,state,'')
        return getdatafor(name,Country,'','')

def displayAccess():
    choice = int(input(' 1 ----> Check Access of Distributor \n 2 ----> Check Permission based on Distributor and Locationnb  '))
    name = input('Enter the Distributor name : ')
    if name in permissions.keys():
        if choice in [1,2,3]:
            if choice == 1:
                print('\n'.join(map(str,permissions[name])))
            else:
                checkDistributorhaspermtocity(name)
        else:
            print('Enter a valid choice')
    else:
        print(name + ' doesnt have any permissions !!!!')

print("Add Distributors")
while(1):
    distributor = input('ENTER DISTRIBUTOR NAME:   ')
    if distributor not in permissions.keys():
        inheritCheck = input('Get Access from Parent Distributor(Y/N): ')
        if inheritCheck.lower() == 'y' or inheritCheck.lower() == 'yes':
            while(1):
                parentDistributor = input('ENTER PARENT DISTRIBUTOR NAME :   ')
                if parentDistributor in permissions.keys():
                    print('Adding Permissions from Parent Distributor : ',parentDistributor)
                    break
                else:
                    print('The Specified Parent doesnt have any Access')
                    if not int(input('Do you want to add access from another parent distributor?(1/0)')):
                        parentDistributor = ''
                        break
        else:
            print('Proceeding to add a new distributor')
            parentDistributor = ''


        while(1):
            IncOrExc = int(input('1 for Inclusion , 2 for Exclusion : '))
            if IncOrExc == 1:
                addAccess(distributor,parentDistributor)
            else:
                removeAccess(distributor,parentDistributor)
            if not int(input('Do you want to add another Permission?(1/0)')):
                break
        
        addDistributorComplete = 0
        cancloseexecution = 0
        
        if not int(input('Do you want to add another Distributor?(1/0)')):
            addDistributorComplete = 1
        
        if addDistributorComplete:
            while(1):
                displayAccess()
                if not int(input('Do you want to check permission again?(1/0)')):
                    cancloseexecution = 1
                    break

        if cancloseexecution:
            break
    else:
        print('Access for this distributor is already set....')
