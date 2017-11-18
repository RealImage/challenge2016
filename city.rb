class City
  attr_accessor :code, :name, :state_code, :country_code

  def initialize params={}
    @code = params["City Code"]
    @name = params["City Name"]
    @state_code = params["Province Code"]
    @country_code = params["Country Code"]
  end

  def self.create params
    return if params["City Code"].nil? or params["City Code"].empty? or params["Province Code"].nil? or params["Province Code"].empty? or params["Country Code"].nil? or params["Country Code"].empty?
    city = find_by_code(params["City Code"], params["Province Code"], params["Country Code"])
    self.new(params) if city.nil?
  end

  def self.all
    ObjectSpace.each_object(self)
  end

  def self.find_by_code(city_code, province_code=nil, country_code=nil)
    all.to_a.detect { |city| city.code == city_code && city.state_code == province_code && city.country_code == country_code }
  end
end