require 'csv'
require_relative 'store'
require_relative 'loader'
require_relative 'distributor'

root_dir = __dir__.match(/(.+)\/app$/)[1]
loader = Loader.new("#{root_dir}/app/commands.csv")
loader.load_data
