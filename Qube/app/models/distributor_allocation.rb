class DistributorAllocation < ApplicationRecord  
  belongs_to :distributor
  belongs_to :country
  belongs_to :province
  belongs_to :city


  def self.allocate_distributor
    puts "Enter the Distributor name:"
    distributor = Distributor.find_by("lower(name) = ?",gets.chomp)
    if distributor.present?
      country = Country.can_allocate_distributor(distributor)
      if country
        province = Province.can_allocate_distributor(distributor,country)  
        if province
          city = City.can_allocate_distributor(distributor,province)
          puts "Do you want to include or exclude"
          status = gets.chomp  
          DistributorAllocation.find_or_create(distributor,country,province,city,status) 
        end
      end
    else
      puts "Invalid Distributor Name"
    end
  end

  def self.find_or_create(distributor,country,province,city,status)
    fields = {"distributor_id" => distributor.id,"country_id" => country.id,"province_id" => province.id,"city_id" => city.id,"status" => status}
    ds = DistributorAllocation.where(fields)
    if ds.present?
      puts "Already distributor exist"
    else
      da = DistributorAllocation.new(fields)
      if da.save!
        puts "Distributor Allocated successfully"
      else
        puts "Could not allocate distributor"
      end
    end
  end


end