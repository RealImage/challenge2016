class Distributor
  attr_accessor :id, :name, :allowed_locations, :unallowed_locations, :extends

  @@list = {}

  def initialize(params)
    @name = Distributor.set_value(params["Distributor Name"])
    @allowed_locations = set_location(params["Allowed locations"])
    @unallowed_locations = set_location(params["Unallowed locations"])
    @extends = set_extends(params["Extends"])
    add_to_list
  end

  def self.perform(input)
    distributor_name, location = self.process_input(input)
    distributor = self.find_by_name(distributor_name)
    return "#{distributor_name} is not found" if distributor.nil?
    return distributor.has_permission?(location)
  end

  def self.process_input(input)
    input = input.split(" ")
    [Distributor.set_value(input.first), Distributor.set_value(input.last)]
  end

  def self.find_by_name(name)
    @@list[name]
  end

  def self.set_value(value)
    return "" if value.nil?
    value.downcase 
  end

  def self.list
    @@list.each do |k, v|
      puts "#{k.capitalize} ->  Allowed: #{v.allowed_locations} , Unallowed: #{v.unallowed_locations}\n"
    end
  end

  def set_extends(extends)
    return [] if extends.nil?
    extends = extends.split(",").collect do |distributor_name| 
      Distributor.find_by_name(distributor_name.downcase)
    end.compact
    extends
  end

  def set_location(locations)
    return [] if (locations.nil? || locations.empty?)
    locations = locations.split(",")
    return locations.collect{ |location| parse_location(location) }
  end

  def parse_location(location)
    location = location.split("-").reverse
    {country: Distributor.set_value(location[0]), state: Distributor.set_value(location[1]), city: Distributor.set_value(location[2])}
  end

  def add_to_list
    @@list[self.name] = self
  end

  def has_permission?(location)
    parsed_location = parse_location(location)
    parent_permissions = self.extends.all?{ |distributor| distributor.check_permission(parsed_location) }
    return display_result(parent_permissions) unless parent_permissions
    permission = self.check_permission(parsed_location)
    display_result(permission)
  end
  
  def check_permission(location)
    is_allowed = self.eval_location("allowed_locations", location)
    is_unallowed = self.eval_location("unallowed_locations", location)
    permission = is_allowed && !is_unallowed
  end

  def eval_location(attribute, location)
    locations = self.send(attribute.to_sym)
    return true if locations.detect{ |i| i == {country: location[:country], state: location[:state], city: location[:city]}}
    return true if locations.detect{ |i| i == {country: location[:country], state: location[:state], city: ""}}
    return true if locations.detect{ |i| i == {country: location[:country], state: "", city: ""}}
    return false
  end

  def display_result(result)
    result ? "Yes" : "No"
  end

end