class DistributorAreaController < ApplicationController
  def index
  end
  def search
    distributor = params[:distributor_name]
    input_areas = params[:area]
    @access = DistributorArea.is_accessible(distributor, input_areas) 
    @access = (@access == true)?  "YES" : "NO"   
  end
end
