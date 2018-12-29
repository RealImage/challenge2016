const config = {
  main_menu: `
  ================ Welcome to Real Image Challenge ================
          This is highly accurate, consistent in memory next generation system.
          Here are the options for you
          1. Create a new distributor
          2. Relate the distributors (example DISTRIBUTOR2 < DISTRIBUTOR1)
          3. List all distributors
          4. Query distributor to area code
          0. To exit
          You need to select 1 or 2 or 3 or 4`,
  menu_error: `Sorry we did not understand your request please try again`,
  distributor_menu: `
  ================ Welcome to Distributor menu ================
        1. Create a distributor
        0. Exit
        You need to select 0 or 1
  `,
  add_distributor: `
First enter the distributor name then a comma separated list of includes, 
the places where he is eligible to sell and then comma separed list of excludes 
where he is not allowed to sell.
The places has to be in codes and if you want to be more specific then go from least specific
to more specific for example Chennai-Tamil Nadu-India and not India-Tamil Nadu-Chennai

Example
  Enter distributor name
  Distributor1
  INCLUDES
  INDIA,UNITEDSTATES
  EXCLUDES
  KARNATAKA-INDIA,CHENNAI-TAMILNADU-INDIA
  `,
  relate_distributor: `
  ==================Welcome to Relate Distributor menu===============
  Enter the relationship like Distributor2 < Distributor1 which means
  Distributor2 has permissions less than Distributor1
  If there are multi level relations then do them one after another
  For example if you want Distributor3 <  Distributor2 < Distributor1
  Then do Distributor2 < Distributor1
  Then Distributor3 < Distributor2
  `,
  query_distributor: `
  =================Welcome to Query Distributor menu==================
  Enter a distributor name to query. Followed by the place to query
  Example
  Enter distributor name
  d1
  Enter place to query
  CHICAGO-ILLINOIS-UNITEDSTATES
  `
};

module.exports = config;
