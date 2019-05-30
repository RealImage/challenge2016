class Province < ApplicationRecord
  belongs_to :country
  has_many :cities
  has_many :distributor_allocations


  
  def self.can_check_distributor(distributor_allocations)
  	puts "Do you want to search by Province [Y/N]"
    option = gets.chomp
    case option
    when "y"
	  	puts "Enter the Province name : "
	  	province = Province.find_by(name: gets.chomp)
	  	if province.present?
	  		if distributor_allocations.present?
	  	  	province_allocated = distributor_allocations.select{|x| x.province_id == province.id && x.status == 'included'}
	  	  end
	  	  if province_allocated
	  	  	puts "province Allocated for this distributor"
	  	  else
	  	  	puts "province Not  Allocated for this distributor"
	      end
	    else
	  	  puts "province not found"
	  	end
    end
  end

  def self.can_allocate_distributor(distributor,country)
  	puts "Enter the Province name:"
    province = Province.find_by("lower(name) = ? and country_id = ?",gets.chomp,country.id)
    if province
    	return province.check_and_allocate(distributor)
    else
    	"Province Does not exist for this country"
    end
  end

  # Only if a country is allowed in parent distributor its allowed for child distributor
    
  def check_and_allocate(distributor)
  	parent_distributors = distributor.ancestors.compact
    if parent_distributors.present?
      provision_included = parent_distributors.distributor_allocations.where("status = 'included' and province_id = #{self.id} and country_id = #{self.country_id}").present?
      not_excluded = parent_distributors.distributor_allocations.where("status = 'excluded' and province_id = #{self.id} and country_id = #{self.country_id}").none?
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
