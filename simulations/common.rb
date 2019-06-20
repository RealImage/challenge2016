require_relative '../constants'

module Simulations
  module Common
    include Constants

    def split_region(region)
      region.split(',')
    end

    def find_type(region)
      splits = split_region(region).size

      splits == 1 ? COUNTRY : ((splits == 2) ? STATE : CITY)
    end

    def get_regions(region)
      part = split_region(region)
      type = find_type(region)

      regions = if type == COUNTRY
                  { country: part[0] }
                elsif type == STATE
                  { country: part[1], state: part[0] }
                else
                  { country: part[2], state: part[1], city: part[0] }
      end

      [type, regions]
    end

    def distribute?(distributor, region)
      type, regions = get_regions(region)
      puts "Distributor #{distributor}, Region #{region}"

      @distributor_manager.permitted_to_distribute?(distributor, type, regions) ? 'YES' : 'NO'
    end
  end
end
