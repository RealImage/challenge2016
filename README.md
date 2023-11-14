cli with 3 options,
    check permission -
        select a distributor then select a city 
        return if distributor is allowed in that city or not
    add - 
        add a distributor - enter distributor Name
            four options includes/ excludes/ parent/ confirm
                select includes ->
                    city - select city to include
                    state - 
                    country - 
                similarly select excludes,
                parent will show list of available distributor and option select one
                confirm -> save and exit
    exit

cities:
0 ct-state-country
1 ct-state-country

distributor:
    name: 'distro'
    includes: [1, 2 ]
    excludes: [3, 4]
    parent: [1]

