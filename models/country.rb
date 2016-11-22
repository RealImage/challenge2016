class Country
  attr_reader :name, :states

  def initialize(country_name)
    @name = country_name
    @states = {}
  end

  def add_state(state)
    @states[state.name] ||= state
  end
end