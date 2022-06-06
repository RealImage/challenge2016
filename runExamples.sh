#!/bin/bash

# build mai.go to the current directory
GO111MODULE=off; go build -o ./perms ./main.go

printf "Run 'perms' daemon...\n"
./perms -daemon &
sleep 1

printf "Create PARENT distributor \n"
./perms -command=create -newDist=PARENT -include=Matanzas-Cuba,Mayabeque-Cuba,Artemisa-Cuba -exclude=Alacranes-Matanzas-Cuba
printf "\n"

printf "get PARENT distributor permissions \n"
./perms -command=get -dist=PARENT
printf "\n"

printf "check if PARENT distributor has permissions to Matanzas-Cuba \n"
./perms -command=has -dist=PARENT -region=Matanzas-Cuba
printf "\n"

printf "check if PARENT distributor has permissions to Alacranes-Matanzas-Cuba \n"
./perms -command=has -dist=PARENT -region=Alacranes-Matanzas-Cuba
printf "\n"

printf "check if PARENT distributor has permissions to Batabano-Mayabeque-Cuba \n"
./perms -command=has -dist=PARENT -region=Batabano-Mayabeque-Cuba
printf "\n"

printf "create CHILD distributor \n"
./perms -command=create -dist=PARENT -newDist=CHILD -include=Mayabeque-Cuba,Artemisa-Cuba -exclude=Batabano-Mayabeque-Cuba,Cabanas-Artemisa-Cuba
printf "\n"

printf "get CHILD distributor permissions \n"
./perms -command=get -dist=CHILD
printf "\n"

printf "check if CHILD distributor has permissions to Matanzas-Cuba \n"
./perms -command=has -dist=CHILD -region=Matanzas-Cuba
printf "\n"

printf "check if CHILD distributor has permissions to Mayabeque-Cuba \n"
./perms -command=has -dist=CHILD -region=Mayabeque-Cuba
printf "\n"

printf "check if CHILD distributor has permissions to Batabano-Mayabeque-Cuba \n"
./perms -command=has -dist=CHILD -region=Batabano-Mayabeque-Cuba
printf "\n"

printf "update PARENT distributor permissions: delete Mayabeque-Cuba from INCLUDES\n"
./perms -command=update -dist=PARENT -include=Mayabeque-Cuba
printf "\n"

printf "get CHILD distributor permissions \n"
./perms -command=get -dist=CHILD
printf "\n"

printf "remove PARENT permissions \n"
./perms -command=remove -dist=PARENT
printf "\n"

printf "Killing 'perms' daemon... \n"
pkill perms
