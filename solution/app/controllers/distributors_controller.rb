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

    unless params["distributor"].nil? || params["distributor"].empty?
      unless params["country"].nil? || params["country"].empty?
        @message = check_permission
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

    def check_permission
      return "The distributor has/doesn't have permissions to ditribute in this location!"
    end

    # Never trust parameters from the scary internet, only allow the white list through.
    def distributor_params
      params.fetch(:distributor, {})
    end
end
