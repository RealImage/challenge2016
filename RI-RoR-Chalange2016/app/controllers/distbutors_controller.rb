class DistbutorsController < ApplicationController
  before_action :set_distbutor, only: [:show, :edit, :update, :destroy]

  # GET /distbutors
  # GET /distbutors.json
  def index
    @distbutors = Distbutor.page(params[:page])
  end

  # GET /distbutors/1
  # GET /distbutors/1.json
  def show
  end
  def permision
    if check_premisions(params[:id] ,params[:permision][:countries], params[:permision][:states], params[:permision][:cities])
      render json: {status: "success", message: "Distbutor has permision for the given location"}
    else
      render json: {status: "warning", message: "Distbutor does not have permision for the given location"}
    end
  end
  # GET /distbutors/new
  def new
    @distbutor = Distbutor.new
  end
  def sub_dist
    @p_d_id = params[:id]
    @distbutor = Distbutor.new
    render :new
  end

  
  # POST /distbutors
  # POST /distbutors.json
  def create
    permision_ok = true
    distbutor_save = false
    @distbutor = Distbutor.new(distbutor_params)
    if params[:distbutor][:primary_dist_id].nil?
      params[:distbutor][:included_cities].each do |city_id|
        if city_id != ""
          city = City.find_by(id: city_id)
          state_id = city.state_id
          country_id = city.state.country_id
          permision_ok = false unless check_premisions(params[:distbutor][:primary_dist_id], country_id, state_id,city_id)
        end
      end
      if permision_ok
        params[:distbutor][:included_states].each do |state_id|
          country_id = State.find_by(id: state_id).country_id
          permision_ok = false unless check_premisions(params[:distbutor][:primary_dist_id], country_id, state_id,"")
        end
      end
      if permision_ok
        params[:distbutor][:included_countries].each do |country_id|
          permision_ok = false unless check_premisions(params[:distbutor][:primary_dist_id], country_id, "","")
        end
      end
    end
    (params[:distbutor][:included_countries] - State.where(id: params[:distbutor][:included_states]).pluck(:country_id).uniq).each{ |country_id| @distbutor.included_countries.build country_id: country_id if country_id != ""}
    (params[:distbutor][:included_states] - City.where(id: params[:distbutor][:included_cities]).pluck(:state_id).uniq).each{ |state_id| @distbutor.included_states.build state_id: state_id if state_id != ""}
    params[:distbutor][:included_cities].each{ |city_id| @distbutor.included_cities.build city_id: city_id if city_id != ""}
    (params[:distbutor][:excluded_states] - City.where(id: params[:distbutor][:excluded_cities]).pluck(:state_id).uniq).each{ |state_id| @distbutor.excluded_states.build state_id: state_id if state_id != ""}
    params[:distbutor][:excluded_cities].each{ |city_id| @distbutor.excluded_cities.build city_id: city_id if city_id != ""}
    respond_to do |format|
      if permision_ok
        distbutor_save = @distbutor.save
      end        
      if distbutor_save
        format.html { redirect_to @distbutor, notice: 'Distbutor was successfully created.' }
        format.json { render :show, status: :created, location: @distbutor }
      else
        format.html { render :new }
        format.json { render json: @distbutor.errors, status: :unprocessable_entity }
      end
    end
  end

  

  # DELETE /distbutors/1
  # DELETE /distbutors/1.json
  def destroy
    @distbutor.included_countries.destroy_all
    @distbutor.included_states.destroy_all
    @distbutor.included_cities.destroy_all
    @distbutor.excluded_states.destroy_all
    @distbutor.excluded_cities.destroy_all
    @distbutor.destroy
    respond_to do |format|
      format.html { redirect_to distbutors_url, notice: 'Distbutor was successfully destroyed.' }
      format.json { head :no_content }
    end
  end

  private
    # Use callbacks to share common setup or constraints between actions.
    def set_distbutor
      @distbutor = Distbutor.find(params[:id])
    end

    # Never trust parameters from the scary internet, only allow the white list through.
    def distbutor_params
      params.require(:distbutor).permit(:name)
    end

    def check_premisions(distbutor_id ,country_id, state_id, city_id)
      if city_id == ""
        if state_id == ""
          if IncludedCountry.find_by(distbutor_id: distbutor_id, country_id:country_id).nil?
            return false
          end
        else
          if IncludedState.find_by(distbutor_id: distbutor_id,state_id:state_id).nil?
            if IncludedCountry.find_by(distbutor_id: distbutor_id, country_id:country_id).nil?
              return false
            end
          end
          unless ExcludedState.find_by(distbutor_id: distbutor_id,state_id:state_id).nil? 
            return false
          end
        end
      else
        if IncludedCity.find_by(distbutor_id: distbutor_id,city_id:city_id).nil?
          if IncludedState.find_by(distbutor_id: distbutor_id,state_id:state_id).nil?
            if IncludedCountry.find_by(distbutor_id: distbutor_id, country_id:country_id).nil?
              return false
            end
          end
        end
        unless ExcludedCity.find_by(distbutor_id: distbutor_id,city_id:city_id).nil? 
          unless ExcludedState.find_by(distbutor_id: distbutor_id,state_id:state_id).nil? 
            return false
          end
        end
      end
      return true
    end
end
