require './csv_util.rb'
require './country.rb'
require './state.rb'
require './city.rb'
require './distributor.rb'

parsed_csv = CsvUtil.parse_csv('./cities.csv')

def add_data(parsed_csv)
  parsed_csv.each do |row|
    params = row.to_hash
    Country.create(params)
    State.create(params)
    City.create(params)
  end
end

puts "Loading data..."
add_data(parsed_csv)

puts "1. Add a distributor in this format \nDistributor Name,Allowed locations,Unallowed locations,Extends \n(Eg. add Distributor1 IN,IT KA-IN,CENAI-TN-IN)\n"
puts "\n2. Please provide distributor name and location for permission as \n(Eg. permission? Distributor1 BENAU-KA-IN)\nType exit if you are done.\n"
process = true

while process
  input = gets.chomp
  abort("Bye") if input == "exit"
  result = Distributor.perform(input)
  puts "\n#{result}\n"
end