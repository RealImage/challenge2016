require './csv_util.rb'
require './country.rb'
require './state.rb'
require './city.rb'
require './distributor.rb'

parsed_csv = CsvUtil.parse_csv('./regions.csv')

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

puts "1. Add a distributor in this format \nadd Distributor Name,Allowed regions,Unallowed regions,Extends \n\nExamples:\nadd Distributor1 IN,IT KA-IN,CENAI-TN-IN)\nadd Distributor2 IN TN-IN Distributor1\nadd Distributor3 IN NULL Distributor2\n\n"
puts "\n2. Please provide distributor name and location for permission as \npermission? Distributor1 REGION\n\nExamples:\npermission? Distributor1 BENAU-KA-IN\n\nType exit if you are done.\n\n"
process = true

while process
  input = gets.chomp
  abort("Bye") if input == "exit"
  result = Distributor.perform(input)
  puts "#{result}\n\n"
end