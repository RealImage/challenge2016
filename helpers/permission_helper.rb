# frozen_string_literal: true

require_relative 'input_helper'

# Get input from user to check the permission for the distributor
def input_check_distributor_permission_input
  puts 'Press 1 to check permission for an distributor'
  puts 'Press 2 to check permission for an sub distributor'
  check_type = gets.chomp
  if check_type.to_i == 1
    check_distributor_permission
  else
    check_sub_distributor_permission
  end
end

# Check permission for a distributor
def check_distributor_permission
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
  input_dist_check_list = input_list('check', 'checking_flow')
  if input_dist_check_list.empty?
    puts 'Data not present to check'
    return
  end
  check_permission_for_distributor(input_distributor, input_dist_check_list, 'main')
end

# Check permission for a sub distributor
def check_sub_distributor_permission
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
  input_sub_dist_check_list = input_list('check', 'checking_flow')
  if input_sub_dist_check_list.empty?
    puts 'Data not present to check'
    return
  end
  check_permission_for_distributor(input_sub_distributor, input_sub_dist_check_list, 'sub')
end

# Check if the distributor has permission
def check_permission_for_distributor(distributor_name, check_list, type = 'main')
  if type == 'main'
    @distributors_list.select do |distributor|
      distributor.name == distributor_name
    end.first.check_distributor_data(check_list)
  else
    @sub_distributors_list.select do |sub_distributor|
      sub_distributor.name == distributor_name
    end.first.check_sub_distributor_data(check_list)
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
