require 'csv'

class CsvUtil

  def self.read_file(file_name)
    File.read(file_name)
  end

  def self.parse_csv(file_name)
    csv_text = self.read_file(file_name)
    parsed_csv = CSV.parse(csv_text, headers: true)
  end

end