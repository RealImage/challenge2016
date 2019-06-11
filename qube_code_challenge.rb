require 'csv'
require "json"
class Distributor
 # exist_region_check is used to check the given region is exist in cities.csv or not
 def self.exist_region_check(str)
  csv_text = File.read('cities.csv')
  csv = CSV.parse(csv_text, :headers => true)
  region = str.split("-")
  case region.length
  when 3
   csv.each do |row|
    if row['City Name'].upcase == region[0].upcase && row['Province Name'].upcase == region[1].upcase && row['Country Name'].upcase == region[2].upcase
     @region_check = true
     break
    else
     @region_check = false
    end
   end
  when 2
   csv.each do |row|
    if row['Province Name'].upcase == region[0].upcase && row['Country Name'].upcase == region[1].upcase
     @region_check = true
     break
    else
     @region_check = false
    end
   end
  when 1
   csv.each do |row|
    if row['Country Name'].upcase == region[0].upcase
     @region_check = true
     break
    else
     @region_check = false
    end
   end
  end
 end
 # include_region function used to insert include region to the distributor
 def self.include_region(include_count,dist_type)
  @sub_include = []
  if include_count
    (1..include_count).each do |we|
     puts "Please enter include region #{we}:"
     region_to_add = gets.chomp.to_s
     exist_region_check(region_to_add)
      if @region_check == true
       if dist_type == "distributor"
        @sub_include<<region_to_add
       else
        file_1 = File.read("qube.json")
        get_file = JSON.parse(file_1)
         if check_distributor_region(get_file,region_to_add,'include') && !check_distributor_region(get_file,region_to_add,'exclude')
          @sub_include<<region_to_add
         else
          p "You don't have permission to include this region."
          get_include_count(dist_type)
         end
       end
      else
       p "Invalid region please check the spell and try again."
       get_include_count(dist_type)
      end
    end
  end
 end
 # exclude_region function used to insert exclude region to the distributor
 def self.exclude_region(exclude_count,include_array, dist_type)
  @sub_exclude = []
  if exclude_count
   (1..exclude_count).each do |we|
    puts "Please enter exclude region #{we}:"
    region_to_add = gets.chomp.to_s
    exist_region_check(region_to_add)
    if @region_check == true && !include_array.include?(region_to_add)
     if dist_type == 'distributor'
      @sub_exclude<<region_to_add
     else
      file_1 = File.read("qube.json")
      get_file = JSON.parse(file_1)
       if !check_distributor_region(get_file,region_to_add, 'exclude')
        @sub_exclude<<region_to_add
       else
        p "You don't have permission to enter this region."
        get_exclude_count(dist_type)
       end
     end
    else
     p "Invalid region please check the spell or you already included this region it can't be exclude."
     get_exclude_count(dist_type)
    end
   end
  end
 end
#add_distributor used to add the distributor name
 def self.add_distributor
  puts "Please enter the distributor name:"
  @distributor_name = gets.chomp.to_s.capitalize
 end

 def self.test
  if add_distributor
   get_include_count('distributor')
   get_exclude_count('distributor')
  list ={list:
             [
                 distributor:{
                     dist_name: @distributor_name,
                     include_dist: {
                         include: @sub_include,
                         exclude: @sub_exclude
                     }
                 }]
  }
   File.open("qube.json", 'w') do |f|
    f.write(JSON.pretty_generate(list))
   end
   check_permission_action
  else
  end
  p "List of permissions for distributors and sub distributors"
  p "********************************************************************"
  pp JSON.parse(File.read("qube.json"))
 end
 # get_include_count used to get the count of include regions.
 def self.get_include_count(type)
  puts "How many region you want to include for #{@distributor_name}:"
  begin
   include_region(Integer(gets.chomp), type)
  rescue
   puts "Please enter number only"
   retry
  end
 end
 # get_exclude_count used to get the count of exnclude regions.
 def self.get_exclude_count(type)
  puts "How many region you want to exclude for #{@distributor_name}:"
  begin
   exclude_region(Integer(gets.chomp), @sub_include, type)
  rescue
   puts "Please enter number only"
   retry
  end
 end
# check_permission_action check the head distributor permitted regions
 def self.check_permission_action
  p "Type YES to know permission for the #{@distributor_name} or type SKIP to go add sub-distributor:"
  gets.chomp.upcase == 'SKIP' ? add_sub_distributor : check_permission
 end
#check_permission display result of the distribution permissions
 def self.check_permission
  p "Enter the region to know:"
  region_permission = gets.chomp
   file_1 = File.read("qube.json")
   get_file = JSON.parse(file_1)
    if check_distributor_region(get_file,region_permission,'include') && !check_distributor_region(get_file,region_permission,'exclude')
     p "Yes #{@distributor_name} have permission"
     check_permission_action
    else
     p "#{@distributor_name} don't have permission"
     check_permission_action
    end
 end
#check_distributor_region check the exist distributor region
 def self.check_distributor_region(list,region,type)
  list.is_a?(Hash) && list.has_key?('list') ?
      check_list = list["list"].map {|s| s['distributor']['include_dist'][type]}.flatten.map(&:upcase) :
      check_list = list[1]['sub_distributor']['include_dist'][type].map(&:upcase)
  include_list = check_list
  region_to_check = region.split("-")
  case region_to_check.length
  when 3
   (include_list.include?("#{region_to_check[0].upcase}-#{region_to_check[1].upcase}-#{region_to_check[2].upcase}") || include_list.include?("#{region_to_check[1].upcase}-#{region_to_check[2].upcase}") || include_list.include?("#{region_to_check[2].upcase}")) ? true : false
  when 2
   (include_list.include?("#{region_to_check[0].upcase}-#{region_to_check[1].upcase}") || include_list.include?("#{region_to_check[1].upcase}")) ? true : false
  when 1
   include_list.include?("#{region_to_check[0].upcase}") ? true : false
  end
 end
#add_sub_distributor add sub-distributor
 def self.add_sub_distributor
  p "Enter how many sub-distributor you want to add for #{@distributor_name}"
  begin
   distributor_count = Integer(gets.chomp)
   sub_distributor(distributor_count)
  rescue
   retry
  end
 end
 #sub_distributor add number of sub-distributors based on given count
 def self.sub_distributor(dis_count)
  if dis_count
   temp_3 =[]
   file_1 = File.read("qube.json")
   get_file = JSON.parse(file_1)
   (1..dis_count).each do |we|
    puts "Please enter the sub-distributor name:"
    sub_distributor_name = gets.chomp
    @distributor_name = sub_distributor_name.capitalize
    get_include_count('sub_distributor')
    get_exclude_count('sub_distributor')
     temp_one = {sub_distributor:{
             dist_name: @distributor_name,
             include_dist: {
                 include: @sub_include,
                 exclude: @sub_exclude
             }
         }}
    temp_3 << temp_one
    kk =get_file['list'] << temp_one
    File.open("qube.json", 'r+') do |f|
     f.write(JSON.pretty_generate(kk))
    end
   end
   end
  end
end
Distributor.test