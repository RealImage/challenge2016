class Country
  attr_accessor :code, :name

  def initialize(params={})
    @code = params["Country Code"]
    @name = params["Country Name"]
  end

  def self.create params
    return if params["Country Code"].nil? or params["Country Code"].empty?
    country = find_by_code(params["Country Code"])
    self.new(params) if country.nil?
  end

  def self.all
    ObjectSpace.each_object(self)
  end

  def self.find_by_code(code)
    all.to_a.detect { |obj| obj.code == code }
  end

  def included?(region)
    country_code, state_code, city_code = region
    self.code == country_code
  end
end