class DistributorsController < ApplicationController

  # GET /distributors
  # GET /distributors.json
  def index
    @distributors = Distributor.all
  end

  # POST /distributors
  # POST /distributors.json
  def create
    # distributor_params = {:name => "distributor2",:parent_id => 1}
    @distributor = Distributor.new(distributor_params)

    if @distributor.save
      render_success("Distributor Created successfully")
    else
      render_error("Something went wrong !!! Cound not create distributors")
    end
  end

  def allocate
    # params = {distributor_id: 1, country_id: 1, province_id: 1, city_id: 1}
    @distributor_allocation = DistributorAllocation.new(params)
    if @distributor_allocation.save
      render_success("Distributor Allocated successfully")
    else
      render_error("Something went wrong !!! Cound not allocate distributors")
    end
  end

  def check
    # params= {distributor_id: 1,city_id: 1,country_id: 1,province_id: 1} 
    # distributor_id and any one of other params is mandatory

    @distributor = Distributor.find_by_id(params[:distributor_id])
    can_distribute = @distributor.can_distribute(params)


  end

  # # PATCH/PUT /distributors/1
  # # PATCH/PUT /distributors/1.json
  # def update
  #   respond_to do |format|
  #     if @distributor.update(distributor_params)
  #       format.html { redirect_to @distributor, notice: 'Distributor was successfully updated.' }
  #       format.json { render :show, status: :ok, location: @distributor }
  #     else
  #       format.html { render :edit }
  #       format.json { render json: @distributor.errors, status: :unprocessable_entity }
  #     end
  #   end
  # end

  # # DELETE /distributors/1
  # # DELETE /distributors/1.json
  # def destroy
  #   @distributor.destroy
  #   respond_to do |format|
  #     format.html { redirect_to distributors_url, notice: 'Distributor was successfully destroyed.' }
  #     format.json { head :no_content }
  #   end
  # end

  def get_countries    
    dis_id = params["distributor_id"].to_i
    @country_status = DistributorAllocation.where("distributor_id LIKE '%#{dis_id}%'").where(country_id: params["country_id"]).present?   
  end  

  def get_provinces
    dis_id = params["distributor_id"].to_i
    if DistributorAllocation.where("distributor_id LIKE '%#{dis_id}%'").where(country_id: params["country_id"]).present?
      @provinces = Country.find(params["country_id"]).provinces
    end 
  end

  def get_cities
    dis_id = params["distributor_id"].to_i
    if DistributorAllocation.where("distributor_id LIKE '%#{dis_id}%'").where(country_id: params["country_id"], province_id: params["id"]).present?
      @cities = Province.find(params[:id]).cities
    end
  end

  def check_city
    dis_id = params["distributor_id"].to_i
    if DistributorAllocation.where("distributor_id LIKE '%#{dis_id}%'").where(country_id: params["country_id"], province_id: params["province_id"], city_id: params["id"]).present?
      @city = true
    end
  end

  private
    # Use callbacks to share common setup or constraints between actions.
    def set_distributor
      @distributor = Distributor.find(params[:id])
    end

    # Never trust parameters from the scary internet, only allow the white list through.
    def distributor_params
      params.require(:distributor).permit(:name, :parent_id)
    end

    def check_allocate_params
      params.require(:distributor).permit(:distributor_id,:city_id,:country_id,:province_id)
    end
end