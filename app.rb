# frozen_string_literal: true

require 'pry'
require 'json'
require_relative 'helpers/csv_helper'
require_relative 'helpers/helpers'

# The MainApp class to run the script.
class MainApp
  attr_reader :distributors_list, :sub_distributors_list

  # Create a hash of the given CSV
  unless File.file?("class/temp.json")
    binding.pry
    File.write("class/temp.json",csv_to_hash('cities.csv').to_json)
  end

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

#  {
#     "countries" => {
#         "india" => {
#             "code" => "in",
#             "province" => {
#                 "Jammu and Kashmir" => {
#                      "code" => "JK",
#                      "cities" => {
#                         "poonch" => {
#                             "code" => "2"
#                             },
#                         "hubli" => {
#                             "code" => "HU"
#                         }
#                     }
#                 }
#             }
#         }
#     }
# }
#     country_hash_data["countries"].merge!({
#         row["Country Name"].to_s => {
#             "code" => row["Country Code"].to_s,
#             "province" => {
#                 row["Province Name"].to_s => {
#                     "code" => row["Province Code"],
#                     "cities" => {
#                         row["City Name"].to_s => {
#                             "code" => row["City Code"].to_s
#                         },
#                     },
#                 },
#             },
#         },
# })
