# frozen_string_literal: true

require 'csv'

# Generate a hash from the csv
# Input - CSV format file that is of csv class
# Returns - A hash of the CSV
def csv_to_hash(_csvfile)
  country_hash_data = {
    'countries' => {}
  }
  csv_data = CSV.read('cities.csv', headers: true)
  csv_data.each do |row|
    unless country_hash_data['countries'].key?(row['Country Name'].to_s.downcase)
      country_hash_data['countries']
        .store(row['Country Name'].to_s.downcase, {})
    end
    country_hash_data['countries'][row['Country Name'].to_s.downcase]
      .merge!('code' => row['Country Code'].to_s.downcase)
    unless country_hash_data['countries'][row['Country Name'].to_s.downcase]
           .key?('province')
      country_hash_data['countries'][row['Country Name'].to_s.downcase]
        .store('province', {})
    end
    unless country_hash_data['countries'][row['Country Name'].to_s.downcase]['province']
           .key?(row['Province Name'].to_s.downcase)
      country_hash_data['countries'][row['Country Name'].to_s.downcase]['province']
        .store(row['Province Name'].to_s.downcase, {})
    end
    country_hash_data['countries'][row['Country Name'].to_s.downcase]['province'][row['Province Name'].to_s.downcase]
      .merge!('code' => row['Province Code'].to_s.downcase)
    unless country_hash_data['countries'][row['Country Name'].to_s.downcase]['province'][row['Province Name'].to_s.downcase]
           .key?('cities')
      country_hash_data['countries'][row['Country Name'].to_s.downcase]['province'][row['Province Name'].to_s.downcase]
        .store('cities', {})
    end
    unless country_hash_data['countries'][row['Country Name'].to_s.downcase]['province'][row['Province Name'].to_s.downcase]['cities']
           .key?(row['City Name'].to_s.downcase)
      country_hash_data['countries'][row['Country Name'].to_s.downcase]['province'][row['Province Name'].to_s.downcase]['cities']
        .store(row['City Name'].to_s.downcase, {})
    end
    country_hash_data['countries'][row['Country Name'].to_s.downcase]['province'][row['Province Name'].to_s.downcase]['cities'][row['City Name'].to_s.downcase]
      .merge!('code' => row['City Code'].to_s.downcase)
  end
  binding.pry
  country_hash_data
end
