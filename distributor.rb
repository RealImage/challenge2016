class Distributor
  attr_accessor :name, :allowed_regions, :unallowed_regions, :extends

  def initialize(params)
    @name = params.first
    @allowed_regions = Distributor.set_region(params[1] || "")
    params[2] = "nil" if (params[2].nil? || params[2] == "NULL")
    @unallowed_regions = Distributor.set_region(params[2])
    @extends = Distributor.find(params[3])
  end

  def self.perform(input)
    return if input.nil? or input.empty?
    input = input.split(" ")
    method_name = input.shift&.strip
    return "Invalid Command" if (method_name.nil? || !(self.respond_to? method_name))
    self.send method_name, input
  end

  def self.add(details)
    valid = validate(details)
    return valid if valid != true 
    self.new(details)
    "success"
  end

  def self.permission?(input)
    distributor_name, region = input
    distributor = find(distributor_name)
    return "Distributor not found" unless distributor 
    return "Invalid region" unless is_valid?(region)
    parent_permissions = distributor.get_all_extends.all?{ |d| d.check_permission(region) }
    return display_result(parent_permissions) unless parent_permissions
    permission = distributor.check_permission(region)
    display_result(permission)
  end

  def self.find distributor_name
    all.to_a.detect { |obj| obj.name.downcase == distributor_name&.downcase }
  end

  def self.all
    ObjectSpace.each_object(self)
  end

  def self.is_valid?(region)
    return false if region.nil?
    !find_region(region).nil?
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

  def self.validate(details)
    return "Distributor exists already" unless find(details[0]).nil?
    return "Include should be given" if details[1].nil? or details[1].empty?
    return "Invalid include for distributor" unless validate_regions(details[1])
    return "Invalid exclude for distributor" if !(details[2].nil? or details[2].empty? or details[2]=="NULL") && !validate_regions(details[2])
    return "Cannot Inherit from #{details[3]}" if !(details[3].nil? or details[3].empty?) && !validate_extends(details[3], details[1], details[2])
    true
  end

  def self.validate_regions(regions)
    regions = regions.split(",")
    regions.all?{ |region| self.is_valid?(region) }
  end

  def self.validate_extends(distributor_name, allowed, unallowed)
    (permission?([distributor_name,allowed]) == "Yes") ? true : false
  end

  def self.set_region(regions)
    regions = regions.split(",")
    return regions.collect{ |region| self.find_region(region) }.compact
  end

  def self.display_result(permission=false)
    permission ? "Yes" : "No"
  end

  def get_all_extends
    distributor = self.extends
    distributor_levels = [self, distributor]
    until distributor == nil
      distributor = distributor.extends
      distributor_levels << distributor if !distributor.nil?
    end
    distributor_levels.compact
  end

  def check_permission(region)
    in_allowed = self.check_region("allowed_regions", region)
    in_unallowed = self.check_region("unallowed_regions", region)
    in_allowed && !in_unallowed
  end

  def check_region(attribute,region)
    region = region.split("-").reverse
    self.send(attribute.to_sym).any? { |r| r.included?(region) }
  end

end