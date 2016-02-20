require_relative "country"

class State
  attr_reader :name, :country

  def initialize(name, country)
    @name = name
    @country = country
    @cities = {}
  end

  def add_city(city)
    @cities[city.name] ||= city
  end
end