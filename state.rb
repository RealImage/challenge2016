class State
  attr_accessor :code, :name, :country_code

  def initialize params={}
    @code = params["Province Code"]
    @name = params["Province Name"]
    @country_code = params["Country Code"]
  end

  def self.create params
    return if params["Country Code"].nil? or params["Country Code"].empty? or params["Province Code"].nil? or params["Province Code"].empty?
    state = find_by_code(params["Province Code"], params["Country Code"])
    self.new(params) if state.nil?
  end

  def self.all
    ObjectSpace.each_object(self)
  end

  def self.find_by_code(province_code, country_code=nil)
    all.to_a.detect { |state| state.code == province_code && state.country_code == country_code }
  end

  def included?(region)
    country_code, state_code, city_code = region
    self.code == state_code && self.country_code == country_code
  end
end