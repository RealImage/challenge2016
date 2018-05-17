require 'pry'
class Distributor
  attr_accessor :id, :name, :allowed_locations, :unallowed_locations, :parents

  @@list = {}

  def initialize(params)
    @name = self.class.case_insentive(params["Name"])
    @allowed_locations = assign_location(params["Allowed locations"])
    @unallowed_locations = assign_location(params["Unallowed locations"])
    @parents = assign_parents(params["Parent name"])
    add_to_list
  end

  class << self
    def perform(distributor_name, location)
      distributor_name = case_insentive(distributor_name)
      location = case_insentive(location)
      distributor = self.find_by_name(distributor_name)
      return "Distributor #{distributor_name} is not found" if distributor.nil?
      distributor.validate_permission(location)
    end

    def find_by_name(name)
      @@list[name]
    end

    def case_insentive(value)
      return "" if value.nil?
      value.downcase.strip
    end
  end

  def assign_parents(parents)
    return [] if parents.nil?
    parents = parents.split(",").collect do |distributor_name| 
      Distributor.find_by_name(distributor_name.downcase)
    end.compact
    parents
  end

  def assign_location(locations)
    return [] if (locations.nil? || locations.empty?)
    locations = locations.split(",")
    locations.collect{ |location| parse_location(location) }
  end

  def parse_location(location)
    location = location.split("-").reverse
    { country: self.class.case_insentive(location[0]), state: self.class.case_insentive(location[1]),
      city: self.class.case_insentive(location[2]) }
  end

  def add_to_list
    @@list[self.name] = self
  end

  def validate_input_location(location)
    country = location[:country]
    province = location[:state]
    city = location[:city]
    @errors = []
    country_details = $country_list.map {|keys, value| value if keys.include? country}.compact.first
    @errors << "#{country} not found in coutry list. So answer is NO" unless country_details
    if province && !province.empty? && @errors.empty?
      validate_province_and_cities(country_details.province_and_cities, province, city)
    end
    @errors
  end

  def validate_province_and_cities(province_details, province, city)
    city_details = province_details.map {|keys, value| value if keys.include? province}.compact.first
    @errors << "#{province} is not valid in the provided country. So answer is NO" unless city_details
    if city && !city.empty? && @errors.empty?
      @errors << "#{city} is not in our list. So answer is NO" unless city_details[:city_list].include? city
    end
  end

  def validate_permission(location)
    parsed_location = parse_location(location)
    validate_input_location(parsed_location)
    parent_permissions = self.parents.all?{ |distributor| distributor.permited?(parsed_location) }
    return generate_result(parent_permissions) unless parent_permissions
    permission = permited?(parsed_location)
    generate_result(permission)
  end

  def permited?(location)
    validate_location('allowed_locations', location) && !validate_location('unallowed_locations', location)
  end

  def validate_location(attribute, location)
    locations = self.send(attribute.to_sym)
    return true if locations.detect{ |i| i == {country: location[:country], state: location[:state],
      city: location[:city]}}
    return true if locations.detect{ |i| i == {country: location[:country], state: location[:state],
      city: ''}}
    return true if locations.detect{ |i| i == {country: location[:country], state: '', city: ''}}
    false
  end

  def generate_result(result)
    permited = (result && @errors.empty?) ? 'Yes' : 'No'
    [permited, @errors]
  end

end
