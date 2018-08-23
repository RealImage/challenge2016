require_relative 'login'
require_relative 'user'
require_relative 'admin'
require_relative 'distributor'
require_relative 'region'
require 'csv'
require 'json'

$ALL_REGIONS = []
class Cinema
  attr_accessor :distributors

  def initialize
    self.distributors = []
  end

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
    result
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
