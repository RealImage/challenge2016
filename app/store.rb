require 'csv'

class Store
  attr_accessor :path, :data

  def initialize(path)
    @path = path
    @data = {}
  end

  def load_data
    CSV.foreach(path, :headers => true) do |row|
      city, province, country = row[0..2]
      area = [city, province, country].join("::")
      data[area] = row
    end
  end
end
