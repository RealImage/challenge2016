require_relative '../../constants'
require_relative '../distributor/distributor'

class DistributorManager
  include Constants
  attr_accessor :distributors

  def initialize
    # Contains the distributor name -> distributor Object map
    @distributors = {}
  end

  def print_status
    distributors.map { |_name, distributor| distributor.details }
  end

  def create_distributor(options = {})
    distributors[options[:name]] = Distributor.new(options)
  end

  def add_included(name, type, region)
    distributor = find_distributor(name)

    distributor.add_included(type, region)
  end

  def add_excluded(name, type, region)
    distributor = find_distributor(name)

    distributor.add_excluded(type, region)
  end

  def valid_parent?(name, type, regions)
    dist_object = find_distributor(name)
    return true if dist_object.inherits.nil?

    permitted_to_distribute?(dist_object.inherits, type, regions)
  end

  def permitted_to_distribute?(name, type, region_code)
    distributor = find_distributor(name)
    parent_distributor = distributor.inherits

    distributor_permitted = send("check_for_#{type}_exclusion", distributor, region_code) && send("check_for_#{type}_inclusion", distributor, region_code)
    if parent_distributor.nil?
      distributor_permitted
    else
      parent_distributor_permitted = permitted_to_distribute?(parent_distributor, type, region_code)
      distributor_permitted && parent_distributor_permitted
    end
  end

  private

  def check_for_city_exclusion(distributor, region_code)
    !distributor.exclude_cities.include?(region_code[:city]) && check_for_state_exclusion(distributor, region_code)
  end

  def check_for_state_exclusion(distributor, region_code)
    !distributor.exclude_states.include?(region_code[:state]) && check_for_country_exclusion(distributor, region_code)
  end

  def check_for_country_exclusion(distributor, region_code)
    !distributor.exclude_countries.include?(region_code[:country])
  end

  def check_for_city_inclusion(distributor, region_code)
    distributor.included_cities.include?(region_code[:city]) || check_for_state_inclusion(distributor, region_code)
  end

  def check_for_state_inclusion(distributor, region_code)
    distributor.included_states.include?(region_code[:state]) || check_for_country_inclusion(distributor, region_code)
  end

  def check_for_country_inclusion(distributor, region_code)
    distributor.included_countries.include?(region_code[:country])
  end

  def find_distributor(name)
    distributors[name]
  end
end
