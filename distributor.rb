class Distributor
  attr_accessor :name, :allowed_regions, :unallowed_regions, :extends

  def initialize(params)
    @name, @allowed_regions, @allowed_regions, @extends = params
    @allowed_regions = Distributor.set_region(params[1] || "")
    @unallowed_regions = Distributor.set_region(params[2] || "")
  end

  def self.perform(input)
    return if input.nil? or input.empty?
    input = input.split(" ")
    method_name = input.shift.strip
    return "Invalid Command" unless self.respond_to? method_name
    self.send method_name, input.strip
  end

  def self.add(details)
    valid = validate(details)
    return valid if valid != true 
    self.new(details)
    "success"
  end

  def self.validate(details)
    return "Include should be given" if details[1].nil? or details[1].empty?
    return "Invalid include for distributor" unless validate_regions(details[1])
    return "Invalid exclude for distributor" if !(details[2].nil? or details[2].empty?) && !validate_regions(details[2])
    true
  end

  def self.permission?(input)
    distributor_name, region = input
    distributor = find(distributor_name)
    return "Distributor not found" unless distributor 
    return "Invalid region" unless is_valid?(region)
    region = region.split("-").reverse
    in_allowed = distributor.allowed_regions.any? { |r| r.included?(region) }
    in_unallowed = distributor.unallowed_regions.any? { |r| r.included?(region) }
    permitted = in_allowed && !in_unallowed 
    (permitted ? "Yes" : "No")
  end

  def self.find distributor_name
    all.to_a.detect { |obj| obj.name.downcase == distributor_name&.downcase }
  end

  def self.all
    ObjectSpace.each_object(self)
  end

  def self.validate_regions(regions)
    regions = regions.split(",")
    regions.all?{ |region| self.is_valid?(region) }
  end

  def self.is_valid?(region)
    !find_region(region).nil?
  end

  def self.set_region(regions)
    regions = regions.split(",")
    return regions.collect{ |region| self.find_region(region) }
  end

  def self.find_region(region)
    region = region.split("-")
    level = region.size
    if level == 3
      City.find_by_code(region[0], region[1], region[2])
    elsif level == 2
      State.find_by_code(region[0], region[1])
    elsif level == 1
      Country.find_by_code(region[0])
    end
  end
end