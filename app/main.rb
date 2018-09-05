require 'csv'
require_relative 'store'
require_relative 'loader'
require_relative 'distributor'
require_relative 'errors/area_code_not_found'

root_dir = __dir__.match(/(.+)\/app$/)[1]
loader = Loader.new("#{root_dir}/app/commands.csv")
loader.load_data
