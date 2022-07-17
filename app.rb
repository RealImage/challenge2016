require './distributors'
class App

def run
  loop do
    puts '1.enter or select the distributor by name'
    puts '2.exit'
    input = gets.chomp

  case input
    when '1'
      puts 'enter the distributor name'
      name = gets.chomp

      distributor = Distributor.find_by_name name

      if distributor.nil?
        distributor = Distributor.new(name: name)
        puts 'Distributor created successfully'
      else
        puts "selected Distributor #{distributor.name}"
      end

      loop do
       puts '1. Assign area to the distributor'
       puts '2. Assign Sub distributor'
       puts '3. listing the distributor Graph'
       puts '0. go back'

      distributor_options = gets.chomp
      case distributor_options
        when '1'
          distributor.assign_area
          puts distributor.inspect
        when '2'
          distributor.sub_distributor
        when '3'
          distributor.listing_distributor
        else
          break
      end
     end
      else
        exit 0
      end
    end
  end
end

App.new.run