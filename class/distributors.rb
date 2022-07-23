# frozen_string_literal: true

require 'JSON'
require_relative '../helpers/helpers'

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
          permissible_data[region_to_check['countries']]['province'][region_to_check['province']].nil?
      display_message(region_to_check['province'], false)
    elsif region_to_check['cities'] &&
          permissible_data[region_to_check['countries']]['province'][region_to_check['province']]['cities'][region_to_check['cities']]
          .nil?
      display_message(region_to_check['cities'], false)
    else
      display_message(region_to_check['countries'], true)
    end
    nil
  end

  # Removes the countries that are not applicable for the distributor
  def remove_excluded_list
    updated_hash_data = JSON.parse(File.read('class/temp.json'))
    exclude_list.each do |excluded_region|
      next if excluded_region['countries'].nil?

      if excluded_region['province'].nil?
        updated_hash_data['countries'].delete(excluded_region['countries'])
        next
      elsif excluded_region['cities'].nil?
        updated_hash_data['countries'][excluded_region['countries']]['province'].delete(excluded_region['province'].to_s)
        next
      else
        updated_hash_data['countries'][excluded_region['countries']]['province'][excluded_region['province']]['cities']
          .delete(excluded_region['cities'].to_s)
      end
    end
    updated_hash_data
  end

  # Adds the countries/provinces that are available for the distributor
  def update_included_list(hash_data)
    updated_hash = {}
    include_list.each do |included_region|
      next if included_region['countries'].nil?

      if included_region['province'].nil?
        unless updated_hash.key?(included_region['countries'].to_s)
          updated_hash.merge!(included_region['countries'] =>
            hash_data['countries'][included_region['countries']])
        end
      else
        updated_hash.merge!(included_region['countries'] => {}) unless updated_hash
                                                                       .key?(included_region['countries'])
        if included_region['cities'].nil?
          unless updated_hash[included_region['countries']].key?('province')
            updated_hash[included_region['countries']].merge!('province' => {})
          end
          unless updated_hash[included_region['countries']]['province'].key?(included_region['province'])
            updated_hash[included_region['countries']]['province'].merge!(included_region['province'] =>
                  hash_data['countries'][included_region['countries']]['province'][included_region['province']])
          end
        else
          unless updated_hash[included_region['countries']].key?('province')
            updated_hash[included_region['countries']]['province'].merge!('province' => {})
          end
          unless updated_hash[included_region['countries']]['province'].key?(included_region['province'])
            updated_hash[included_region['countries']]['province'].merge!(included_region['province'] => {})
          end
          unless updated_hash[included_region['countries']]['province'][included_region['province']]
                 .key?('cities')
            updated_hash[included_region['countries']]['province'][included_region['province']]
              .merge!('cities' => {})
          end
          unless updated_hash[included_region['countries']]['province'][included_region['province']]['cities']
                 .key?(included_region['cities'])
            updated_hash[included_region['countries']]['province'][included_region['province']]['cities']
              .merge!(included_region['cities'] => {})
          end
          updated_hash[included_region['countries']]['province'][included_region['province']]['cities']
            .merge!(included_region['cities'] =>
              hash_data['countries'][included_region['countries']]['province'][included_region['province']]['cities'][included_region['cities']])
        end
      end
    end
    updated_hash
  end

  # Returns hash that are permitted for an distributor
  def permissible_data_hash
    removed_countries_hash = remove_excluded_list
    update_included_list(removed_countries_hash)
  end
end
