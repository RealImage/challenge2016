class DistributorArea < ActiveRecord::Base
  belongs_to :distributor

  def self.is_accessible(distributor, input_areas)
    dis_record = Distributor.find_by_name(distributor)
    areas = input_areas.split('-')
    country = areas.last
    if areas.size == 3
      city = areas[0] 
      province = areas[1]
    end
    province = areas[0] if areas.size == 2
    country_code = Country.find_by_name(country).code if country 
    province_code = Province.find_by_name(province).code if province
    city_code = City.find_by_name(city).code if city
    anscestors_ids = dis_record.ancestors.pluck(:id)
    root_id = anscestors_ids.first || dis_record.id
    dis_ids = (anscestors_ids << dis_record.id)
    
    country_access = self.where(distributor_id: root_id, country_code: country_code, is_included: true).exists?
    return false unless country_access

    if province_code.present? && city_code.present?
      no_state_access = self.where(province_code: province_code, is_included: false, distributor_id: dis_ids).exists? 
      return false if no_state_access
      no_city_access = self.where(city_code: city_code, is_included: false, distributor_id: dis_record.id).exists?
      return false if no_city_access
    elsif province_code.present? && !city_code.present?
      no_state_access = self.where(province_code: province_code, is_included: false, distributor_id: dis_record.id).exists? 
      return false if no_state_access
    elsif !province_code.present? && city_code.present?
      no_city_access = self.where(city_code: city_code, is_included: false, distributor_id: dis_record.id).exists?
      return false if no_city_access
    elsif !province_code.present? && !city_code.present? && country_code.present?
      country_access = self.where(distributor_id: dis_record.id, country_code: country_code, province_code: nil, city_code: nil, is_included: true).exists?
      return false unless country_access
    end
    return true
  end
end
