require './csv_util.rb'

CsvUtil.parse_csv('./input.csv')

puts "Please provide distributor name and location as (Eg. Distributor1 BANGALORE-KARNATAKA-INDIA)\nEnter exit if you are done.\n"
process = true

while process
  process = false
  puts "Input your Distributor name and location:"
  input = gets.chomp
  unless input == "exit"
    result = Distributor.perform(input)
    process = true
    puts "Has Permission? - #{result}"
  end
end