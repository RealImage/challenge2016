require_relative "city"
require_relative "state"
require_relative "country"

class Distributor
  attr_reader :name, :includes, :excludes
  attr_accessor :parent_distributor

  def initialize(name)
    @name = name
    @parent_distributor = nil
    @includes = {cities: {}, states: {}, countries: {}}
    @excludes = {cities: {}, states: {}, countries: {}}
  end

  def self.parse(distributors_seq, distributors_hash)
    distributors = distributors_seq.split("<").map(&:strip).map(&:upcase)
    distributor = Distributor.new(distributors.first)
    distributor.add_as_parent(distributors[1], distributors_hash)
    distributor
  end

  def add_as_parent(distributor_name, distributors_hash)
    return if distributor_name.nil?
    self.parent_distributor = distributors_hash[distributor_name]
  end

  def add_include(location, repo)
    city, state, country = *split(location.upcase)

    @includes[:cities][city] = repo.cities[city] and return if city
    @includes[:states][state] = repo.states[state] and return if state
    @includes[:countries][country] = repo.countries[country] if country
  end

  def add_exclude(location, repo)
    city, state, country = *split(location.upcase)

    @excludes[:cities][city] = repo.cities[city] and return if city
    @excludes[:states][state] = repo.states[state] and return if state
    @excludes[:countries][country] = repo.countries[country] if country
  end

  def can_distribute_in?(location, repo)
    return false if excluded_anywhere?(location, repo)
    return true if included_anywhere?(location, repo)
    false
  end

  def included_anywhere?(location, repo)
    result = included?(location, repo)
    return result if self.parent_distributor.nil?
    result || self.parent_distributor.included_anywhere?(location, repo)
  end

  def excluded_anywhere?(location, repo)
    result = excluded?(location, repo)
    return result if self.parent_distributor.nil?
    result || self.parent_distributor.excluded_anywhere?(location, repo)
  end

  private

  def split(location)
    ([nil, nil].concat location.split("-"))[-3, 3]
  end

  def included?(location, repo)
    to_bool(@includes, location, repo, true)
  end

  def excluded?(location, repo)
    to_bool(@excludes, location, repo, false)
  end

  def to_bool(source, location, repo, default)
    city, state, country = *split(location.upcase)

    city_instance = repo.cities[city]
    state_instance = repo.states[state]
    country_instance = repo.countries[country]

    city_bool = source[:cities][city_instance.name] if city_instance
    state_bool = source[:states][state_instance.name] if state_instance
    country_bool = source[:countries][country_instance.name] if country_instance

    city_bool || state_bool || country_bool
  end

  def has_no_parent_distributors?
    self.parent_distributor.nil?
  end
end