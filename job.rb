require './region_distributor.rb'

	reg_distributor=RegionDistributor.new
	puts "No of distributors to be inputted"
	number = gets.chomp
	for i in 0..(number-1)
 	  puts "enter name"
		name = gets.chomp 
 	  puts "enter included regions as an array"
		inc_region = gets.chomp
	  puts "enter exxcluded regions as an array"
   		exc_region = gets.chomp
		reg_distributor.distributor({:name=>name,:inc_region=>inc_region.to_i,:exc_region=>exc_region.to_i})	
	end
        
	continue = 'y' 
        while (continue == 'y')
		puts "enter distributor to be found"
		distributor_name = gets.chomp
		puts "enter region to be found"
		region = gets.chomp
		permission = check_permissions(distributor_name,region)
		puts permission
		puts "continue? y : n"
		continue = gets.chomp
	end
