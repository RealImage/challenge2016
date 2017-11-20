###################################

Environment:

Ruby - 2.3.1

###################################

Assumptions:

1. Include(Allowed regions) should always be there
2. Validations are done for inputs like valid region, valid distributor, multi level hierarchy etc
1. Incase of not providing Exclude(Unallwed regions) give it as NULL Eg. add Distributor3 IN NULL Distributor2
2. Exclude regions will be a valid one i.e a subset of include

Files:

1. regions.csv - Cities.csv has so many data inorder to load the data quick created regions.csv
2. process.rb - Loads the data from regions.csv. Gets the input from user and prints the result
3. csv_util - Reads file and gives parsed csv result
4. distributor.rb - Create new Distributor objects and check permissions
5. country.rb - Has country objects
6. state.rb - has state objects
7. city.rb - has city objects
8. test_cases folder contains test cases for the above files

###################################

Schema: 

Country - code, name
State - code, name, country_code
City - code, name, country_code, state_code
Distributor - name, allowed_regions, unallowed_regions, extends
Each of them treated as Models instead od DB, all the inputs are stored in-memory as objects

Command to run the program: 

$ ruby process.rb

Inputs are given by user (add distributor / check permission od distributor). inputs are given by space separated

1. Add a distributor in this format 
add Distributor Name,Allowed regions,Unallowed regions,Extends 

add Distributor1 IN,IT KA-IN,CENAI-TN-IN)
add Distributor2 IN TN-IN Distributor1
add Distributor3 IN NULL Distributor2

2. Please provide distributor name and location for permission as 

(Eg. permission? Distributor1 BENAU-KA-IN)

###################################


# Real Image Challenge 2016

In the cinema business, a feature film is usually provided to a regional distributor based on a contract for exhibition in a particular geographical territory.

Each authorization is specified by a combination of included and excluded regions. For example, a distributor might be authorzied in the following manner:
```
Permissions for DISTRIBUTOR1
INCLUDE: INDIA
INCLUDE: UNITEDSTATES
EXCLUDE: KARNATAKA-INDIA
EXCLUDE: CHENNAI-TAMILNADU-INDIA
```
This allows `DISTRIBUTOR1` to distribute in any city inside the United States and India, *except* cities in the state of Karnataka (in India) and the city of Chennai (in Tamil Nadu, India).

At this point, asking your program if `DISTRIBUTOR1` has permission to distribute in `CHICAGO-ILLINOIS-UNITEDSTATES` should get `YES` as the answer, and asking if distribution can happen in `CHENNAI-TAMILNADU-INDIA` should of course be `NO`. Asking if distribution is possible in `BANGALORE-KARNATAKA-INDIA` should also be `NO`, because the whole state of Karnataka has been excluded.

Sometimes, a distributor might split the work of distribution amount smaller sub-distiributors inside their authorized geographies. For instance, `DISTRIBUTOR1` might assign the following permissions to `DISTRIBUTOR2`:

```
Permissions for DISTRIBUTOR2 < DISTRIBUTOR1
INCLUDE: INDIA
EXCLUDE: TAMILNADU-INDIA
```
Now, `DISTRIBUTOR2` can distribute the movie anywhere in `INDIA`, except inside `TAMILNADU-INDIA` and `KARNATAKA-INDIA` - `DISTRIBUTOR2`'s permissions are always a subset of `DISTRIBUTOR1`'s permissions. It's impossible/invalid for `DISTRIBUTOR2` to have `INCLUDE: CHINA`, for example, because `DISTRIBUTOR1` isn't authorized to do that in the first place. 

If `DISTRIBUTOR2` authorizes `DISTRIBUTOR3` to handle just the city of Hubli, Karnataka, India, for example:
```
Permissions for DISTRIBUTOR3 < DISTRIBUTOR2 < DISTRIBUTOR1
INCLUDE: HUBLI-KARNATAKA-INDIA
```
Again, `DISTRIBUTOR2` cannot authorize `DISTRIBUTOR3` with a region that they themselves do not have access to. 

We've provided a CSV with the list of all countries, states and cities in the world that we know of - please use the data mentioned there for this program. *The codes you see there may be different from what you see here, so please always use the codes in the CSV*. This Readme is only an example. 

Write a program in any language you want (If you're here from Gophercon, use Go :D) that does this. Feel free to make your own input and output format / command line tool / GUI / Webservice / whatever you want. Feel free to hold the dataset in whatever structure you want, but try not to use external databases - as far as possible stick to your langauage without bringing in MySQL/Postgres/MongoDB/Redis/Etc.

To submit a solution, fork this repo and send a Pull Request on Github. 

For any questions or clarifications, raise an issue on this repo and we'll answer your questions as fast as we can.


