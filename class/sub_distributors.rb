# frozen_string_literal: true

class Subdistributors < Distributors
  attr_accessor :permissible_data, :main_distributor_permissible_list

  def initialize(name, include_list, exclude_list , main_distributor_permissible_list)
    super(name,include_list,exclude_list, '')
    sub_distributors.flatten!
    @main_distributor_permissible_list = main_distributor_permissible_list
    @permissible_data = permissible_data
    binding.pry
    puts 'hi'
  end

  #Returns permissible data for an given sub-distributor
  def permissible_data
    updated_hash_data = @main_distributor_permissible_list
    exclude_list.each do |excluded_region|
        if excluded_region['countries'].nil?
          next
        elsif excluded_region['province'].nil?
          updated_hash_data.delete(excluded_region['countries'])
        elsif excluded_region['cities'].nil?
          updated_hash_data[excluded_region['countries']].delete(excluded_region['province'])
        else
          updated_hash_data[excluded_region['countries']][excluded_region['province']]
            .delete(excluded_region['cities'])
        end
      end
    updated_hash_data
  end

  # Checks if the distributor object has permission or not.
  def check_distributor_data(check_list)
    region_to_check = check_list.first
    if permissible_data[region_to_check['countries']].nil?
      display_message(region_to_check['countries'], false)
    elsif region_to_check['province'] && permissible_data[region_to_check['countries']][region_to_check['province']].nil?
      display_message(region_to_check['province'], false)
    elsif region_to_check['cities'] && permissible_data[region_to_check['countries']][region_to_check['province']][region_to_check['cities']].nil?
      display_message(region_to_check['cities'], false)
    else
      display_message(region_to_check['countries'], true)
    end
    nil
  end
end
