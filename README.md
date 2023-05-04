# Real Image Challenge 2016

This program checks if a given city is allowed to be distributed to by a distributor or not.

To get started with this program, follow these steps:

# Usage
When prompted, enter distributor permissions in the following format:

INCLUDE: <country_code> - allows distribution in the specified country
EXCLUDE: <state_code>-<country_code> - disallows distribution in the specified state of the specified country

For example, you can try the following as input.

Enter the distributor permissions:
INCLUDE: IN
EXCLUDE: TN-IN
EXCLUDE: KA-IN

Enter city:VISAKHAPATNAM-AP-IN

When prompted, enter the city name, state code, and country code to check if it is allowed for distribution
The program will output either "YES" or "NO" .
This program uses data from the cities.csv file to check if a city is allowed or not.

This program checks distributor permissions in the following way:

If the distributor has not included the country in their permissions, the city is not allowed.
If the distributor has excluded the state in the country from their permissions, the city is not allowed.
If the city is not in the database for the specified country and state, the city is not allowed.



