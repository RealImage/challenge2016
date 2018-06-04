require 'csv'
require 'json'
class Distributor
  attr_accessor :id, :name, :included_permission, :excluded_permission, :parent_distrubutor

  def initialize(params)
  	self.id = Random.rand(10000)
    self.name = params[:name]
    self.included_permission = params[:included_permission]
    self.excluded_permission = params[:excluded_permission]
  end

  def set_distrubtion_permission(params)
  	self.included_permission = params[:included_permission]
    self.excluded_permission = params[:excluded_permission]
  end

  def autorize_ditrubtion(distrubutor, params)
    owner_can_distribute = self.allowed_permission
    owner_can_not_distribute = self.no_permission
    autorize_can_didtribute = permission_array(params[:included_permission])
    autorize_can_not_didtribute = permission_array(params[:excluded_permission])
    permission_status = Distributor.check_permission(owner_can_distribute, owner_can_not_distribute, autorize_can_didtribute, autorize_can_not_didtribute)
    if permission_status
      put "not autorize distrubutor"
    else
      put "Can distrubute"
    end
  end

  def self.check_permission(owner_can_distribute, owner_can_not_distribute, autorize_can_didtribute, autorize_can_not_didtribute)
    permission = false
    owner_permited_country = owner_can_distribute[0]
    owner_permited_state = owner_can_distribute[1]
    owner_permited_city = owner_can_distribute[2]
    owner_restricted_country = owner_can_not_distribute[0]
    owner_restricted_state = owner_can_not_distribute[1]
    owner_restricted_city = owner_can_not_distribute[2]
    owner_permited_region = @json_data.select {|x| owner_permited_country.include? x[:country_name] }.select {|x| owner_permited_state.include? x[:province_name] }.
    select {|x| owner_permited_city.include? x[:city_name] }
    owner_restricted_region = @json_data.select {|x| owner_restricted_country.include? x[:country_name] }.select {|x| owner_restricted_state.include? x[:province_name] }.
    select {|x| owner_restricted_city.include? x[:city_name] }

    owner_restricted_region.each do |region|
      if region[:country_name] == autorize_can_didtribute[0]
        if region[:province_name] == autorize_can_didtribute[1]
          if region[:city_name] == autorize_can_didtribute[2]
            permission = true
          end
        end
      elsif region[:province_name] == autorize_can_didtribute[1]
        if region[:city_name] == autorize_can_didtribute[2]
          permission = true
        end
      elsif region[:city_name] == autorize_can_didtribute[2]
        permission = true
      else
        permission = false
      end
    end   
  end

  def allowed_permission
    allowed_permission = []
    permission_list = self.included_permission.split(',')
    permission_list.each do |permission|
      allowed_permission << permission.split('-').uniq
    end
    allowed_permission
  end

  def no_permission
    no_permission = []
    permission_list = self.excluded_permission.split(',')
    permission_list.each do |permission|
      no_permission << permission.split('-').uniq
    end
    no_permission
  end

  def permission_array(data)
    permission = []
    data_list = data.split(',')
    data_list.each do |data|
      permission << data.spli('-').uniq
    end
    permission
  end

  def create_db
    csv_data = File.open('cities.csv').read
    data = CSV.new(csv_data, :headers => true, :header_converters => :symbol, :converters => :all)
    @json_data = data.to_a.map {|row| row.to_hash }
  end
end