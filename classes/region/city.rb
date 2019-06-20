class City
  def initialize(options = {})
    @code = options[:code]
    @name = options[:name]
    @state = options[:state]
    @country = options[:country]
  end
end
