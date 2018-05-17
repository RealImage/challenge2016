# To manage a list of coutry and its province and cities
class Country
  attr_accessor :id, :name, :code, :province_and_cities

  $country_list = {}

  def initialize(params)
    self.name = case_insentive(params["Country Name"])
    self.code = case_insentive(params["Country Code"])
    self.province_and_cities = {}
    create_or_find_country(params)
  end

  def create_or_find_country(params)
    country = $country_list[[self.name, self.code]]
    country = $country_list[[self.name, self.code]] = self unless country 
    assign_state_and_cities(country, params)
  end

  def assign_state_and_cities(country, params)
    province_name = case_insentive(params["Province Name"])
    province_code = case_insentive(params["Province Code"])
    country.province_and_cities[[province_name, province_code]] ||= {}
    create_or_update_city_details(country.province_and_cities[[province_name, province_code]], params)  
  end

  def create_or_update_city_details(province_info, params)
    city_name = case_insentive(params["City Name"])
    city_code = case_insentive(params["City Code"])
    province_info[:city_list] ||=[]
    province_info[:city_list].push(city_name) unless province_info[:city_list].include? city_name
    province_info[:city_list].push(city_code) unless province_info[:city_list].include? city_code
  end

  def case_insentive(value)
    return "" if value.nil?
    value.downcase.strip
  end
end
