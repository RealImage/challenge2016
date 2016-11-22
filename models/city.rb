require_relative "state"

class City
  attr_reader :name, :state

  def initialize(city_name, state)
    @name, @state = city_name, state
  end

end