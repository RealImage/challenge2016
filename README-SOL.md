Rule 1 => Every geographical place must begin and end with "" (double qoute). -->
Rule 2 => Every geographical place must be separated from each other by :(single colon sign) within the ending double quote.
For example => correct => "MADIKERI-KARNATAKA-INDIA:"
incorrect => MADIKERI-KARNATAKA-INDIA:KOLKATA-WEST BENGAL-INDIA:"
incorrect => MADIKERI-KARNATAKA-INDIA:
incorrect => "MADIKERI-KARNATAKA-INDIA:
incorrect => "MADIKERI-KARNATAKA-INDIA"
Rule 3 => The Rule 1 and Rule 2 applies for distributor name also.
For example => correct => "Tom:Harry:Sejul" with Sejul is the GrandParent Distributor, Harry is the parent distributor and Tom is the child.
Rule 4 => Dummy Run the program on the first attemp to understand the process of executing the python script.
Rule 5 => If you want to run the program with multiple search query then please enter one query(pair of single geographic place and single name) in every cycle of run and continue pressing 'y' (small letter y)


 <!-- pyhton3 app.py
 -I <included geographical area/ [Place name must be same with CSV file[ [Please follow README-SOL.md]/each place ends with ':']>
  -E <excluded geographical area / [Place name must be same with CSV file [Please follow README-SOL.md]/each place ends with ':']>
 -D <distributors name / [Each distributor name must end with ':']>
 -F <name of csv file>

Please enter the distributor name now[No Spaces in each name/ Name must end with ':']   :"Tom:"
Please enter the geographical place where distributor's PERMISSION [No Spaces in each name/ Name must end with ':']     :"KOLKATA-WEST BENGAL-INDIA:"
Tom:
2Yes.    The distributor Tom has PERMISSION to release film in these input regions [KOLKATA-WEST BENGAL-INDIA-
Press y to continue entering geo locations.Any key to exit      :y
Please enter the distributor name now[No Spaces in each name/ Name must end with ':']   :"Tom:"
Please enter the geographical place where distributor's PERMISSION [No Spaces in each name/ Name must end with ':']     :"MADIKERI-KARNATAKA-INDIA:"
Tom:
No.      The distributor Tom has no PERMISSION to release film in these input regions [MADIKERI-KARNATAKA-INDIA-
Press y to continue entering geo locations.Any key to exit      :y
Please enter the distributor name now[No Spaces in each name/ Name must end with ':']   :"tom:"
Please enter the geographical place where distributor's PERMISSION [No Spaces in each name/ Name must end with ':']     :"KARNATAKA-INDIA:"
Tom:
No.      The distributor Tom has no PERMISSION to release film in these input regions [KARNATAKA-INDIA-
Press y to continue entering geo locations.Any key to exit      :y
Please enter the distributor name now[No Spaces in each name/ Name must end with ':']   :"Tom:"
Please enter the geographical place where distributor's PERMISSION [No Spaces in each name/ Name must end with ':']     :"Kolkata-West Bengal:"
Tom:
2Yes.    The distributor Tom has PERMISSION to release film in these input regions [Kolkata-West Bengal-
Press y to continue entering geo locations.Any key to exit      :t

 -->
