class Distributor < User
  attr_accessor :type, :sub_distributors, :include_regions, :exclude_regions, :parent

  def initialize
    self.include_regions = []
    self.exclude_regions = []
    self.sub_distributors = []
  end

end