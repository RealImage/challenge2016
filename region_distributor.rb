class RegionDistributor

attr_accessor :name,:inc_region,:exc_region,:parent_distributor

@@distribution_object={}

	def distributor(params)
		@name = params[:name]
		@inc_region = params[:inc_region]
		@exc_region = params[:exc_region]
		@parent_distributor=params[:distributor].present? ? params[:distributor] : nil
		@exc_region << (@@distribution_object[:@parent_distributor]).exc_region
		create_distributor 
		parent = distribution_object[self.parent_distributor]
		for i in 0..(self.inc_region.size-1)
			puts "This distribution has been accepted" if check_regions(i,parent.inc_reg,parent.exc_region)
			remove_distributor(params[:name]) if !check_regions(i,parent.inc_reg,parent.exc_region)
		end

	end


	def create_distributor

		@@distribution_object[self.name] = self
		

	end
	
	def get_city_state(region)
		
		region_array = region.split("-")
		region_hash[:country] = region_array[0]
		region_hash[:state] = region_array[1].present? ? region_array[1] : ''"
		region_hash[:city] = region_array[2].present? ? region_array[2] : ''"
	end
 
	def check_regions(region1,parent_inc_reg,parent_exc_reg)

		reg_hash = get_city_state(region)
                parent_inc_reg.each do |inc|
			

			permitted = false
			permitted = true if (reg_hash[:country] == inc[:coutry] && reg_hash[:state] == inc[:state] && reg_hash[:city] == inc[:city])
               		permitted = true if (reg_hash[:country] == inc[:coutry] && reg_hash[:state] == inc[:state] && inc[:city] == "")
              		permitted = true if (reg_hash[:country] == inc[:coutry] && inc[:state] == "" && inc[:city] == "")
		end
		parent_exc_reg.each do |exc|

			permitted = false if (reg_hash[:city] == exc[:city] || reg_hash[:state] == exc[:state])
		end
		
		return permitted
	end

	def remove_distributor(name)
		
		@@distribution_object.remove!(:name)
		puts "Distributor has been removed"
	
	end
	
	def check_permissions(distributor_name,region)
		
		region_hash = get_city_state(region)
		distributor = @@distributor_object[:distributor_name]
		return true if check_regions(region_hash,distributor.inc_region,distributor.exc_region)	
		return false

	end

end
