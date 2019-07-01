I have written code in ruby v 2.4. Running the code is very simple, just hitting the run/start in the ruby terminal is enough.

Below is the demonstration of the code working:

1) First, we have to load the region details in the csv to the class variable, for that we must give the correct csv file path.
2) It will ask for the distributor name to create, followed by whether the distributor has a paren distributor. If parent distributor is present then it will ask for the name of the parent distributor, if the parent distributor is present it will assign that distributor as a parent or else it will through an error.
3) Then, it will ask for what of the assignment we want to make, include or exclude.
4) Then, when creating a first distributor it will ask user, whether they know how to assign regions if user types Y, then it will display the info of how to assign regions.
5) After that, it will ask for the regions to include. If the region typed is valid it will assign the region. Then, it will ask for whether we want to continue the assignment. If typed as Y, then will again start this point.
6) After assigning regions, it will ask for whether we need to create another distributor. If typed Y, then it will again start from the 2nd point.
7) After creating distributors, it will ask whether to check permissions for the distributors. If typed Y, then it will ask the distributor name and then the region. After that, it will validate the region and gives the result yes or no.
