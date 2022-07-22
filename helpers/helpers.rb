# frozen_string_literal: true

require 'pry'
require_relative '../class/distributors'
require_relative '../class/sub_distributors'

# Method to get the province details
# Input - Input type to print to user
# Returns - A array of array of city, state and country.
def get_input_list(list_type_to_display, check_or_input_flow = 'input_flow', type_of_distributor = 'main')
  input_data = []
  loop do
    puts "Please enter the #{list_type_to_display} regions:"
    puts 'Please enter COUNTRY or press enter if its empty:'
    input_country = gets.chomp.downcase
    puts 'Please enter PROVINCE or press enter if its empty:'
    input_state = gets.chomp.downcase
    if input_state.empty? && type_of_distributor == 'main'
      input_city = ''
    else
      puts 'Please enter CITY separated or press enter if its empty:'
      input_city = gets.chomp.downcase
    end
    input_data << {
      'countries' => input_country,
      'province' => input_state,
      'cities' => input_city
    }.delete_if { |_k, v| v.empty? }
    if check_or_input_flow == 'input_flow'
      puts 'Please enter y/n if you want to add more data(y/n):'
      input_option = gets.chomp.downcase
      break if input_option.to_s == 'n'
    else
      break
    end
  end
  input_data
end

# Get the information about the distributor/sub distributor and store the objects
def input_distributor_info
  puts 'Press 1 to enter distributor data:'
  puts 'Press 2 to enter sub distributor data:'
  distributor_input = gets.chomp
  input_distributor_data if distributor_input.to_i == 1
  input_sub_distributor_data if distributor_input.to_i == 2
end

# Gets distributor data and stores the object
def input_distributor_data
  puts 'Please enter distributor name:'
  input_dist_name = gets.chomp.downcase
  include_list = get_input_list('Include')
  exclude_list = get_input_list('Exclude')
  puts 'Please select if you want to enter sub distributor data if available (y/n):'
  sub_dist_data_available = gets.chomp.downcase
  if sub_dist_data_available == 'y'
    puts 'Please enter name of sub distributor:'
    input_sub_distributor_name = gets.chomp.downcase
    puts 'Please enter included area of sub distributor:'
    input_sub_dist_include_list = get_input_list('Include', 'subdistributor')
    puts 'Please enter exluded area of sub distributor:'
    input_sub_dist_exclude_list = get_input_list('Exclude', 'subdistributor')
    @distributors_list <<
      Distributors.new(
        input_dist_name.to_s, include_list, exclude_list, input_sub_distributor_name
      )
    @sub_distributors_list <<
      Subdistributors.new(
        input_sub_distributor_name.to_s, input_sub_dist_include_list, input_sub_dist_exclude_list,
        main_dist_permissible_data(input_dist_name.to_s))
  else
    input_sub_distributor_name = nil
    @distributors_list <<
      Distributors.new(
        input_dist_name.to_s, include_list, exclude_list, input_sub_distributor_name
      )
  end
end

# Gets subdistributor data and stores the object
def input_sub_distributor_data
  if @distributors_list.empty?
    puts 'Distributor data is empty. Please add distributor first.'
    return
  end
  puts 'Please select the name of distributor from the list:'
  @distributors_list.each do |distributor|
    puts distributor.name.downcase
  end
  puts '-------------------'
  input_main_distributor = gets.chomp.downcase
  puts 'Please enter name of sub distributor:'
  input_sub_distributor_name = gets.chomp.downcase
  puts 'Please enter included area of sub distributor:'
  input_sub_dist_include_list = get_input_list('Included')
  puts 'Please enter exluded area of sub distributor:'
  input_sub_dist_exclude_list = get_input_list('Excluded')
  @distributors_list.map do |dist|
    next unless dist.name == input_main_distributor.to_s

    dist.sub_distributors << input_sub_distributor_name.to_s
    Subdistributors.new(
      input_sub_distributor_name.to_s, input_sub_dist_include_list, input_sub_dist_exclude_list,
      main_dist_permissible_data(input_main_distributor.to_s))
  end
end

def main_dist_permissible_data(main_dist)
  @distributors_list.each do |dist|
    if dist.name == main_dist
      return dist.permissible_data
    end
  end
end

# Get input from user to check the permission for the distributor
def input_check_distributor_permission_input
  puts 'Press 1 to check permission for an distributor'
  puts 'Press 2 to check permission for an sub distributor'
  check_type = gets.chomp
  if check_type.to_i == 1
    if @distributors_list.empty?
      puts 'The distributors list is empty. Please enter the data first and try again'
      puts
      puts
      return
    end
    puts 'Please select the distributor from the list'
    @distributors_list.each do |distributor|
      puts distributor.name
    end
    puts '-------------------'
    input_distributor = gets.chomp.downcase
    if @distributors_list.select do |distributor|
         distributor.name == input_distributor
       end.empty?
      puts 'Please enter valid data from list and try again'
      return
    end
    input_dist_check_list = get_input_list('check', 'checking_flow')
    if input_dist_check_list.empty?
      puts 'Data not present to check'
      return
    end
    check_permission_for_distributor(input_distributor, input_dist_check_list, type = 'main')
  else
    if @sub_distributors_list.empty?
      puts 'The sub distributor list is empty. Please enter the data first and try again'
      puts
      puts
      return
    end
    puts 'Please select the sub distributor from the list'
    @sub_distributors_list.each do |sub_distributor|
      puts sub_distributor.name
    end
    puts '-------------------'
    input_sub_distributor = gets.chomp
    if @sub_distributors_list.select do |distributor|
         distributor.name == input_sub_distributor
       end.empty?
      puts 'Please enter valid data from list and try again'
      return
    end
    input_sub_dist_check_list = get_input_list('check', 'checking_flow')
    if input_sub_dist_check_list.empty?
      puts 'Data not present to check'
      return
    end
    check_permission_for_distributor(input_sub_distributor, input_sub_dist_check_list, type = 'sub')
  end
end

# Check if the distributor has permission
def check_permission_for_distributor(distributor_name, check_list, type)
  binding.pry
  if type == 'main'
    @distributors_list.select do |distributor|
      distributor.name == distributor_name
    end.first.check_distributor_data(check_list)
  else
    @sub_distributors_list.select do |sub_distributor|
      sub_distributor.name == distributor_name
    end.first.check_distributor_data(check_list)
  end
end

# Display message if the distributor is authorised or not.
def display_message(region_to_display, authorised)
  if authorised
    puts "Congratulations!! The distributor can distribute in the region #{region_to_display}"
  else
    puts "The distributor is not authorised to distribute in the region #{region_to_display}"
  end
end

# # Checks if the distributor object has permission or not.
# def check_distributor_data(check_list)
#   binding.pry
#   region_to_check = check_list.first
#   if permissible_data[region_to_check['countries']].nil?
#     display_message(region_to_check['countries'], false)
#   elsif region_to_check['province'] && permissible_data[region_to_check['countries']][region_to_check['province']].nil?
#     display_message(region_to_check['province'], false)
#   elsif region_to_check['cities'] && permissible_data[region_to_check['countries']][region_to_check['province']][region_to_check['cities']].nil?
#     display_message(region_to_check['cities'], false)
#   else
#     display_message(region_to_check['countries'], true)
#   end
#   nil
# end
