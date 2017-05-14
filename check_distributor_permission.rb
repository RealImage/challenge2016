# General comments - Always use Country-Province-City heirarchy (for what ever input) to check distributors permission
# eg: India-Tamil Nadu-Chennai (allowed)
# Use only Name of the regions, this porgram runs based on the names rather than code.
# currently the distributor's rights are hardcoded inside a hash in this file - Change the permissions if they need to be altered.
# usage with example:
#
# ATH017460:qube tlogesh$ ruby check_distributor_permission.rb
#   Enter distributor name:
#   distributor1
#   Enter the region to check:
#   India-West Bengal-Titagarh
#   The distributor doesnt have rights to distribute in India-West Bengal-Titagarh
#
# ATH017460:qube tlogesh$ ruby check_distributor_permission.rb
#   Enter distributor name:
#   distributor2
#   Enter the region to check:
#   India-Tamil Nadu-Chennai
#   The distributor has rights to distribute in India-Tamil Nadu-Chennai
#

require 'csv'

# this method is used to form a tree structure based on the data fetched from the CSV.
# tree Structure:
# world_regions {
#	 :India => {
#		 :Tamil Nadu => {
#          :Chennai => Chennai
#        }
#        :Andhra Pradesh => {
#			:Hyderabad => Hyderabad
#        }
#	 }
#    :Country Name => {
#       :Province Name => {
#		   :City Name => City Name
#        }
# 	 }
# }
def form_tree_from_csv
  world_regions = {}

  CSV.foreach('cities.csv', headers: true) do |row|
    # intialize key (if doesn't exist) for each country to hold a hash
    world_regions[row.to_hash['Country Name']] = {} unless
                                                  world_regions.key?(row.to_hash['Country Name'])

    # intialize key (if doesn't exist) for each province inside country hash to hold a hash
    world_regions[row.to_hash['Country Name']][row.to_hash['Province Name']] = {} unless
                        world_regions[row.to_hash['Country Name']].key?(row.to_hash['Province Name'])

    # push the city name as the end nodes
    world_regions[row.to_hash['Country Name']][row.to_hash['Province Name']][row.to_hash['City Name']] = row.to_hash['City Name']
  end

  # return the tree structure
  world_regions
end

# Hash containing distributors permission
distributors = {
  distributor1: {
    include: 'India,United States',
    exclude: 'India-Tamil Nadu'
  },
  distributor2: {
    include: 'India-Tamil Nadu',
    exclude: 'India-Tripura',
    inherits: 'distributor1' # Hard code all the inherits
  }
}

# Method to intialize a distributor
#
# then sets the actual permissions by calling set_distributor_rights
# distributor => Hash of distributors
# world_regions => data retrived from csv
# distributor_name => Name of the distributor for whom the permission has to be checked

def configure_distributor(
  distributors:,
  world_regions:,
  distributor_name:
)
  distributor_rights = {}
  if distributors.key?(distributor_name.to_sym)

    # configure the permission of inherits first
    if distributors[distributor_name.to_sym].key?('inherits'.to_sym)
      distributors[distributor_name.to_sym]['inherits'.to_sym].split(',').each do |a|
        include_string = distributors[a.to_sym]['include'.to_sym]
        exclude_string = distributors[a.to_sym]['exclude'.to_sym]

        distributor_rights = distributor_rights.merge(set_distributor_rights(
                                                        includeString: include_string,
                                                        world_regions: world_regions
        ))
        distributor_rights = exclude_distributor_rights(
          excludeString: exclude_string,
          distributor_rights: distributor_rights
        )
      end
    end

    # configure the permission specific to the disributor, this has to be seperate to avoid overwritting
    include_string = distributors[distributor_name.to_sym]['include'.to_sym]
    exclude_string = distributors[distributor_name.to_sym]['exclude'.to_sym]

    distributor_rights = distributor_rights.merge(set_distributor_rights(
                                                    includeString: include_string,
                                                    world_regions: world_regions
    ))

    distributor_rights = exclude_distributor_rights(
      excludeString: exclude_string,
      distributor_rights: distributor_rights
    )
  end
  distributor_rights
end

