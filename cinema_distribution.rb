$ALL_REGIONS = []
class Cinema
  attr_accessor :distributors

  def initialize
    self.distributors = []
  end

  require 'csv'
  require 'json'

  def self.all_regions
    file = File.open("cities.csv", "r")
    keys = file.readline().chomp!.split(",")
    result = []

    file.each_line do | line |
      i = 0
      obj = {}
      line.chomp.split(",").each do |v|
        obj[keys[i]] = v
        i += 1
      end
      result.push(obj)
    end
    JSON.parse(result.to_json)
  end

  def add_distributors(total_distributors_count)
    @cond = true
    while @cond
      if total_distributors_count > 0
        @cond = false
        (1..total_distributors_count).each do |i|
          new_distributor =  registration_inputs(i, "Distributor")
          distributors.push(new_distributor)
          Region.set_region(new_distributor)
          assign_sub_distributors(new_distributor)
        end
      else
        puts "Warning! please enter valid number, should be numeric and greater than 0"
        total_distributors_count = Cinema.number_of_distributors
      end
    end
    Login.logout
  end

  def assign_sub_distributors(new_distributor)
    print "\nDo you want to add sub-distributors for #{new_distributor.username.capitalize}? type 'yes' or 'no': "
    yes_or_no = gets.chomp
    if yes_or_no == "yes"
      print "Enter number of sub-distributors: "
      sub_distributors_count = gets.chomp.to_i
      (1..sub_distributors_count).each do |sd|
        sub_distributor = registration_inputs(sd, "SubDistributor")
        sub_distributor.parent = new_distributor
        new_distributor.sub_distributors.push(sub_distributor)
        Region.set_region(sub_distributor)
      end
    end
  end

  def registration_inputs(i, name)
    distributor = Distributor.new
    distributor.type = name
    print "\nEnter #{name} #{i} username: "
    distributor.username = gets.chomp
    print "Enter #{name} #{i} password: "
    distributor.password = gets.chomp
    distributor
  end

  def self.number_of_distributors
    print "\nEnter number of distributors you want to add: "
    gets.chomp.to_i
  end
end

class Login
  def self.authenticate_user(user)
    puts "\nLogin with default Admin credentials!"
    print "Enter username: "
    username = gets.chomp
    print "Enter password: "
    password = gets.chomp

    @login = true
    while @login
      if username.eql?(user.username) && password.eql?(user.password)
        puts "\nLogged in successfully. Welcome #{user.username.capitalize}!"
        @login = false
      else
        puts "Invalid username or password. Please try again."
        self.authenticate_user(user)
      end
    end
  end

  def self.logout
    print "\nTo Logged out type 'logout': "
    logout = gets.chomp
    puts "\nLogged out successfully!" if logout == 'logout'
  end
end

class User
  attr_accessor :username, :password
end

class Admin < User

end

class Distributor < User
  attr_accessor :type, :sub_distributors, :include_regions, :exclude_regions, :parent

  def initialize
    self.include_regions = []
    self.exclude_regions = []
    self.sub_distributors = []
  end

end

class Region
  def self.set_region(new_distributor)
    print "\nEnter number of regions you want to INCLUDE for #{new_distributor.username.capitalize}: "
    regions_count = gets.chomp.to_i
    puts "\nAssign distrubution regions for #{new_distributor.username.capitalize} like below format:"
    puts "eg1: India\neg2: Tamil Nadu-India\neg3: yawal-karnataka-india\neg4: United States"
    assign_regions(regions_count,new_distributor)
    print "\nEnter number of regions you want to EXCLUDE for #{new_distributor.username.capitalize}: "
    exclude_regions_count = gets.chomp.to_i
    assign_regions(exclude_regions_count,new_distributor,false)
  end

  def self.assign_regions(regions_count, new_distributor, include=true)
    (1..regions_count).each do |r|
      print "\nEnter region #{r}: "
      region = gets.chomp
      if !include && new_distributor.include_regions.include?(region)
        puts "You can't exclude included region. Please try with other region."
        assign_regions(regions_count, new_distributor, include)
      end
      result = new_distributor.type == "Distributor" ? Region.search_region(region) :
      (!Region.search_region(region, new_distributor.parent.exclude_regions, false) && ( Region.search_region(region, new_distributor.parent.include_regions, include) && Region.search_region(region) ) )
      if result
        include ? new_distributor.include_regions.push(region) : new_distributor.exclude_regions.push(region)
        puts "\nAssigned #{region} region successfully"
      else
        puts new_distributor.type == "Distributor" ? "\nEntered region is not available or check spell mistake!" : "\nYou don't have distribution rights in this region. Please try inside your region!"
        assign_regions(regions_count, new_distributor, include)
      end
    end
  end

  def self.search_region(keyword, regions = nil, include=nil)
    permit_regions = keyword.split('-')
    regions_count = permit_regions.length
    headers = ["City Name", "Province Name", "Country Name"]
    rows_or_regions = regions || $ALL_REGIONS
    i = regions_count - 1
    row_len = headers.length - 1
    rows_or_regions.each do |row|
      incr = 0
      if regions
        row_len = row.split('-').length - 1
        i = row_len < i ? row_len : i
        return false if !include && row_len != i
      end
      (0..i).each do |e|
        region_key = permit_regions[(regions_count - 1) - e].gsub(/\s+/, "").upcase == (regions ? row.split('-')[row_len - e].gsub(/\s+/, "").upcase : row[headers[row_len-e]].gsub(/\s+/, "").upcase)
        region_key ? incr += 1 : break
      end
      return true if incr == (i + 1)
    end
    false
  end
