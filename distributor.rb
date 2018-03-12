#!/usr/bin/env ruby
# ==================Importing required modules==============
require 'csv'
require_relative 'distributor_helper'

# ==================Including DistributorHelper module ==============
include DistributorHelper
class Distributor
	# ==================declaring class variables ==============
	@@region_names={}
	@@distributors=[]
	# @@country_codes={}
	# @@city_codes={}
	# @@province_codes={}

	# ==================Intializing class object with required data ==============
	def initialize(distributor_name,include_regions,exclude_regions)
		# @name=distributor_name.capitalize   #unused variable
		@include_regions=include_regions||{}
		@exclude_regions=exclude_regions||{}
		@@distributors<<@name unless @@distributors.include?(@name)
	end

	# ==================Method for distributor check==============
	def self.check_distributor_exists(distributor)
		@@distributors.include?(distributor.to_s.capitalize)
	end

	# ==================Method for getting the regions==============
	def get_regions()
		{"include_regions"=>@include_regions,"exclude_regions"=>@exclude_regions}
	end

	# ==================Method for verifying input regions with CSV data==============
	def self.verify_regions(region)
		if !region["Country Name"].empty?
			valid = @@region_names.keys.include?(region["Country Name"].first)
			if !region["Province Name"].empty? && valid
				valid = @@region_names[region["Country Name"].first].keys.include?(region["Province Name"].first)
				if !region["City Name"].empty? && valid
					valid = @@region_names[region["Country Name"].first][region["Province Name"].first].keys.include?(region["City Name"].first)
				end
			end
		end
		valid
	end

	# ==================Method for reading CSV check==============
	#============Can me moved to helper module=====

	def self.read_regions(file_name='cities.csv')
	    CSV.foreach(file_name, headers: true) do |row|
	      row_hash=row.to_hash
	      @@region_names[row_hash['Country Name']] = {} unless @@region_names.key?(row_hash['Country Name'])
	      @@region_names[row_hash['Country Name']][row_hash['Province Name']] = {} unless @@region_names[row_hash['Country Name']].key?(row_hash['Province Name'])	
	      @@region_names[row_hash['Country Name']][row_hash['Province Name']][row_hash['City Name']] = {}
	     	
	     	#=========Unused code================
	      # @@country_codes[row_hash['Country Name']] = row_hash['Country Code'] unless @@country_codes.key?(row_hash['Country Name'])
	      # @@city_codes[row_hash['City Name']] = row_hash['City Code'] unless @@city_codes.key?(row_hash['City Name'])
	      # @@province_codes[row_hash['Province Name']] = row_hash['Province Code'] unless @@province_codes.key?(row_hash['Province Name'])
	    end	
	end

	# ==================Method for distributor permission checks==============
	def check_distributor_permissions(region)
		if !(region["Country Name"].empty?||@include_regions["Country Name"].empty?)
			if @include_regions["Country Name"].include?(region["Country Name"].first) && !@exclude_regions["Country Name"].include?(region["Country Name"].first)
				return true if region["Province Name"].empty?||@include_regions["Province Name"].empty?
				if @include_regions["Province Name"].include?(region["Province Name"].first)
					return true if region["City Name"].empty?||@include_regions["City Name"].include?(region["City Name"].first)||(@include_regions["City Name"].empty?)
				end
			elsif @include_regions["Country Name"].include?(region["Country Name"].first) && @exclude_regions["Country Name"].include?(region["Country Name"].first)
				return true if region["Province Name"].empty?||(@include_regions["Province Name"].empty?&& !@exclude_regions["Province Name"].include?(region["Province Name"].first))
				if @include_regions["Province Name"].include?(region["Province Name"].first) && !@exclude_regions["Province Name"].include?(region["Province Name"].first)
					return true if region["City Name"].empty?||@include_regions["City Name"].include?(region["City Name"].first)||(@include_regions["City Name"].empty?)
				elsif @include_regions["Province Name"].include?(region["Province Name"].first) && @exclude_regions["Province Name"].include?(region["Province Name"].first)				
					return true if region["City Name"].empty?||@include_regions["City Name"].include?(region["City Name"].first)||((@include_regions["City Name"].empty?)&&!@exclude_regions["City Name"].include?(region["City Name"].first))
				end
			end
		end
		return false
	end

end

	# ==================Main Method invoked to run the entire process==============
