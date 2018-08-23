class Region
  def self.set_region(new_distributor)
    print "\nEnter number of regions you want to INCLUDE for #{new_distributor.username.capitalize}: "
    regions_count = gets.chomp.to_i
    puts "\nAssign distrubution regions for #{new_distributor.username.capitalize} like below format:"
    puts "eg1: India\neg2: Tamil Nadu-India\neg3: yawal-karnataka-india\neg4: United States"
    assign_regions(regions_count,new_distributor)
    print "\nEnter number of regions you want to EXCLUDE for #{new_distributor.username.capitalize}: "
    exclude_regions_count = gets.chomp.to_i
    assign_regions(exclude_regions_count,new_distributor,false)
  end

  def self.assign_regions(regions_count, new_distributor, include=true)
    (1..regions_count).each do |r|
      print "\nEnter region #{r}: "
      region = gets.chomp
      if !include && new_distributor.include_regions.include?(region)
        puts "You can't exclude included region. Please try with other region."
        assign_regions(regions_count, new_distributor, include)
      end
      result = new_distributor.type == "Distributor" ? Region.search_region(region) :
      (!Region.search_region(region, new_distributor.parent.exclude_regions, false) && ( Region.search_region(region, new_distributor.parent.include_regions, include) && Region.search_region(region) ) )
      if result
        include ? new_distributor.include_regions.push(region) : new_distributor.exclude_regions.push(region)
        puts "\nAssigned #{region} region successfully"
      else
        puts new_distributor.type == "Distributor" ? "\nEntered region is not available or check spell mistake!" : "\nYou don't have distribution rights in this region. Please try inside your region!"
        assign_regions(regions_count, new_distributor, include)
      end
    end
  end

  def self.search_region(keyword, regions = nil, include=nil)
    permit_regions = keyword.split('-')
    regions_count = permit_regions.length
    headers = ["City Name", "Province Name", "Country Name"]
    rows_or_regions = regions || $ALL_REGIONS
    i = regions_count - 1
    row_len = headers.length - 1
    rows_or_regions.each do |row|
      incr = 0
      if regions
        row_len = row.split('-').length - 1
        i = row_len < i ? row_len : i
        return false if !include && row_len != i
      end
      (0..i).each do |e|
        region_key = permit_regions[(regions_count - 1) - e].gsub(/\s+/, "").upcase == (regions ? row.split('-')[row_len - e].gsub(/\s+/, "").upcase : row[headers[row_len-e]].gsub(/\s+/, "").upcase)
        region_key ? incr += 1 : break
      end
      return true if incr == (i + 1)
    end
    false
  end
end