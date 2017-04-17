require 'csv'
require 'json'
cities = CSV.read("../app/assets/csv/cities.csv");
json = []

cities.each_with_index do |city, index|
	skip_add_state = 0
	skip_add_country = 0

	country = cities[index][2]
	state = cities[index][1]
	city = cities[index][0]

	country_name = cities[index][5]
	state_name = cities[index][4]
	city_name = cities[index][3]
	
	json.each do |j|
		if j['code'] == country
			j['children'].each do |child| # list of states of a country
				if child['code'] == state
					# add city
					state_id = child['id']
					cities_count = child['children'].count + 1
					add_city = Hash.new
					add_city["id"] = state_id.to_s+"."+cities_count.to_s
					add_city["code"] = city
					add_city["name"] = city_name
					child['children'] << add_city
					skip_add_state = 1
					skip_add_country = 1
					break
				end
			end
			# add state, add city
			if skip_add_state == 0
				country_id = j["id"]
				state_count = j["children"].count + 1
				state_id = country_id.to_s+"."+state_count.to_s

				add_state = Hash.new
				add_state["id"] = state_id
				add_state["code"] = state
				add_state["name"] = state_name
				add_state["children"] = []

				add_city = Hash.new
				add_city["id"] = state_id.to_s+"."+1.to_s
				add_city["code"] = city
				add_city["name"] = city_name

				temp_city = []
				temp_city << add_city
				add_state["children"] = temp_city

				j['children'] << add_state
				skip_add_country = 1
			end
		end
	end
	# add country, add state, add city
	if skip_add_country == 0
		add_country = Hash.new
		country_count = json.count + 1
		add_country["id"] = country_count.to_s
		add_country["code"] = country
		add_country["name"] = country_name
		add_country["children"] = []

		add_state = Hash.new
		add_state["id"] = country_count.to_s+"."+1.to_s
		add_state["code"] = state
		add_state["name"] = state_name
		add_state["children"] = []

		add_city = Hash.new
		add_city["id"] = country_count.to_s+"."+1.to_s+"."+1.to_s
		add_city["code"] = city
		add_city["name"] = city_name
		
		temp_city = []
		temp_city << add_city
		add_state["children"] = temp_city

		temp_state = []
		temp_state << add_state
		add_country["children"] = temp_state

		json << add_country
		puts country_name
	end
end

File.open("places.json","w") do |f|
  f.write(json.to_json)
end
puts "Completed"