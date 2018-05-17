require './csv_import_util.rb'
CsvImportUtil.parse_cities_csv('./cities.csv')
CsvImportUtil.parse_distributors_csv('./distributors.csv')

continue = 'y'
while (continue == 'y')
  puts 'Input your Distributor name'
  distributor = gets.chomp
  puts 'Input your Distributor location in the following format city-province-country (eg: UDUPI-KARNATAKA-INDIA)'
  region = gets.chomp
  permission, @errors = Distributor.perform(distributor, region)
  puts permission
  puts "\n #{@errors.first}" if @errors && !@errors.empty?
  puts 'continue? y : any other key'
  continue = gets.chomp
end