end


def roll_camera_action()
  admin = Admin.new
  admin.username = "admin"
  admin.password = "cinema"

  puts "\nWelcome to Cinema Distribution World!!"
  puts "Default Admin Account Details:\nusername: #{admin.username}\npassword: #{admin.password}"
  Login.authenticate_user(admin)

  c = Cinema.new
  $ALL_REGIONS = Cinema.all_regions
  total_distributors_count = Cinema.number_of_distributors
  c.add_distributors(total_distributors_count)
end

roll_camera_action()



#================= SAMPLE OUTPUT =================

# coppernine01@coppernine01:~/ruby/challenge2016$ ruby cinema_distribution.rb

# Welcome to Cinema Distribution World!!
# Default Admin Account Details:
# username: admin
# password: cinema

# Login with default Admin credentials!
# Enter username: admin
# Enter password: cinee
# Invalid username or password. Please try again.

# Login with default Admin credentials!
# Enter username: admin
# Enter password: cinema

# Logged in successfully. Welcome Admin!

# Enter number of distributors you want to add: 2

# Enter Distributor 1 username: distb01
# Enter Distributor 1 password: 000000

# Enter number of regions you want to INCLUDE for Distb01: 2

# Assign distrubution regions for Distb01 like below format:
# eg1: India
# eg2: Tamil Nadu-India
# eg3: yawal-karnataka-india
# eg4: United States

# Enter region 1: indiaaa

# Entered region is not available or check spell mistake!

# Enter region 1: india

# Assigned india region successfully

# Enter region 2: united states

# Assigned united states region successfully

# Enter number of regions you want to EXCLUDE for Distb01: 2

# Enter region 1: chennai-tamilnadu-india

# Assigned chennai-tamilnadu-india region successfully

# Enter region 2: california-india

# Entered region is not available or check spell mistake!

# Enter region 2: california-united states

# Assigned california-united states region successfully

# Do you want to add sub-distributors for Distb01? type 'yes' or 'no': yes
# Enter number of sub-distributors: 1

# Enter SubDistributor 1 username: sub-distb01
# Enter SubDistributor 1 password: 000000

# Enter number of regions you want to INCLUDE for Sub-distb01: 1

# Assign distrubution regions for Sub-distb01 like below format:
# eg1: India
# eg2: Tamil Nadu-India
# eg3: yawal-karnataka-india
# eg4: United States

# Enter region 1: Australia

# You don't have distribution rights in this region. Please try inside your region!

# Enter region 1: India

# Assigned India region successfully

# Enter number of regions you want to EXCLUDE for Sub-distb01: 1

# Enter region 1: telangana-india

# Assigned telangana-india region successfully

# Enter Distributor 2 username: distb02
# Enter Distributor 2 password: 000000

# Enter number of regions you want to INCLUDE for Distb02: 1

# Assign distrubution regions for Distb02 like below format:
# eg1: India
# eg2: Tamil Nadu-India
# eg3: yawal-karnataka-india
# eg4: United States

# Enter region 1: Australia

# Assigned Australia region successfully

# Enter number of regions you want to EXCLUDE for Distb02: 1

# Enter region 1: Queensland-australia

# Assigned Queensland-australia region successfully

# Do you want to add sub-distributors for Distb02? type 'yes' or 'no': yes
# Enter number of sub-distributors: 1

# Enter SubDistributor 1 username: sub-distb02
# Enter SubDistributor 1 password: 000000

# Enter number of regions you want to INCLUDE for Sub-distb02: 1

# Assign distrubution regions for Sub-distb02 like below format:
# eg1: India
# eg2: Tamil Nadu-India
# eg3: yawal-karnataka-india
# eg4: United States

# Enter region 1: india

# You don't have distribution rights in this region. Please try inside your region!

# Enter region 1: Queensland-australia

# You don't have distribution rights in this region. Please try inside your region!

# Enter region 1: Australia

# Assigned Australia region successfully

# Enter number of regions you want to EXCLUDE for Sub-distb02: 1

# Enter region 1: South Australia-Australia

# Assigned South Australia-Australia region successfully

# To Logged out type 'logout': logout

# Logged out successfully!
