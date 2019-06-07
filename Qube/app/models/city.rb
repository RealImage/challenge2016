class City < ApplicationRecord
  belongs_to :province
  has_many :distributor_allocations



  def self.check_distributor(distributor_allocations)
  	puts "Do you want to search by city [Y/N]"
    option = gets.chomp
    case option
    when "y"
	  	puts "Enter the city name : "
	  	city = City.find_by(name: gets.chomp)
	  	if city.present?
	  		if distributor_allocations.present?
	  	  	city_allocated = distributor_allocations.select{|x| x.city_id == city.id && x.status == 'included'}
	  	  end
	  	  if city_allocated
	  	  	puts "city Allocated for this distributor"
	  	  else
	  	  	puts "city Not  Allocated for this distributor"
	      end
	    else
	  	  puts "city not found"
	  	end
    end
  end


  def self.can_allocate_distributor(distributor,province)
  	puts "Enter the city name:"
    city = City.find_by("lower(name) = ? and province_id = ?",gets.chomp,province.id)
    if city
    	return city.check_and_allocate(distributor)
    else
    	"City Does not exist for this province"
    end
  end

  # Only if a country is allowed in parent distributor its allowed for child distributor
    
  def check_and_allocate(distributor)
  	parent_distributors = distributor.ancestors.compact
    if parent_distributors.present?
      provision_included = parent_distributors.distributor_allocations.where("status = 'included' and province_id = #{self.province_id} and country_id = #{self.country_id} and city_id = #{self.id}").present?
      not_excluded = parent_distributors.distributor_allocations.where("status = 'excluded' and province_id = #{self.province_id} and country_id = #{self.country_id} and city_id = #{self.id}").none?
      if provision_included || not_excluded
        return self
      else
      	puts "Cannot allocate to this distributor, parent distributor have no permission for distribution"
      	return nil
      end
    else
      return self #No parent exist, so it can be allocated blindly
    end
  end


end
