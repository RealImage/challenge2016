# This file should contain all the record creation needed to seed the database with its default values.
# The data can then be loaded with the rake db:seed (or created alongside the db with db:setup).
#
# Examples:
#
#   cities = City.create([{ name: 'Chicago' }, { name: 'Copenhagen' }])
#   Mayor.create(name: 'Emanuel', city: cities.first)

# require 'active_record'
# require 'activerecord-import'
# require 'csv'

# path = Rails.root.join("public", "cities.csv")
# csv_text = File.read(path)

# city_codes = []
# province_codes = []
# country_codes = []

# country_records = []
# province_records = []
# city_records = []

# csv = CSV.parse(csv_text, :headers => true)
# csv.each do |row|
#   unless country_codes.include? row["Country Code"]
#     country_codes << row["Country Code"]
#     country_records << Country.new(code: row["Country Code"], name: row["Country Name"])
#   end
#   unless province_codes.include? row["Province Code"]
#     province_codes << row["Province Code"]
#     province_records << Province.new(code: row["Province Code"], name: row["Province Name"], country_code: row["Country Code"])
#   end
#   unless city_codes.include? row["City Code"]
#     city_codes << row["City Code"]
#     city_records << City.new(code: row["City Code"], name: row["City Name"], province_code: row["Province Code"])
#   end
# end

# Country.import(country_records)
# Province.import(province_records)
# City.import(city_records)


dis1 = Distributor.create!(name: "distributor1", parent_id: nil)
dis1.reload
dis2 = Distributor.create!(name: "distributor2", parent_id: dis1.id)
dis2.reload
dis3 = Distributor.create!(name: "distributor3", parent_id: dis2.id)


dis_area1 = DistributorArea.create!(distributor_id: dis1.id, country_code: "IN", is_included: "true")
dis_area1 = DistributorArea.create!(distributor_id: dis1.id, country_code: "US", is_included: "true")
dis_area1 = DistributorArea.create!(distributor_id: dis1.id, country_code: "IN", province_code: "KA", is_included: "false")
dis_area2 = DistributorArea.create!(distributor_id: dis2.id, country_code: "IN", province_code: "TN", is_included: "true")









