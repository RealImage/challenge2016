require_relative '../../constants'

class Distributor
  include Constants

  attr_accessor :name, :included, :excluded, :inherits

  def initialize(options)
    @name = options[:name]
    @included = {}
    @excluded = {}
    @inherits = options[:inherits]
  end

  def details
    {
      name: name,
      included: included,
      excluded: excluded,
      inherits: inherits
    }
  end

  def add_included(type, code)
    included[type] = [] if included[type].nil?

    included[type] << code
  end

  def add_excluded(type, code)
    excluded[type] = [] if excluded[type].nil?

    excluded[type] << code
  end

  def included_cities
    included[:city] || []
  end

  def included_states
    included[:state] || []
  end

  def included_countries
    included[:country] || []
  end

  def exclude_cities
    excluded[:city] || []
  end

  def exclude_states
    excluded[:state] || []
  end

  def exclude_countries
    excluded[:country] || []
  end
end
