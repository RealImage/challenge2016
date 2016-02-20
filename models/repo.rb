require_relative "city"
require_relative "state"
require_relative "country"
require_relative "distributor"

class Repo
  attr_reader :cities, :states, :countries

  def initialize(cities_csv, distributors_file)
    @cities = {}
    @states = {}
    @countries = {}
    @distributors = {}

    print "Loading Repo... "
    load_cities(cities_csv)
    load_distributors(distributors_file)
    puts "Done"
  end

  def can_distribute_in?(distributor_name, location)
    @distributors[distributor_name.upcase].can_distribute_in?(location, self)
  end

  private

  def load_cities (cities_csv)
    cities = File.open(cities_csv).read
    cities.each_line do |line|
      city_code, province_code, country_code, city_name, province_name, country_name = line.split(",").map(&:strip).map(&:upcase)
      @countries[country_name] ||= Country.new(country_name)
      @states[province_name] ||= State.new(province_name, @countries[country_name])
      @cities[city_name] = City.new(city_name, @states[province_name])

      @countries[country_name].add_state(@states[province_name])
      @states[province_name].add_city(@cities[city_name])
    end
  end

  def load_distributors(distributors_file)
    distributors = File.open(distributors_file).read
    current_distributor = nil
    distributors.each_line do |distributor|
      if distributor.length == 0
        current_distributor = nil
        next
      end

      if distributor =~ /^Permissions for(.*)/i
        current_distributor = Distributor.parse($1.strip, @distributors)
        @distributors[current_distributor.name] ||= current_distributor
      elsif distributor =~ /^INCLUDE:(.*)/i
        current_distributor.add_include($1.strip, self)
      elsif distributor =~ /^EXCLUDE:(.*)/i
        current_distributor.add_exclude($1.strip, self)
      end
    end
  end
end