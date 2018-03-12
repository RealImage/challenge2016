module DistributorHelper
	def format_regions(region_names=[])
		regions_hash={}
		region_order={0=>'Country Name',1=>'Province Name',2=>'City Name'}
		region_order.values.each{|x|regions_hash[x]=[]}
		(region_names||[]).each do|regions|
			regions_list = regions.split('-').map{|x|x.squish}
			regions_list.reverse.each_with_index do |code,i|
				regions_hash[region_order[i]]<<code.titleize if !regions_hash[region_order[i]].include?(code)
			end
		end
		regions_hash
	end

	def check_include_parent_permissions(child_regions,parent_include_regions={},parent_exclude_regions={})
		parent_include_regions||={}
		parent_exclude_regions||={}
		if !(parent_include_regions.empty? || child_regions["Country Name"].empty?)
			if parent_include_regions["Country Name"].include?(child_regions["Country Name"].first) && !parent_exclude_regions["Country Name"].include?(child_regions["Country Name"].first)
				return true if child_regions["Province Name"].empty? || parent_include_regions["Province Name"].empty?
				if parent_include_regions["Province Name"].include?(child_regions["Province Name"].first)
					return true if child_regions["City Name"].empty? || parent_include_regions["City Name"].empty?||parent_include_regions["City Name"].include?(child_regions["City Name"].first)
				end
			elsif parent_include_regions["Country Name"].include?(child_regions["Country Name"].first) && parent_exclude_regions["Country Name"].include?(child_regions["Country Name"].first)
				return true if child_regions["Province Name"].empty?||(parent_include_regions["Province Name"].empty?&& !parent_exclude_regions["Province Name"].include?(child_regions["Province Name"].first))
				if parent_include_regions["Province Name"].include?(child_regions["Province Name"].first) && !parent_exclude_regions["Province Name"].include?(child_regions["Province Name"].first)
					return true if child_regions["City Name"].empty?||parent_include_regions["City Name"].include?(child_regions["City Name"].first)||(parent_include_regions["City Name"].empty?)
				elsif parent_include_regions["Province Name"].include?(child_regions["Province Name"].first) && parent_exclude_regions["Province Name"].include?(child_regions["Province Name"].first)				
					return true if child_regions["City Name"].empty?||parent_include_regions["City Name"].include?(child_regions["City Name"].first)||((parent_include_regions["City Name"].empty?)&&!parent_exclude_regions["City Name"].include?(child_regions["City Name"].first))
				end
			end
		end
		return false
	end

	def check_exclude_parent_permissions(include_regions,exclude_regions)
		if !(exclude_regions["Country Name"].empty?||include_regions["Country Name"].empty?)||!include_regions["Country Name"].include?(exclude_regions["Country Name"].first)
			if include_regions["Country Name"].include?(exclude_regions["Country Name"].first)
				return true if include_regions["Province Name"].empty? && !exclude_regions["Province Name"].empty?
				if include_regions["Province Name"].include?(exclude_regions["Province Name"].first)
					return true if (include_regions["City Name"].empty? && !exclude_regions["City Name"].empty?)||include_regions["City Name"].include?(exclude_regions["City Name"].first)
				end
			end
		end
		return false
	end

	def titleize
  		self.split(" ").map {|word| word.capitalize}.join(" ")
	end

	def squish
	  gsub!(/\A[[:space:]]+/, '')
	  gsub!(/[[:space:]]+\z/, '')
	  gsub!(/[[:space:]]+/, ' ')
	  self
	end
end