def main()

	file_name="cities.csv"
	Distributor.read_regions(file_name)

	puts "Welcome to Distributor Contract\n Configure Distributors\n"
	puts "Enter Number of distributors:"
	total_distributor = STDIN.gets.to_s.chomp.to_i
	i=1
	distributors_list = {}
	parent_distributor=""
	distributor_name=""

	while (i<=total_distributor)
		loop do
			puts "Enter Valid Distributor #{i} Name:"
			distributor_name = STDIN.gets.to_s.chomp
			distributor_name_check =Distributor.check_distributor_exists(distributor_name)
			puts "Distributor Name already taken" if distributor_name_check
			break if !distributor_name_check
		end

		loop do
			puts "Enter Valid Distributor #{i} Parent Name:"
			parent_distributor = STDIN.gets.to_s.chomp
			if parent_distributor.empty?
				puts "No Parent distributor"
				break
			end
			parent_distributor_check = Distributor.check_distributor_exists(parent_distributor)
			puts "Parent Distributor Not Found" if !parent_distributor_check
			break if parent_distributor_check
		end

		puts "Enter Number of Include Regions:"
		total_include=STDIN.gets.to_s.chomp.to_i
		if total_include>0
			k=0
			includes=[] 
			while(k<total_include)
				puts "INCLUDE:"
				data=STDIN.gets.to_s.chomp 
				formatted_data = format_regions([data])
				parent_distributor_regions = distributors_list["#{parent_distributor.to_s.upcase}"].get_regions() if !parent_distributor.squish.empty?
				if Distributor.verify_regions(formatted_data) &&(parent_distributor.squish.empty?|| check_include_parent_permissions(formatted_data,parent_distributor_regions["include_regions"],parent_distributor_regions["exclude_regions"]))
					includes<<data
				else
					puts "Invalid Input data:\t Retry with correct input"
					k-=1
				end
				k+=1
			end

			puts "Enter Number of Exclude Regions:"
			total_exclude=STDIN.gets.to_s.chomp.to_i
			if total_exclude>0
				j=0
				excludes=[] 
				while(j<total_exclude)
					puts "EXCLUDE:"
					data=STDIN.gets.to_s.chomp
					formatted_data = format_regions([data])
					if (Distributor.verify_regions(formatted_data) && (parent_distributor.empty?||check_exclude_parent_permissions(format_regions(includes),formatted_data)))
						excludes<<data
					else				
						puts "Invalid Input data:\t Retry with correct input"
						j-=1
					end
					j+=1
				end
			end
		end	

		distributors_list["#{distributor_name.upcase}"]=Distributor.new(distributor_name,format_regions(includes),format_regions(excludes))
		i+=1
	end	
	scanning_distributors(distributors_list)
end

def scanning_distributors(distributors_list={})
	puts "\n\n======Scan for Distributors and their Regions======\n"
	while true
		puts "Enter 'exit' or 'E' at any time to Stop\n"
		puts "Enter distributor name:"
		data=STDIN.gets.to_s.chomp
		break if data=="exit"||data=="E"
		if Distributor.check_distributor_exists(data)
			distributor = distributors_list["#{data.upcase}"]
			puts "Enter region to check:"
			check_region = STDIN.gets.to_s.chomp
			break if check_region=="exit"||check_region=="E"
			formatted_check_region = format_regions([check_region])
			if Distributor.verify_regions(formatted_check_region)
				if distributor.check_distributor_permissions(formatted_check_region)
					puts "YES"
				else
					puts "NO"
				end
			else
				puts"Entered Region is Not Matched With Our data(CSV)\n"
			end
		else
			puts "Distributor does not exist"
		end
	end
end

main()


#===========================sample output data========================


# Welcome to Distributor Contract
#  Configure Distributors
# Enter Number of distributors:
# 2
# Enter Valid Distributor 1 Name:
# distributor1
# Enter Valid Distributor 1 Parent Name:

# No Parent distributor
# Enter Number of Include Regions:
# 1
# INCLUDE:
# india
# Enter Number of Exclude Regions:
# 1
# EXCLUDE:
# china
# Enter Valid Distributor 2 Name:
# distributor2
# Enter Valid Distributor 2 Parent Name:
# distributor1
# Enter Number of Include Regions:
# 1
# INCLUDE:
# tamil nadu-india
# Enter Number of Exclude Regions:
# 2
# EXCLUDE:
# chennai-india
# Invalid Input data:	 Retry with correct input
# EXCLUDE:
# Keelakarai-tamil nadu-india
# EXCLUDE:
# Wellington-tamil nadu-india


# ======Scan for Distributors and their Regions======
# Enter 'exit' or 'E' at any time to Stop
# Enter distributor name:
# distributor 2
# Distributor does not exist
# Enter 'exit' or 'E' at any time to Stop
# Enter distributor name:
# distributor2
# Enter region to check:
# china
# NO
# Enter 'exit' or 'E' at any time to Stop
# Enter distributor name:
# distributor2
# Enter region to check:
# chennai-tamil nadu-india
# YES
# Enter 'exit' or 'E' at any time to Stop
