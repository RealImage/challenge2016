require 'csv'
require './distributor.rb'
require './country.rb'

class CsvImportUtil
  class << self
    def parse_cities_csv(file)
      parsed_csv = parse_csv_file(file)
      parsed_csv.each do |row|
      Country.new(row.to_hash)
      end
    end

    def parse_distributors_csv(file)
      parsed_csv = parse_csv_file(file)
      parsed_csv.each do |row|
      Distributor.new(row.to_hash)
      end
    end

    def parse_csv_file(file)
      csv_text = File.read(file)
      CSV.parse(csv_text, headers: true)
    end
  end
end
