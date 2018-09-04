**Running the program**

* `cd app/`
* `ruby main.rb`

**About the code**

The logic passes all sample test cases and other cases I made up. Currently it does not make use of the cities CSV data at all. If we can ensure that the input data is correct the code can satisfy all the conditions without prior knowledge of the regions.

**How it works**

When a new distributor is added it creates a new instance of the Distributor object. This object saves inclusions and exclusions and can also extend data from another distributor object (parent distributor).

It builds a hash table for inclusions and exclusions. Authorization checks are performed against these hashes so the algorithm works fast at O(1) time complexity.

Key logic is inside the *authorization_for* private method inside Distributor class.

**Notes**

* Users can now easily edit input commands to the program by updating `app/commands.csv`
* On Running `app/main.rb` the program performs the commands and prints output on the console
* Area codes should be of the format **CITY::PROVINCE::COUNTRY** to be interpreted correctly by the parser

**Commands**

The command formats accepted by `app/commands.csv` are listed below

* CREATE,distributor name
* INCLUDE,distributor name,AREA::CODE
* EXCLUDE,distributor name,AREA::CODE
* EXTEND,child distributor name,parent distributor name
* VERIFY,distributor name,AREA::CODE
