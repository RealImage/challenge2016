# Ruby 2.3.0

require 'csv'

class ClassFactory
  def self.create_class(new_class, parent, inc_name, exc_name)
    c = Class.new do
    end

    klass = Kernel.const_set(new_class, c)
    klass.class_eval do
      attr_accessor :inc_name, :exc_name
      
      define_method(:initialize) do |*values|
        if parent
          p_class = Kernel.const_get(parent).new
          @inc_name = []
          @exc_name = []
          inc_name.each do |inc|
            inc = inc.split(',').flatten.join('-')
            @inc_name << inc if permission_check(inc, p_class)
          end
          @exc_name = p_class.exc_name + exc_name
        else
          @inc_name = inc_name
          @exc_name = exc_name
        end
      end
    end
  end
end

def permission_check(search, s_class)
  result = false
  search = search.split('-')
  s_class.inc_name.each do |inc|
    next if result
    inc = inc.split(',').flatten
    if inc.length > search.length
      result = false
    else
      result = inc[0].capitalize.eql?(search[0].capitalize)
      result = inc[1].capitalize.eql?(search[1].capitalize) unless inc[1].nil?
      result = inc[2].capitalize.eql?(search[2].capitalize) unless inc[2].nil?
    end
  end

  s_class.exc_name.each do |exc|
    next if !result
    exc = exc.split(',').flatten
    if exc.length > search.length
      result = true
    else
      result = !exc[0].capitalize.eql?(search[0].capitalize)
      result = !exc[1].capitalize.eql?(search[1].capitalize) unless exc[1].nil?
      result = !exc[2].capitalize.eql?(search[2].capitalize) unless exc[2].nil?
    end
  end
  return result
end

def read_csv(country, state, city)
  batch = CSV.read('/home/ajithror/Downloads/cities.csv', :headers=>true)
  batch = batch.select {|a| a["Country Name"].capitalize.eql?(country.capitalize)} unless country.nil?
  batch = batch.select {|a| a["Province Name"].capitalize.eql?(state.capitalize)} unless state.nil?
  batch = batch.select {|a| a["City Name"].capitalize.eql?(city.capitalize)} unless city.nil?
  return batch.count != 0
end

def get_country
  puts "Enter Country Name"
  country = gets.chomp.capitalize
  if country.eql?('')
    puts "Country name should not be blank."
    get_country
  else
    return country
  end
end

def fetch_data
  country = get_country
  puts "Enter Province Name"
  state = gets.chomp.capitalize
  state = state.eql?('') ? nil : state
  city =  if state
            puts "Enter City Name"
            gets.chomp.capitalize
          else
            nil
          end
  if read_csv(country, state, city.eql?('') ? nil : city)
    return [country, state, city.eql?('') ? nil : city].compact.join(',')
  else
    puts "Given details was not exist in our directory. Try agian."
    fetch_data
  end
end

def main_fun
  puts "Enter Total distributors"
  total = gets.chomp.to_i
  if total.eql?(0)
    puts "Please enter above 0"
    main_fun
  else
    (1..total).each do |a|
      puts "Enter distributor name"
      class_name = gets.chomp.capitalize
      if (a != 1)
        puts "Enter parent distributor name"
        parent = gets.chomp.capitalize
      else
        parent = nil
      end
      inc_name = []
      puts "No of includes"
      i = gets.chomp.to_i
      (1..i).each do |a|
        inc_name << fetch_data
      end
      exc_name = []
      puts "No of excludes"
      i = gets.chomp.to_i
      (1..i).each do |a|
        exc_name << fetch_data
      end
      ClassFactory.create_class(class_name, parent, inc_name, exc_name)
    end
  end
end

def input_check(d_name)
  puts "#{d_name} has permission to distribute in (country-province-city)"
  check = gets.chomp.split('-')
  country, state, city = check[0], check[1], check[2]
  if read_csv(country, state, city)
    return check.join('-')
  else
    puts "Given details was not exist in our directory. Try agian."
    input_check(d_name)
  end
end

main_fun()

puts "Enter no of times you want run the test"
n = gets.chomp.to_i
(1..n).each do |a|
  puts "Enter distributor name"
  d_name = gets.chomp.capitalize
  s_class = Kernel.const_get(d_name)
  c_in = s_class.new
  check = input_check(d_name)
  if permission_check(check, c_in)
    puts 'YES'
  else
    puts 'NO'
  end
end
