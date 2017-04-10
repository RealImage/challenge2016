class DistributorsController < ApplicationController
  before_action :set_distributors
  before_action :set_distributor, only: [:edit, :update]
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
    included_locations = create_location_template(params[:permitted_city_0], params[:permitted_province_0], params[:permitted_country_0])
    excluded_locations = create_location_template(params[:restricted_city_0], params[:restricted_province_0], params[:restricted_country_0])

    @distributor = {
      name: params[:name],
      parent_id: params["parent_distributor"],
      include: [included_locations],
      exclude: [excluded_locations]
    }

    distributor_id = @distributors.keys.map {|a| a.to_i }.max.next.to_s

    @distributors[distributor_id] = @distributor
    File.open(Rails.root.join('db', 'distributors.json'), 'w') { |f| f.write(JSON.pretty_generate(@distributors)) }
    redirect_to root_url
  end

  # PATCH/PUT /distributors/1
  # PATCH/PUT /distributors/1.json
  def update
    included_array = []
    count = 0
    loop do
      temp_city_symbol = "permitted_city_" + count.to_s
      temp_province_symbol = "permitted_province_" + count.to_s
      temp_country_symbol = "permitted_country_" + count.to_s
      break if params[temp_country_symbol.to_sym].nil?
      count += 1
      next if params[temp_country_symbol.to_sym].empty?
      unless create_location_template(params[temp_city_symbol.to_sym], params[temp_province_symbol.to_sym], params[temp_country_symbol.to_sym])
        return
      end
      included_array.push(create_location_template(params[temp_city_symbol.to_sym], params[temp_province_symbol.to_sym], params[temp_country_symbol.to_sym]))
    end

    excluded_array = []
    count = 0
    loop do
      temp_city_symbol = "restricted_city_" + count.to_s
      temp_province_symbol = "restricted_province_" + count.to_s
      temp_country_symbol = "restricted_country_" + count.to_s
      break if params[temp_country_symbol.to_sym].nil?
      count += 1
      next if params[temp_country_symbol.to_sym].empty?
      unless create_location_template(params[temp_city_symbol.to_sym], params[temp_province_symbol.to_sym], params[temp_country_symbol.to_sym])
        return
      end
      excluded_array.push(create_location_template(params[temp_city_symbol.to_sym], params[temp_province_symbol.to_sym], params[temp_country_symbol.to_sym]))
    end

    @distributor = {
      name: params[:name],
      parent_id: params["parent_distributor"],
      include: included_array,
      exclude: excluded_array
    }

    @distributors[params[:id]] = @distributor
    File.open(Rails.root.join('db', 'distributors.json'), 'w') { |f| f.write(JSON.pretty_generate(@distributors)) }
    redirect_to root_url
  end

  # DELETE /distributors/1
  # DELETE /distributors/1.json
  def destroy
    @distributors.delete(params[:id])
    File.open(Rails.root.join('db', 'distributors.json'), 'w') { |f| f.write(JSON.pretty_generate(@distributors)) }
    redirect_to root_url
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

    def create_location_template(city, province, country)
      if country.empty?
        flash[:danger] = "Some issue with locations"
        redirect_to root_url
        return false
      elsif province.empty?
        return "#{country}"
      elsif city.empty?
        if @locations.select { |x| x["Province Code"] == province && x["Country Code"] == country }.count.zero?
          flash[:danger] = "Province location doesn't matched!!!"
          redirect_to root_url
          return false
        end
        return "#{province}_#{country}"
      else
        if @locations.select { |x| x["City Code"] == city && x["Province Code"] == province && x["Country Code"] == country }.count.zero?
          flash[:danger] = "City and Province location doesn't matched!!!"
          redirect_to root_url
          return false
        end
        return "#{city}_#{province}_#{country}"
      end
    end
end
