require 'csv'
require './country.rb'

class CsvImportUtil
  class << self
    def parse_cities_csv(file)
      csv_text = File.read(file)
      parsed_csv = CSV.parse(csv_text, headers: true)
      parsed_csv.each do |row|
        Country.new(row.to_hash)
      end
    end
  end
end
