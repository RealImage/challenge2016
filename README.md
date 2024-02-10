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

---------------------------- Arvinder Pal -----------------------
=> distributor struct represents a distributor with a name, inclusion, exclusion lists, and an array of sub-distributors.

=> Location struct holds information about a geographical location (city, state, country).
Global Variables:

=> LocationMap: A map that stores location information based on city, state, and country.

=> Distributors: A map that stores distributor information, where each distributor has a name and associated details.

=> The main function provides a menu-driven interface for the user to interact with the program.

=> Options include 
1. adding a distributor
2. adding permissions
3. searching permissions by location
4. adding a sub-distributor
5. exiting the program

=> AddPermission Function:
=> Allows the user to add permissions (inclusion or exclusion) to a distributor or sub-distributor based on geographical locations (city, state, country)
=> Checks if the parent distributor authorizes the permissions for a sub-distributor.
=> SearchByLocation Function:

=> Allows the user to search for permissions of a distributor or sub-distributor based on a specified location (city, state, country).

=> IsAuthorized Method:
=> Part of the distributor struct. Checks if a distributor or sub-distributor is authorized for a specific region (city, state, country) based on inclusion and exclusion lists.

=> FindParentDistributor Function:
Finds and returns the parent distributor of a given distributor name.
CheckDistributor Function:

=> Checks if a distributor or sub-distributor exists, and returns the distributor along with the relationship type (parent or child).

=>AddSubDistributor Function:
Allows the user to add a sub-distributor to an existing parent distributor.


