class Distributor < ApplicationRecord	
  has_closure_tree

  has_many :distributor_allocations

  def self.check_allocation
  	puts "Enter Distributor name"    
    distributor = Distributor.includes(:distributor_allocations).find_by('lower(name) = ?', gets.chomp.downcase)
    if distributor.present?
      parent_distributors = distributor.ancestors.compact
      distributor_allocations = parent_distributors.map(&:distributor_allocations)
      Country.check_distributor(distributor_allocations)  	  
      Province.check_distributor(distributor_allocations)
      City.check_distributor(distributor_allocations)
    else
  	  puts "Invalid Distributor"
  	end  	
  end





end