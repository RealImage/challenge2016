class HomesController< ApplicationController
	before_action :get_distributors_json_from_assets
	before_action :get_cities_json_from_assets
	
	def index
	end

	def get_parent
		parent_id = params[:parent_id].to_i
		@distributors.each do |distributor|
			if distributor["id"] == parent_id
				render json: distributor
			end
		end
		
	end

	def get_countries
		cities = []
		@cities.each do |city|
			cities << [city["code"], city["name"]]
		end
		render json: cities
	end

	def get_states_for_country
		country_id = params[:country_id]
		states = []
		@cities.each do |city|
			if city["code"] == country_id
				city["children"].each do |state|
					states << [ state["code"], state["name"] ]
				end
			end
		end
		render json: states
	end

	def get_cities_for_state
		state_id = params[:state_id]
		country_id = params[:country_id]
		cities = []
		@cities.each do |city|
			if city["code"] == country_id
				city["children"].each do |state|
					if state["code"] == state_id
						state["children"].each do |taken_city|
							cities << [ [taken_city["code"]], taken_city["name"]]
						end
					end
				end
			end
		end
		render json: cities
	end

	def distributors_list
	end

	def show_distributors
		id = params[:id].to_i
		@distributors.each do |distributor|
			if distributor["id"] == id
				@distributor = distributor
				break
			end
		end
	end

	def new_distributor
	end

	def new_sub_distributor
		@parent_id = params[:parent_id]
	end

	def create_distributor
		create_new_distributor
	end

	def create_sub_distributor
		create_new_distributor(params[:parent_id])
	end

	def add_new_location
		@id = params[:id]
	end

	def create_new_location
		id = params[:id].to_i
		@distributors.each do |distributor|
			if distributor["id"] == id
				is_included = false
				is_excluded = false

				is_parent = distributor["parent_id"].nil? && distributor["parent_id"].blank?

				inclusion_location_with_name =  find_ids(params[:inc_country], params[:inc_state], params[:inc_city]) unless (params[:inc_country].blank? && params[:inc_country].empty?)
				# binding.pry
				unless inclusion_location_with_name.nil? && inclusion_location_with_name.blank?
					unless inclusion_location_with_name[0].nil? && inclusion_location_with_name[0].blank?
					 if check_for_inclusion(get_locations_for_inclusion(id), inclusion_location_with_name[0], is_parent)
					 		distributor["inc_locations"] << inclusion_location_with_name[0]
					 		distributor["inc_locations_with_name"] << inclusion_location_with_name[1]
					 		save_distributors_to_json_file
					 		is_included = true
					 else
					 		flash[:notice] = "Cannot add Inclusion Location"
					 		redirect_to list_path and return
					 end
					end
				end

				exclusion_location_with_name = find_ids(params[:exc_country], params[:exc_state], params[:exc_city]) unless (params[:exc_country].blank? && params[:exc_country].empty?)

				#binding.pry
				unless exclusion_location_with_name.nil? && exclusion_location_with_name.blank?
					unless exclusion_location_with_name[0].nil? && exclusion_location_with_name[0].blank?
						if check_for_exclusion(get_locations_for_exclusion(id), exclusion_location_with_name[0])
							distributor["exc_locations"] << exclusion_location_with_name[0]
							distributor["exc_locations_with_name"] << exclusion_location_with_name[1]
							save_distributors_to_json_file
							is_excluded = true
						else
							flash[:notice] = "Cannot add Exclusion Location"
							redirect_to list_path and return
						end
					end
				end

				
				# get_locations_for_inclusion(id)
				# get_locations_for_exclusion(id)
				# binding.pry

				#save_distributors_to_json_file
				if is_included || is_excluded
					flash[:notice] = "Created Successfully..!"
					redirect_to list_path
				else
					redirect_to list_path
				end
			end
		end
	end

	def check_authorisation
		@id = params[:id]
	end

	def check_authorisation_for_given_location
		id = params[:id].to_i
		@distributors.each do |distributor|
			if distributor["id"] == id
				is_parent = distributor["parent_id"].nil? && distributor["parent_id"].blank?

				location =  find_ids(params[:inc_country], params[:inc_state], params[:inc_city]) unless (params[:inc_country].blank? && params[:inc_country].empty?)
				unless location.nil? && location.blank?
					if check_for_authorisation(get_locations_for_authorisation(id), location, is_parent)
						flash[:notice] = "Distributor have access to this location"
							redirect_to list_path and return
					else
						flash[:notice] = "Distributor doesn't have access to this location"
							redirect_to list_path and return
					end
				end
			end
		end
	end

	private

	def save_distributors_to_json_file
		File.open("app/assets/json/distributors.json","w") do |f|
		  f.write(@distributors.to_json)
		end
	end

	def get_distributors_json_from_assets
		file = File.read("app/assets/json/distributors.json")
		@distributors = []
		unless file.empty?
			@distributors = JSON.parse(file)
		end
	end

	def get_cities_json_from_assets
		file = File.read("app/assets/json/places.json")
		@cities = JSON.parse(file)
	end

	def create_new_distributor(parent_id = nil)
		distributors_count = 1
		unless @distributors.nil?
			distributors_count = @distributors.count + 1
		end

		new_distributor = Hash.new
		new_distributor["id"] = distributors_count
		new_distributor["name"] = params[:distributor_name]
		new_distributor["inc_locations"] = []
		new_distributor["inc_locations_with_name"] = []
		new_distributor["exc_locations"] = []
		new_distributor["exc_locations_with_name"] = []
		new_distributor["parent_id"] = parent_id
		@distributors << new_distributor
		save_distributors_to_json_file
		flash[:notice] = "Created Successfully..!"
		redirect_to list_path
	end

	def find_ids(given_country, given_state, given_city)
		@cities.each do |city|
			if city['code'] == given_country
				unless given_state.empty?
					city['children'].each do |state|
						if state['code'] == given_state
							unless given_city.empty?
								state['children'].each do |taken_city|
									return taken_city['id'], taken_city['name'] if taken_city['code'] == given_city
								end
							else
								return state['id'], state['name']
							end
						end
					end
				else
					return city['id'], city['name']
				end
			end
		end
		return [], []
	end

	def insert_to_array(array, array_to_be_inserted)
		array_to_be_inserted.each do |item|
			array << item
		end
	end

	def get_parent_locations(distributor_id, check_inclusion, check_exclusion)
		@distributors.each do |distributor|
			if distributor["id"] == distributor_id
				unless distributor["parent_id"].nil? && distributor["parent_id"].blank?
					insert_to_array(check_inclusion, distributor["inc_locations"])
					insert_to_array(check_exclusion, distributor["exc_locations"])
					get_parent_locations(distributor["parent_id"].to_i, check_inclusion, check_exclusion)
				else
					insert_to_array(check_inclusion, distributor["inc_locations"])
					insert_to_array(check_exclusion, distributor["exc_locations"])
					break
				end
			end
		end
		return check_inclusion, check_exclusion
	end

	def get_locations_for_inclusion(distributor_id, check_inclusion = [], check_exclusion = [])
		@distributors.each do |distributor|
			if distributor["id"] == distributor_id
				unless distributor["parent_id"].nil? && distributor["parent_id"].blank?
					insert_to_array(check_exclusion, distributor["exc_locations"])
					return get_parent_locations(distributor["parent_id"].to_i, check_inclusion, check_exclusion)
				else
					insert_to_array(check_exclusion, distributor["exc_locations"])
					break
				end
			end
		end
		return check_inclusion, check_exclusion
	end

	def get_locations_for_exclusion(distributor_id, check_inclusion = [], check_exclusion = [])
		@distributors.each do |distributor|
			if distributor["id"] == distributor_id
				insert_to_array(check_inclusion, distributor["inc_locations"])
				insert_to_array(check_exclusion, distributor["exc_locations"])
				break
			end
		end
		return check_inclusion, check_exclusion
	end

	def get_locations_for_authorisation(distributor_id, check_inclusion = [], check_exclusion = [])
		@distributors.each do |distributor|
			if distributor["id"] == distributor_id
				unless distributor["parent_id"].nil? && distributor["parent_id"].blank?
					insert_to_array(check_inclusion, distributor["inc_locations"])
					insert_to_array(check_exclusion, distributor["exc_locations"])
					return get_parent_locations(distributor["parent_id"].to_i, check_inclusion, check_exclusion)
				else
					insert_to_array(check_inclusion, distributor["inc_locations"])
					insert_to_array(check_exclusion, distributor["exc_locations"])
					break
				end
			end
		end
		return check_inclusion, check_exclusion
	end

	# check the given location can be added or not.

	def check_country_presence(locations, location_to_be_added)
		locations.each do |location|
			if location.length == 1
				return true if location.to_i == location_to_be_added.to_i
			end
		end
		false
	end

	def check_state_presence(locations, location_to_be_added)
		splited_locations = location_to_be_added.split('.')
		return true if check_country_presence(locations, splited_locations[0])
		locations.each do |location|
			if location.length == 3
				return true if location.to_s == location_to_be_added.to_s
			end
		end
		false
	end

	def check_city_presence(locations, location_to_be_added)
		splited_locations = location_to_be_added.split('.')
		return true if check_country_presence(locations, splited_locations[0])
		return true if check_state_presence(locations, (splited_locations[0]+"."+splited_locations[1]).to_s)
		locations.each do |location|
			if location.length == 5
				return true if location.to_s == location_to_be_added.to_s
			end
		end
		false
	end

	def check_given_location(locations, location_to_be_added)
		if location_to_be_added.length == 1
			return check_country_presence(locations, location_to_be_added)
		elsif location_to_be_added.length == 3
			return check_state_presence(locations, location_to_be_added)
		elsif location_to_be_added.length == 5
			return check_city_presence(locations, location_to_be_added)
		else
			return false
		end
	end

	def check_for_inclusion(locations, location_to_be_added, is_parent)
		inc_locations = locations[0]
		exc_locations = locations[1]

		if is_parent
			return !check_given_location(exc_locations, location_to_be_added)
		else
			return (check_given_location(inc_locations, location_to_be_added) && !check_given_location(exc_locations, location_to_be_added))
		end
	end

	def check_for_exclusion(locations, location_to_be_added)
		inc_locations = locations[0]
		exc_locations = locations[1]
		
		return (check_given_location(inc_locations, location_to_be_added) && !check_given_location(exc_locations, location_to_be_added))
	end

	def check_for_authorisation(locations, location_to_be_added, is_parent)
		inc_locations = locations[0]
		exc_locations = locations[1]

		return (check_given_location(inc_locations, location_to_be_added) && !check_given_location(exc_locations, location_to_be_added))
	end
end