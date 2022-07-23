# frozen_string_literal: true

require 'csv'

# Generate a hash from the csv
# Input - CSV format file that is of csv class
# Returns - A hash of the CSV
def csv_to_hash(_csvfile)
  country_hash_data = {}
  csv_data = CSV.read('cities.csv', headers: true)
  csv_data.each do |row|
    unless country_hash_data.key?(row['Country Name'].to_s.downcase)
      country_hash_data.store(row['Country Name'].to_s.downcase, {})
    end
    country_hash_data[row['Country Name'].to_s.downcase]
      .merge!('code' => row['Country Code'].to_s.downcase)
    unless country_hash_data[row['Country Name'].to_s.downcase].key?(row['Province Name'].to_s.downcase)
      country_hash_data[row['Country Name'].to_s.downcase]
        .store(row['Province Name'].to_s.downcase, {})
    end
    country_hash_data[row['Country Name'].to_s.downcase][row['Province Name'].to_s.downcase]
      .merge!('code' => row['Province Code'].to_s.downcase)
    unless country_hash_data[row['Country Name'].to_s.downcase][row['Province Name'].to_s.downcase]
           .key?(row['City Name'].to_s.downcase)
      country_hash_data[row['Country Name'].to_s.downcase][row['Province Name'].to_s.downcase]
        .store(row['City Name'].to_s.downcase, {})
    end
    country_hash_data[row['Country Name'].to_s.downcase][row['Province Name'].to_s.downcase][row['City Name'].to_s.downcase]
      .merge!('code' => row['City Code'].to_s.downcase)
  end
  country_hash_data
end
