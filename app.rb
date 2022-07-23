# frozen_string_literal: true

require 'json'
require_relative 'helpers/csv_helper'
require_relative 'helpers/input_helper'
require_relative 'helpers/permission_helper'

# The MainApp class to run the script.
class MainApp
  attr_reader :distributors_list, :sub_distributors_list

  # Create a hash of the given CSV
  File.write('class/cities.json', csv_to_hash('cities.csv').to_json) unless File.file?('class/cities.json')

  @distributors_list = []

  @sub_distributors_list = []

  loop do
    puts 'Press 1 to enter distributor data'
    puts 'Press 2 to check permission for entered distributors'
    puts 'Press 3 to Exit'
    input = gets.chomp
    case input.to_i
    when 1
      input_distributor_info
    when 2
      input_check_distributor_permission_input
    end
    break if input.to_i == 3
  end
end
