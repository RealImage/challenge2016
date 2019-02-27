require 'json'

  class String
		def squish
	  	gsub!(/\A[[:space:]]+/, '')
	  	gsub!(/[[:space:]]+\z/, '')
	  	gsub!(/[[:space:]]+/, ' ')
	  	self
		end
	end

class Qube

  def input_parser(json_file_path)
  	file_input = File.read(json_file_path)
  	json = JSON.parse(file_input)
  	@input_distributor = {}
  	json.each do |key,value|
  		@input_distributor[key] = value
  	end
  end

  def region_checker(distributor_name,user_input_region,json_input_region)
  	# p distributor_name
  	# p user_input_region
  	# p json_input_region
  	distributor_right = false
  	included_region = json_input_region["includes"]  	
  	excluded_region = json_input_region["excludes"]
  	user_input_region = user_input_region.split("$")
  	if user_input_region.count == 1
  	  # 1 means it has country      
  	  included_region.each do |region|
        region_split = region.split("$")
  	  	if region_split.count == 1          
          if (region_split[0].downcase.squish == user_input_region[0].downcase.squish)
            distributor_right = true
          end
        elsif region_split.count == 2           
          if (region_split[1].downcase.squish == user_input_region[0].downcase.squish)
            distributor_right = true
          end
        elsif region_split.count == 3
          if (region_split[2].downcase.squish == user_input_region[0].downcase.squish)
            distributor_right = true
          end
  	  	end
  	  end # included region loop end
  	# 2 means it has state and country  
  	elsif user_input_region.count == 2
  	  included_region.each do |region|
        region_split = region.split("$")
        if region_split.count == 1
          # checking for country, if country is there then distributor can distribute
          if (region_split[0].downcase.squish == user_input_region[1].downcase.squish)
            distributor_right = true
          end
        elsif region_split.count == 2
          # checing for both state and country
          if (region_split[0].downcase.squish == user_input_region[0].downcase.squish and region_split[1].downcase.squish == user_input_region[1].downcase.squish )
            distributor_right = true
          end
        elsif region_split.count == 3
          if (region_split[0].downcase.squish == user_input_region[0].downcase.squish and region_split[1].downcase.squish == user_input_region[1].downcase.squish)
            distributor_right = true
          end
        end
      end
    # 3 means it has state, city, country 
  	elsif user_input_region.count == 3
      included_region.each do |region|
        region_split = region.split("$")
        if region.split("$").count == 1          
          # checking for country, if country is there then distributor can distribute
          if (region_split[0].downcase.squish == user_input_region[2].downcase.squish)
            distributor_right = true
          end
        elsif region.split("$").count == 2
          # checing for both state and country
          if (region_split[0].downcase.squish == user_input_region[0].downcase.squish and region_split[1].downcase.squish == user_input_region[1].downcase.squish )
            distributor_right = true
          end
        elsif region.split("$").count == 3
          if (region_split[0].downcase.squish == user_input_region[0].downcase.squish and region_split[1].downcase.squish == user_input_region[1].downcase.squish and region_split[2].downcase.squish == user_input_region[2].downcase.squish)
            distributor_right = true
          end
        end
      end
  	end

    # CHECKING FOR EXCLUDE CONDITION
    if distributor_right      
      if user_input_region.count == 1
        # 1 means it has country      
        excluded_region.each do |region|
          region_split = region.split("$")
          if region_split.count == 1          
            if (region_split[0].downcase.squish == user_input_region[0].downcase.squish)
              distributor_right = false
            end
          end
        end # excluded region loop end
      # 2 means it has state and country  
      elsif user_input_region.count == 2
        excluded_region.each do |region|
          region_split = region.split("$")
          if region_split.count == 1
            # checking for country, if country is there then distributor can distribute
            if (region_split[0].downcase.squish == user_input_region[1].downcase.squish)
              distributor_right = false
            end
          elsif region_split.count == 2
            # checing for both state and country                        
            if (region_split[0].downcase.squish == user_input_region[0].downcase.squish and region_split[1].downcase.squish == user_input_region[1].downcase.squish)
              distributor_right = false
            end          
          end
        end
      # 3 means it has state, city, country 
      elsif user_input_region.count == 3        
        excluded_region.each do |region|  
          region_split = region.split("$")        
          if region.split("$").count == 3                     
            if (region_split[0].downcase.squish == user_input_region[0].downcase.squish and region_split[1].downcase.squish == user_input_region[1].downcase.squish and region_split[2].downcase.squish == user_input_region[2].downcase.squish)
              distributor_right = false
            end
          elsif region.split("$").count == 2
             if (region_split[0].downcase.squish == user_input_region[1].downcase.squish and region_split[1].downcase.squish == user_input_region[2].downcase.squish)
              distributor_right = false
            end 
          end
        end
      end
  	end
  	if distributor_right == false
  		 return "NO"
  	elsif distributor_right == true
  		 return "YES"
  	end  	
  end

  def distributor_checker(distributor_name,user_input_region)
  	main_distributor = @input_distributor[distributor_name]
  	#if main distributor is not present, check for sub distributor
  	if main_distributor == nil  		
  		@input_distributor.each do |key,value|
        if @input_distributor[key]["sub_distributors"]!=nil
    			if @input_distributor[key]["sub_distributors"].keys.include? distributor_name 
    				sub_distributor_include_excludes = @input_distributor[key]["sub_distributors"][distributor_name]
    				#p "Sub-Distributor found"  	
    				p region_checker(distributor_name,user_input_region,sub_distributor_include_excludes)
    			else
    				p "No Main/Sub-Distributor found"
    			end
        end
  		end
  	else
  		#p "Main Distributor"  		
  		p region_checker(distributor_name,user_input_region,main_distributor)
  	end
  end
end

  # INPUT JSON FILE PATH
  json_file_path = "/home/rahul/Work/Qube/input.json"

  # INPUT PERMISSION TO CHECK
  #input_distributor_name = "distributor1"
  #city_state_country_to_check = "Chennai-Tamil Nadu-India"
  puts "Please enter the Distributor Name" 
  input_distributor_name = gets.chomp
  puts "Please enter the Location to check (seperated by $):"
  city_state_country_to_check = gets.chomp

  # CALLING METHOD
  a = Qube.new
  a.input_parser(json_file_path)
  a.distributor_checker(input_distributor_name,city_state_country_to_check)
