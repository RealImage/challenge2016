# frozen_string_literal: true

require_relative '../class/distributors'
require_relative '../class/sub_distributors'

# Method to get the province details
# Input - Input type to print to user
# Returns - A array of array of city, state and country.
def input_list(list_type_to_display, check_or_input_flow = 'input_flow', type_of_distributor = 'main')
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
  include_list = input_list('Included')
  exclude_list = input_list('Excluded')
  puts 'Please select if you want to enter sub distributor data if available (y/n):'
  sub_dist_data_available = gets.chomp.downcase
  if sub_dist_data_available == 'y'
    puts 'Please enter name of sub distributor:'
    input_sub_distributor_name = gets.chomp.downcase
    puts 'Please enter included area of sub distributor:'
    input_sub_dist_include_list = input_list('Included', 'subdistributor')
    puts 'Please enter exluded area of sub distributor:'
    input_sub_dist_exclude_list = input_list('Excluded', 'subdistributor')
    @distributors_list <<
      Distributors.new(
        input_dist_name.to_s, include_list, exclude_list, input_sub_distributor_name
      )
    @sub_distributors_list <<
      Subdistributors.new(
        input_sub_distributor_name.to_s, input_sub_dist_include_list, input_sub_dist_exclude_list,
        main_dist_permissible_data(input_dist_name.to_s, 'distributor')
      )
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
  puts 'Please select'
  puts '1 if you want to add sub distributor for a distributor'
  puts '2 if you want to add sub distributor for a sub distributor'

  input_option = gets.chomp
  case input_option.to_i
  when 1
    print_distributor_data
  when 2
    print_sub_distributor_data
  end
  input_main_distributor = gets.chomp.downcase
  puts 'Please enter name of sub distributor:'
  input_sub_distributor_name = gets.chomp.downcase
  puts 'Please enter included area of sub distributor:'
  input_sub_dist_include_list = input_list('Included')
  puts 'Please enter exluded area of sub distributor:'
  input_sub_dist_exclude_list = input_list('Excluded')
  case input_option.to_i
  when 1
    @distributors_list.map do |dist|
      next unless dist.name == input_main_distributor.to_s

      dist.sub_distributors << input_sub_distributor_name.to_s
    end
  when 2
    @sub_distributors_list.map do |dist|
      next unless dist.name == input_main_distributor.to_s

      dist.sub_distributors << input_sub_distributor_name.to_s
    end
  end
  input_option = input_option_to_string(input_option.to_i)
  @sub_distributors_list <<
    Subdistributors.new(
      input_sub_distributor_name.to_s, input_sub_dist_include_list, input_sub_dist_exclude_list,
      main_dist_permissible_data(input_main_distributor.to_s, input_option)
    )
end

# Print distributor data
def print_distributor_data
  if @distributors_list.empty?
    puts 'Distributor data is empty. Please add distributor first.'
    return
  end
  puts 'Please select the name of distributor from the list:'
  @distributors_list.each do |distributor|
    puts distributor.name.downcase
  end
end

# Print Sub distributor data
def print_sub_distributor_data
  if @sub_distributors_list.empty?
    puts 'Sub Distributor data is empty. Please add Sub distributor first.'
    return
  end
  puts 'Please select the name of sub distributor from the list:'
  @sub_distributors_list.each do |distributor|
    puts distributor.name.downcase
  end
  puts '-------------------'
end

# Returns distributor or sub distributor
def input_option_to_string(input)
  if input == 1
    'distributor'
  else
    'sub_distributor'
  end
end

# Gets the allowed data for the main distributer to pass to subdistributor object
def main_dist_permissible_data(main_dist, input_option)
  if input_option == 'distributor'
    @distributors_list.each do |dist|
      return dist.permissible_data if dist.name == main_dist
    end
  else
    input_option == 'sub_distributor'
    @sub_distributors_list.each do |dist|
      return dist.permissible_data if dist.name == main_dist
    end
  end
end
