class Country < ApplicationRecord
  has_many :provinces
  has_many :distributor_allocations


  def self.check_distributor(distributor_allocations)
  	puts "Do you want to search by country [Y/N]"
    option = gets.chomp
    case option
    when "y"
	  	puts "Enter the country name : "
	  	country = Country.find_by(name: gets.chomp)
	  	if country.present?
	  	  country_allocated = distributor_allocations.select{|x| x.country_id == country.id && x.status == 'included'}
	  	  if country_allocated
	  	  	puts "Country Allocated for this distributor"
	  	  else
	  	  	puts "Country Not  Allocated for this distributor"
	      end
	    else
	  	  puts "Country not found"
	  	end
    end
  end


  def self.can_allocate_distributor(distributor)
  	puts "Enter the country name:"
    c = Country.find_by("lower(name) = ?",gets.chomp)
    parent_distributors = distributor.ancestors.compact
    # Only if a country is allowed in parent distributor its allowed for child distributor
    if parent_distributors.present?
      country_included = parent_distributors.distributor_allocations.where("status = 'included' and country_id = #{c.id}").present?
      country_excluded = parent_distributors.distributor_allocations.where("status = 'excluded' and country_id = #{c.id}").none?
      if country_included || country_excluded
        return c
      else
      	puts "Cannot allocate to this distributor, parent distributor have no permission for distribution"
      	return nil
      end
    else
      return c
    end
  end


end
