require 'csv'
require './distributor.rb'

class CsvUtil

	def self.parse_csv(file_name)
		csv_text = File.read(file_name)
		parsed_csv = CSV.parse(csv_text, headers: true)
		parsed_csv.each do |row|
			Distributor.new(row.to_hash)
		end
	end
end