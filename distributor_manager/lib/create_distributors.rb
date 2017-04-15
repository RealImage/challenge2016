require 'json'
require 'csv'

distributors = []
file = File.read("../app/assets/json/places.json")
cities = JSON.parse(file)

def find_ids(inc_country, inc_state, inc_city, cities)
	cities.each do |city|
		if city['code'] == inc_country
			unless inc_state.empty?
				city['children'].each do |state|
					if state['code'] == inc_state
						unless inc_city.empty?
							state['children'].each do |city|
								return city['id'] if city['code'] == inc_city
							end
						else
							return state['id']
						end
					end
				end
			else
				return city['id']
			end
		end
	end
end

name =  "sre" # name
inc_country =  "IN" # country
inc_state =  "JK" # state
inc_city =  "PUNCH" # city

exc_country =  "IN" # country
exc_state =  "JK" # state
exc_city =  "PUNCH" # city


new_distributor = Hash.new
distributors_count = distributors.count + 1
new_distributor["id"] = distributors_count
new_distributor["name"] = name
new_distributor["inc_locations"] = find_ids(inc_country, inc_state, inc_city, cities)
#new_distributor["exc_locations"] = 

distributors << new_distributor

File.open("distributors.json","w") do |f|
  f.write(distributors.to_json)
end