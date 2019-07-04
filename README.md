# test.js - Working

checkDistributorAccess is the main function that I have implemented.

There is a distributor 'dist1',

```
    const dist1 = {
        name: "d1",
        includedPlaces: ["IN", "US"],
        excludedPlaces: ["KA-IN", "CH-TN-IN"]
    };

```

Here I am using place codes for storing data, 'includedPlaces' is the array of places where the distributor have distributing access, 'excludedPlaces' is the array of places where distributor is restricted.

If I have to find whether a distributor have access in 'TN-IN' I will call 'checkDistributorAccess('TN-IN')' like that.

How to Run the application ?

command -

```
node test.js
```

if you want to check the access on any different area, you have to edit the code and call the checkDistributorAccess with that particular place code - save the code - and run again

" \***_ Here I have just included my logic, not created any application due to my time concern. Please forgive that. _** "
