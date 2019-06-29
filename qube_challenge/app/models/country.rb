require 'csv'
class Country < ActiveRecord::Base
# require 'active_record'
# require 'activerecord-import'

#   def self.import
#     path = Rails.root.join("public", "cities_test.csv")
#     csv_text = File.read(path)
    
#     city_codes = []
#     province_codes = []
#     country_codes = []

#     country_records = []
#     province_records = []
#     city_records = []

#     csv = CSV.parse(csv_text, :headers => true)
#     csv.each do |row|
#       unless country_codes.include? row["Country Code"]
#         country_codes << row["Country Code"]
#         country_records << Country.new(code: row["Country Code"], name: row["Country Name"])
#       end
#       unless province_codes.include? row["Province Code"]
#         province_codes << row["Province Code"]
#         province_records << Province.new(code: row["Province Code"], name: row["Province Name"])
#       end
#       unless city_codes.include? row["City Code"]
#         city_codes << row["City Code"]
#         city_records << City.new(code: row["City Code"], name: row["City Name"])
#       end
#     end
#     Country.import country_records
#     Province.import province_records
#     City.import city_records
#   end
end
