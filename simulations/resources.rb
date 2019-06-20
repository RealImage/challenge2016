require 'csv'
require_relative './config'

module Simulations
  module Resources
    include Config

    def create_regions
      cities = CSV.read('cities.csv')

      cities.each do |city|
        @region_manager.create_region(city)
      end
    end

    def create_distributors
      distributors.each do |distributor|
        create_distributor(distributor, nil)
      end
    end

    def create_distributor(distributor, inherits)
      name = distributor[:name]
      @distributor_manager.create_distributor(name: name, inherits: inherits)

      (distributor[:included] || []).each do |region|
        type, regions = get_regions(region)

        if @distributor_manager.valid_parent?(name, type, regions)
          @distributor_manager.add_included(name, type, regions[type])
        end
      end

      (distributor[:excluded] || []).each do |region|
        type, regions = get_regions(region)

        if @distributor_manager.valid_parent?(name, type, regions)
          @distributor_manager.add_excluded(name, type, regions[type])
        end
      end

      (distributor[:sub_distributors] || []).each do |sub_distributor|
        inherits = distributor[:name]
        create_distributor(sub_distributor, inherits)
      end
    end
  end
end
