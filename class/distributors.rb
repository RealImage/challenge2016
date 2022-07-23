# frozen_string_literal: true

require 'json'
require_relative '../helpers/input_helper'
require_relative '../helpers/data_helper'
require_relative '../helpers/permission_helper'

# Distributor class to hold the distributor data
class Distributors
  attr_accessor :name, :sub_distributors, :include_list, :exclude_list, :permissible_data

  def initialize(name, include_list, exclude_list, sub_distributors = '')
    @name = name
    @include_list = include_list
    @include_list.delete_if(&:empty?)
    @exclude_list = exclude_list
    @exclude_list.delete_if(&:empty?)
    @sub_distributors = [] << sub_distributors
    self.sub_distributors.compact!
    self.sub_distributors.flatten!
    @permissible_data = permissible_data_hash
  end

  # Checks if the distributor object has permission or not.
  def check_distributor_data(check_list)
    region_to_check = check_list.first
    if permissible_data[region_to_check['countries']].nil?
      display_message(region_to_check['countries'], false)
    elsif region_to_check['province'] &&
          permissible_data[region_to_check['countries']][region_to_check['province']].nil?
      display_message(region_to_check['province'], false)
    elsif region_to_check['cities'] &&
          permissible_data[region_to_check['countries']][region_to_check['province']][region_to_check['cities']]
          .nil?
      display_message(region_to_check['cities'], false)
    else
      display_message('', true)
    end
    nil
  end

  # Removes the countries that are not applicable for the distributor
  def remove_excluded_list
    updated_hash_data = JSON.parse(File.read('class/cities.json'))
    exclude_list.each do |excluded_region|
      next if excluded_region['countries'].nil?

      if excluded_region['province'].nil?
        updated_hash_data.delete(excluded_region['countries'])
        next
      elsif excluded_region['cities'].nil?
        updated_hash_data[excluded_region['countries']].delete(excluded_region['province'].to_s)
        next
      else
        updated_hash_data[excluded_region['countries']][excluded_region['province']]
          .delete(excluded_region['cities'].to_s)
      end
    end
    updated_hash_data
  end

  # Returns hash that are permitted for an distributor
  def permissible_data_hash
    removed_countries_hash = remove_excluded_list
    update_included_list(removed_countries_hash)
  end
end
