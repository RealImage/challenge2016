# Using the hard coded data, as source of backend
reserved_permission={
  DISTRIBUTOR1: {
    include: 'INDIA, UNITEDSTATES',
    exclude: 'KARNATAKA-INDIA, CHENNAI-TAMILNADU-INDIA',
    super_class: nil
  },
  DISTRIBUTOR2: {
    include: 'INDIA',
    exclude: 'TAMILNADU-INDIA',
    super_class: 'DISTRIBUTOR1'
  }
}

def check_permission(distributor_detials, reserved_permission)
  # check for the DISTRIBUTOR to exist in records
  if reserved_permission.key?(distributor_detials[:distributor_name])
    name = distributor_detials[:distributor_name].to_sym
    super_distributor_name = reserved_permission[name][:super_class] ? reserved_permission[name][:super_class].to_sym : ''
    region_names = distributor_detials[:region_name].split(', ')

    region_names.each do |region|
      include_region = reserved_permission[name][:include].split(', ')
      exclude_region = reserved_permission[name][:exclude].split(', ')
      if exclude_region.include?(region)
        puts "\nNo, Distributor \"#{distributor_detials[:distributor_name]}\" don't have screening rights in #{region} region.\n"
      elsif include_region.include?(region)
        puts "\nYes!!!!, Distributor \"#{distributor_detials[:distributor_name]}\" have screening rights in #{region} region.\n"
        exclude_region.each do |exc_region|
          if exc_region.include?(region)
            puts "excluding #{exc_region}"
          end
        end
      elsif reserved_permission.key?(super_distributor_name)
        super_distributor = reserved_permission[super_distributor_name]
        super_include_region = super_distributor[:include].split(', ')
        unless super_include_region.include?(region)
          puts "\nNo, Distributor \"#{distributor_detials[:distributor_name]}\" don't have screening rights in #{region} region.\n"
        end
      else
        puts "\nNo, Distributor \"#{distributor_detials[:distributor_name]}\" don't have screening rights in #{region} region.\n"
      end
    end
  else
    puts "\nDistributor with the name \"#{distributor_detials[:distributor_name]}\", does not exists in records\n"
  end
end

def get_distributor_detials
  puts "\nEnter distributor name:\n"
  distributor_name = gets.chomp.to_sym
  puts "\nEnter the region to check permission. (eg: CITY-STATE-COUNTRY, CITY-STATE-COUNTRY)\n"
  region_name = gets.chomp
  { distributor_name: distributor_name, region_name: region_name }
end

distributor_detials = get_distributor_detials
check_permission(distributor_detials, reserved_permission)
