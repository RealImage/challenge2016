require_relative './simulations/resources'
require_relative './simulations/common'
require_relative './classes/managers/distributor_manager'
require_relative './classes/managers/region_manager'

include Simulations::Common
include Simulations::Resources

@region_manager = RegionManager.new
@distributor_manager = DistributorManager.new

# Not creating the city, state, country objects as I don't see a use case for them now. We can attach it to the distributor object when they have individual states.
# create_regions
create_distributors

puts 'Distributor Status'
puts @distributor_manager.print_status
puts

puts distribute?('DISTRIBUTOR1', 'PUNCH,JK,IN')
puts distribute?('DISTRIBUTOR1', 'JK,IN')
puts distribute?('DISTRIBUTOR1', 'IN')
puts distribute?('DISTRIBUTOR1', 'TN,IN')
puts distribute?('DISTRIBUTOR2', 'TN,IN')
puts distribute?('DISTRIBUTOR1', 'KA,IN')
puts distribute?('DISTRIBUTOR2', 'KA,IN')
puts distribute?('DISTRIBUTOR3', 'KA,IN')
puts distribute?('DISTRIBUTOR3', 'TN,IN')
puts distribute?('DISTRIBUTOR4', 'TN,IN')
puts distribute?('DISTRIBUTOR4', 'KNGLM,TN,IN')