# The actual method that sets the permission
# includeString => string representing the regions that has to be included
def set_distributor_rights(
  world_regions:,
  includeString:
)
  distributor_rights = {}
  includeString.split(',').each do |region|
    # in this case set access to the entire country (province and city)
    # Assuming the heirarchy City cannot exists without a province.
    if region.split('-')[0] &&
       !region.split('-')[1]

      distributor_rights = distributor_rights
                           .merge(set_country_rights(
                                    world_regions: world_regions,
                                    set_rights_to: region.split('-')[0]
                           ))
    end

    # Set only specific provincial access inside a country
    if region.split('-')[1] &&
       !region.split('-')[2]

      # Intialize the country key to hold an hash, if the key doesn't exist already
      distributor_rights[region.split('-')[0]] = {} unless distributor_rights.key?(region.split('-')[0])

      # set provincial rights
      distributor_rights[region.split('-')[0]] = distributor_rights[region.split('-')[0]]
                                                 .merge(set_provincial_rights(
                                                          world_regions: world_regions,
                                                          set_rights_to: region.split('-')[1],
                                                          country: region.split('-')[0]
                                                 ))
    end

    # set only specific city access
    next unless region.split('-')[2]

    # intialize the country and province key to hold an hash, if they don't exist
    distributor_rights[region.split('-')[0]] = {} unless distributor_rights.key?(region.split('-')[0])
    distributor_rights[region.split('-')[0]][region.split('-')[1]] = {} unless
                                                                        distributor_rights[region.split('-')[0]].key?(region.split('-')[1])

    # set city rigths
    distributor_rights[region.split('-')[0]][region.split('-')[1]] = distributor_rights[region.split('-')[0]][region.split('-')[1]]
                                                                     .merge(set_city_rights(
                                                                              world_regions: world_regions,
                                                                              set_rights_to: region.split('-')[2],
                                                                              country: region.split('-')[0],
                                                                              province: region.split('-')[1]
                                                                     ))
  end

  # return the distributor rights in a tree structure back again
  #
  #   allowed_country=>{
  #     allowed_province=>{
  #       allowed_city =>
  #     }
  #   }
  #
  distributor_rights
end

# This removes the excluded items from the allowed permission tree returned by the previous method
# excludeString => String representing the regions that has to be exclude from the included regions

def exclude_distributor_rights(
  excludeString:,
  distributor_rights:
)
  excludeString.split(',').each do |region|
    # remove city access
    if region.split('-')[2]
      distribution_rights = remove_city_rights(
        distributor_rights: distributor_rights,
        country: region.split('-')[0],
        province: region.split('-')[1],
        rights_to_remove: region.split('-')[2]
      )
    end

    # remove province access
    if region.split('-')[1] &&
       !region.split('-')[2]
      distributor_rights = remove_province_rights(
        distributor_rights: distributor_rights,
        country: region.split('-')[0],
        rights_to_remove: region.split('-')[1]
      )
    end

    # removing country access is skipped - very rare chances to occur.
  end

  # return the final tree strucutre containing distributors distribution area
  distributor_rights
end

# remove the city specific rights
def remove_city_rights(
  distributor_rights:,
  country:,
  province:,
  rights_to_remove:
)
  distributor_rights[country][province] = distributor_rights[country][province].reject { |k, v| k == rights_to_remove && v == rights_to_remove }
  distributor_rights
end

# remove the province rights (also remove the following city rights)
def remove_province_rights(
  distributor_rights:,
  country:,
  rights_to_remove:
)
  distributor_rights[country] = distributor_rights[country].reject { |k, _v| k == rights_to_remove }
  distributor_rights
end

# Sets country specific access by allowing access to all the
# province and city beneath the country in tree struture.
def set_country_rights(
  world_regions:,
  set_rights_to:
)
  tmp_hash = world_regions.select { |k, _v| k == set_rights_to.strip }
  tmp_hash
end

# Sets provinice specific access by allowing access to all the
# city beneath the province in tree struture.
def set_provincial_rights(
  world_regions:,
  set_rights_to:,
  country:
)
  tmp_hash = (world_regions[country].select { |k, _v| k == set_rights_to.strip })
  tmp_hash
end

# sets pmerission for specific city
def set_city_rights(
  world_regions:,
  set_rights_to:,
  country:,
  province:
)
  tmp_hash = (world_regions[country][province].select { |k, _v| k == set_rights_to.strip })
  tmp_hash
end

def check_distributor_rights(
  distribution_rights:,
  region_name:
)
  # omitting country leve search here, since search are more specific either to province
  # or city
  if region_name.split('-')[1] && !region_name.split('-')[2]
    if distribution_rights[region_name.split('-')[0]].key?(region_name.split('-')[1])
      puts 'The distributor has rights to distribute in '\
        + region_name.split('-')[0] + '-' + region_name.split('-')[1]
    else
      puts "The distributor doesn't have rights to distribute in "\
        + region_name.split('-')[0] + '-' + region_name.split('-')[1]
    end
  end

  # city specific search
  if region_name.split('-')[1] && region_name.split('-')[2]
    if distribution_rights[region_name.split('-')[0]].key?(region_name.split('-')[1])
      if distribution_rights[region_name.split('-')[0]][region_name.split('-')[1]].key?(region_name.split('-')[2])
        puts 'The distributor has rights to distribute in '\
          + region_name.split('-')[0] + '-' + region_name.split('-')[1] + '-' + region_name.split('-')[2]
      else
        puts 'The distributor doesnt have rights to distribute in '\
          + region_name.split('-')[0] + '-' + region_name.split('-')[1] + '-' + region_name.split('-')[2]
      end
    else
      puts 'The distributor doesnt have rights to distribute in '\
        + region_name.split('-')[0] + '-' + region_name.split('-')[1] + '-' + region_name.split('-')[2]
    end
  end
end

world_regions = form_tree_from_csv

puts 'Enter distributor name:'
distributor_name = gets.chomp

puts 'Enter the region to check:'
region_name = gets.chomp

distribution_rights = configure_distributor(
  distributor_name: distributor_name,
  world_regions: world_regions,
  distributors: distributors
)

check_distributor_rights(
  distribution_rights: distribution_rights,
  region_name: region_name
)
