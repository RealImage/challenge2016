Dir['../region/*'].each {|f| require f}

# As of now unused.
class RegionManager
    attr_accessor :countries, :states, :cities

    def initialize
        @countries = {}
        @states = {}
        @cities = {}
    end

    def create_region(codes)
        city_code, city_name = codes[0], codes[3]
        state_code, state_name = codes[1], codes[4]
        country_code, country_name = codes[2], codes[5]

        country = create_country({code: country_code, name: country_name})
        state = create_state({country: country, code: state_code, name: state_name})
        city = create_city({country: country, state: state, code: city_code, name: city_name})

        countries[country_code] = country
        states[state_code] = state
        cities[city_code] = city
    end

    private

    def create_country(options)
        Country.new(options)
    end

    def create_state(options)
        State.new(options)
    end

    def create_city(options)
        City.new(options)
    end
end