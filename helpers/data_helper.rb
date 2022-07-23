# frozen_string_literal: true

# Returns only the required country for a given distributor list
def update_included_list(hash_data)
  updated_hash = {}
  include_list.each do |included_region|
    next if included_region['countries'].nil?

    if included_region['province'].nil?
      unless updated_hash.key?(included_region['countries'].to_s)
        updated_hash.merge!(included_region['countries'] =>
          hash_data[included_region['countries']])
      end
    else
      updated_hash.merge!(included_region['countries'] => {}) unless updated_hash
                                                                     .key?(included_region['countries'])
      if included_region['cities'].nil?
        unless updated_hash[included_region['countries']].key?(included_region['province'])
          updated_hash[included_region['countries']].merge!(included_region['province'] =>
                hash_data[included_region['countries']][included_region['province']])
        end
      else
        unless updated_hash[included_region['countries']].key?(included_region['province'])
          updated_hash[included_region['countries']].merge!(included_region['province'] => {})
        end
        unless updated_hash[included_region['countries']][included_region['province']]
               .key?(included_region['cities'])
          updated_hash[included_region['countries']][included_region['province']]
            .merge!(included_region['cities'] => {})
        end
        updated_hash[included_region['countries']][included_region['province']]
          .merge!(included_region['cities'] =>
            hash_data[included_region['countries']][included_region['province']][included_region['cities']])
      end
    end
  end
  updated_hash
end
