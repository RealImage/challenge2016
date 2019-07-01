require 'csv'

class Distributor

  @@region_names = {}

  attr_accessor :name, :included_region, :excluded_region, :parent_distributor


  def initialize(name, included_region = [], excluded_region = [])
    @name = name
    @included_region = included_region
    @excluded_region = excluded_region
  end

  def check_distributor_permission(region)
    @have_permission = false
    region = region.split('-')
    verify_region('included', region)
    verify_region('excluded', region)
    @have_permission
  end

  def verify_region(type, region_given)
    region = type == 'included' ? self.included_region : self.excluded_region

    region.each do |region_present|
      if (region_present.length == 1 && region_given[-1] == region_present[-1]) || (self.parent_distributor.nil? && region_present.length == 2 && region_given[1] == region_present[0] &&
          region_given[-1] == region_present[-1]) || (self.parent_distributor.present? && region_present.length == 2 && region_given[1] == region_present[1] &&
          region_given[-1] == region_present[-1]) || (region_present.length == 3 && region_given[0] == region_present[0] && region_given[1] == region_present[1] &&
          region_given[-1] == region_present[-1])
        if type == 'included'
          @have_permission = true
        else
          @have_permission = false
        end
        break
      end
    end
    @have_permission
  end

  def self.valid_region?(region)
    return p "Please enter a valid region" if region.split('-').count > 3

    assigned_region = region.split('-')
    region_count = assigned_region.count
    is_valid = false
    @@region_names['all_regions'].each do |region|

      case region_count
      when 1
        if region[2] == assigned_region[0]
          is_valid = true
        end
      when 2
        if region[1] == assigned_region[0] && region[2] == assigned_region[1]
          is_valid = true
        end
      when 3
        if region[0] == assigned_region[0] && region[1] == assigned_region[1] && region[2] == assigned_region[2]
          is_valid = true
        end
      end
    end

    return is_valid
  end

  def self.fetch_regions(csv_file_path)
    regions = []
    begin
      CSV.foreach(csv_file_path, headers: true) do |row|
        row_hash = row.to_hash
        region_array = [row_hash['City Code'], row_hash['Province Code'], row_hash['Country Code']]
        regions << region_array.map {|remove_spaces| remove_spaces}
        @@region_names['all_regions'] = regions
      end
    rescue
      p 'Please enter the correct file path'
    end
    @@region_names
  end

  def self.check_name(name)
    ObjectSpace.each_object(self).to_a.detect {|distributor| distributor.name == name}
  end

  def self.distributor_count
    ObjectSpace.each_object(self).to_a.count
  end

  def self.distribution_ways
    p 'Following are the distribution ways'
    p '1) If you want to include country wise: Type the country name - Example: INDIA'
    p '2) If you want to include state wise: Type the state name and then country name, separated by - Example: KARNATAKA-INDIA'
    p '3) If you want to include city wise: Type the city name, state name, country name, separated by - Example: CHENNAI-TAMILNADU-INDIA'
    p 'You can assign one or more regions by typing the regions to next line'
  end

end

Distributor.fetch_regions('cities.csv')

loop do

  p 'Create a distributor with name'

  name_of_distributor = gets.chomp

  distributor = Distributor.new(name_of_distributor)

  p 'Do the distributor has any parent_distributor? Y/N'

  have_parent_distributor = gets.chomp


  if !(have_parent_distributor == 'Y' || have_parent_distributor == 'N')
    p 'Please type a valid option.'
    break
  else
    if have_parent_distributor == 'Y'
      p 'Type the name of the parent distributor'

      parent_distributor = Distributor.check_name(gets.chomp)

      unless parent_distributor
        p 'Please enter a valid parent name'
        break
      else
        distributor.parent_distributor = parent_distributor.name
      end
    end
  end

  loop do
    p 'Do you want to include or exclude region ? I/E'
    assign_type = gets.chomp
    if !(assign_type == 'I' || assign_type == 'E')
      p 'Please type a valid options'
      break
    else
      if Distributor.distributor_count == 0
        p 'Do you know how to assign regions? Y/N'
        knowledge_on_assigning = gets.chomp
        if !(knowledge_on_assigning == 'Y' || knowledge_on_assigning == 'N')
          p 'Please type a valid options'
          break
        end

        Distributor.distribution_ways if knowledge_on_assigning == 'N'
      end

      assign_type == 'I' ? (p 'Please type a region to include') : (p 'Please type a region to exclude')

      region = gets.chomp
      if Distributor.valid_region?(region)
        if distributor.parent_distributor.nil?
          assign_type == 'I' ? distributor.included_region << region.split('-') : distributor.excluded_region << region.split('-')
          assign_type == 'I' ? (p 'Region included successfully') : (p 'Region excluded successfully')
        else
          if assign_type == 'I' && parent_distributor.check_distributor_permission(region)
            distributor.included_region << region
          else
            p 'Please type a valid region.'
          end
          if assign_type == 'E'
            distributor.excluded_region << region
          end
        end
      else
        p 'Please type a valid region.'
        break
      end
      p 'Do you want to continue assigning regions? Y/N'
      continue_assigning = gets.chomp
      if !(continue_assigning == 'Y' || continue_assigning == 'N')
        p 'Please type a valid options'
        break
      end
      break unless continue_assigning == 'Y'
    end
  end

  p 'Do you want to create another distributor ? Y/N'

  continue_distribution = gets.chomp

  if !(continue_distribution == 'Y' || continue_distribution == 'N')
    p 'Please type a valid options'
    break
  elsif continue_distribution == 'N'
    break
  end
end

p 'Do you want to check permissions for the distributor ? Y/N'

permisssion_check = gets.chomp

if !(permisssion_check == 'Y' || permisssion_check == 'N')
  p 'Please type a valid options'
elsif permisssion_check == 'Y'
  p 'Please enter the name the distributor'

  distributor = Distributor.check_name(gets.chomp)
  p distributor

  unless distributor
    p 'Please enter a correct name'
  else
    p 'Please enter the region you want to check permission, in city-state-country format'
    region = gets.chomp

    unless Distributor.valid_region?(region)
      p 'Please enter a valid region'
    else
      if distributor.check_distributor_permission(region)
        p 'YES'
      else
        p 'NO'
      end
    end
  end
end