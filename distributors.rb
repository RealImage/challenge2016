require 'set'
require 'csv'

class Distributor
  @@distributors = {}
  @@current_sequence = 0
  @@cities = []

  @@rows = CSV.foreach('cities.csv') do |row|
    @@cities.push([row[3].downcase.strip, row[4].downcase.strip, row[5].downcase.strip].join(','))
  end

  attr_accessor :parent_id, :children_ids, :exclude_path, :include_path, :id, :name

  def initialize args
    sequence = self.class.next_sequence

    @name =args[:name]
    @parent_id = args[:parent_id]
    @children_ids = (args[:children_ids].nil? ? [] : args[:children_ids])
    @exclude_path =args[:exclude] || []
    @include_path =args[:include] || []
    @id = sequence

    add_distributor
  end

  def self.next_sequence
    @@current_sequence +=1
  end

  def self.find_by_name name
    @@distributors[name.strip]
  end

  def assign_area main_distributor=nil
     loop do
       puts 'enter include area with comma seperate values'
       area_include = gets.chomp
       if main_distributor && !main_distributor.include_path.include?(area_include)
         puts 'Not include for main distributor'
       elsif !(valid_area? area_include)
         puts 'Not a valid area'
       else
         self.exclude_path.reject{|a| a.eql?(area_include.downcase.strip)}
         self.include_path.push area_include
       end

       puts 'enter Exclude area with comma seperate values'
       area_exclude = gets.chomp

       if !(valid_area? area_exclude)
         puts 'Not a valid area'
       else
         self.include_path.reject{|a| a.eql?(area_exclude.downcase.strip)}
         self.exclude_path.push area_exclude
       end

       puts 'press Y to continue,N to Skip'
       selection =gets.chomp
       if !selection.eql? 'Y'
         break
       end
     end
  end

  def valid_area? area
    area = area.downcase.split(',').map(&:strip).join(',')
    selected_cities = @@cities.select{|area_places| (area_places.eql?(area))}
    !selected_cities.empty?
  end

  def sub_distributor
    puts 'enter the name of the sub_distributor'
    sub_distributor_name = gets.chomp
    sub_distributor = self.class.find_by_name sub_distributor_name

    if sub_distributor.nil?
      puts 'would you want to create a sub distributor,Y or N'
      sub_distributor_selection = gets.chomp
        if sub_distributor_selection.eql? "Y"
         sub_distributor= Distributor.new(name: sub_distributor_name, parent_id:self.id)
         children_ids.push(sub_distributor.id)
          sub_distributor.assign_area self
        end
    elsif !sub_distributor.parent_id.nil? || !sub_distributor.parent_id.eql?(self.id)
      puts 'sub distributor already Exists and linked with other distributor'
    else
      sub_distributor.assign_area self
    end
  end

  def listing_distributor
    distributor= self

    puts 'Current distributor details'
    distributor.print_distributor_info

    if !distributor.nil?
      parent_id = distributor.parent_id

      while !parent_id.nil?
        puts 'Parent distributors'
        parent = Distributor.find_by_id parent_id
        parent.print_distributor_info
        parent_id = parent.parent_id
      end

      children_ids = distributor.children_ids
      while !children_ids.empty?
        puts 'Sub distributors'
        childrens = distributor.find_by_childrens children_ids
        children_ids = []

        childrens.each do |children|
          children.print_distributor_info
          children_ids = children_ids + children.children_ids
        end
      end
    else
     puts  "Distributed not exists"
    end
  end

  def find_by_childrens children_ids
    @@distributors.select {|name, distributor| children_ids.include? distributor.id}.values
  end

  def self.find_by_id id
    distributor = nil

    @@distributors.each do |key, value|
      if (value.id.eql? id)
        distributor = value
      end
    end

    distributor
  end

  def print_distributor_info
    puts "Include #{self.include_path}"
    puts "Exclude #{self.exclude_path}"
    puts "Name #{self.name}"
  end

  def add_distributor
    @@distributors[self.name] = self
    puts 'Added distributor'
    puts @@distributors.inspect
  end

  def self.distributors
    @@distributors
  end
end