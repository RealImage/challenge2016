class DistributorsController < ApplicationController
  before_action :set_distributors, except: [:destroy]
  before_action :set_distributor, only: [:edit, :update, :destroy]
  before_action :set_locations, except: [:index, :destroy]

  require 'json'
  require 'csv'
  # GET /distributors
  # GET /distributors.json
  def index
  end

  # GET /distributors/new
  def new
    @distributor = Hash.new
  end

  def check
    @distributor = Hash.new

    unless params["distributor_id"].nil? || params["distributor_id"].empty?
      unless params["country"].nil? || params["country"].empty?
        if check_permission(params["distributor_id"], params["country"], params["province"], params["city"])
          @message = "The distributor has permission for this location"
        else
          @message = "The distributor does not have permission for this location"
        end
      else
        @message = "Please enter a location to check permission!"
      end
    else
      @message = "Please select a distributor to check for permissions!"
    end
  end

  # GET /distributors/1/edit
  def edit
  end

  # POST /distributors
  # POST /distributors.json
  def create
    @distributor = Distributor.new(distributor_params)

    respond_to do |format|
      if @distributor.save
        format.html { redirect_to @distributor, notice: 'Distributor was successfully created.' }
        format.json { render :show, status: :created, location: @distributor }
      else
        format.html { render :new }
        format.json { render json: @distributor.errors, status: :unprocessable_entity }
      end
    end
  end

  # PATCH/PUT /distributors/1
  # PATCH/PUT /distributors/1.json
  def update
    respond_to do |format|
      if @distributor.update(distributor_params)
        format.html { redirect_to @distributor, notice: 'Distributor was successfully updated.' }
        format.json { render :show, status: :ok, location: @distributor }
      else
        format.html { render :edit }
        format.json { render json: @distributor.errors, status: :unprocessable_entity }
      end
    end
  end

  # DELETE /distributors/1
  # DELETE /distributors/1.json
  def destroy
    @distributor.destroy
    respond_to do |format|
      format.html { redirect_to distributors_url, notice: 'Distributor was successfully destroyed.' }
      format.json { head :no_content }
    end
  end

  private
    # Use callbacks to share common setup or constraints between actions.
    def set_distributor
      @distributor = @distributors[params[:id]]
    end

    def set_distributors
      @distributors = JSON.parse(File.read(Rails.root.join('db', 'distributors.json')))
    end

    def set_locations
      @locations = CSV.read(Rails.root.join('db', 'cities.csv'), :headers=>true)
    end

    def check_permission(distributor_id, country, province, city)
      # Check for the current distributor exclusion
      excluded = check_exclusion(distributor_id, country, province, city)
      return false if excluded

      traverse_path = []
      traverse_path.push(distributor_id)
      # else check until parent reached
      parent_id = @distributors[distributor_id]['parent_id']
      until parent_id.empty?
        return false if check_exclusion(parent_id, country, province, city)
        traverse_path.push(parent_id)
        parent_id = @distributors[parent_id]['parent_id']
      end

      traverse_path.reverse.each do |dist_id|
        included = check_inclusion(dist_id, country, province, city)
        return false unless included
      end
      return true
      # included = check_exclusion(distributor_id, country, province, city)
      # return true if included
    end

    def check_exclusion(distributor_id, country, province, city)
      excluded_locations = @distributors[distributor_id]['exclude']

      excluded_locations.each do |location|
        if location =~ /^#{country}$/ || location =~ /^#{province}_#{country}$/ || location =~ /^#{city}_#{province}_#{country}$/
          return true
        end
      end
      return false
    end

    def check_inclusion(distributor_id, country, province, city)
      included_locations = @distributors[distributor_id]['include']

      included_locations.each do |location|
        if location =~ /^#{country}$/ || location =~ /^#{province}_#{country}$/ || location =~ /^#{city}_#{province}_#{country}$/
          return true
        end
      end
      return false
    end

    # Never trust parameters from the scary internet, only allow the white list through.
    def distributor_params
      params.fetch(:distributor, {})
    end
end
