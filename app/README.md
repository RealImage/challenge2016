**Running the program**

* cd app/
* ruby main.rb

**About the code**

The logic passes all sample test cases and other cases I made up. Currently it does not make use of the CSV data at all. If we can ensure that the input data is correct the code can satisfy all the conditions without prior knowledge of the regions.

**How it works**

When a new distributor is added it creates a new instance of the Distributor object. This object saves inclusions and exclusions and can also extend data from another distributor object (parent distributor).

It builds a hash table for inclusions and exclusions. Authorization checks are performed against these hashes so the algorithm works fast at O(1) time complexity.

Key logic is inside the *authorization_for* private method inside Distributor class.

**Notes**

* Running main.rb gives you a prompt with which you can control distributor logic
* You can find commented out lines for quick debugging at the end of main.rb
* Area codes interpreted by the parser should be of the format **CITY::PROVINCE::COUNTRY**
