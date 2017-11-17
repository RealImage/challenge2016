require './csv_util.rb'

parsed_csv = CsvUtil.parse_csv('./cities.csv')


def add_cities(parsed_csv)
  parsed_csv.each do |row|
    # add country , state and city
  end
end

add_cities(parsed_csv)
