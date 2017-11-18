require './csv_util.rb'
require './country.rb'
require './state.rb'
require './city.rb'

parsed_csv = CsvUtil.parse_csv('./cities.csv')


def add_data(parsed_csv)
  parsed_csv.each do |row|
    params = row.to_hash
    Country.create(params)
    State.create(params)
    City.create(params)
  end
end

add_data(parsed_csv)