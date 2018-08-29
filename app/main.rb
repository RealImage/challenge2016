require_relative 'distributor'

@distributors = []

def print_menu
  puts "  0 - Quit\n" +
       "  1 - Create new distributor\n" +
       "  2 - Include area for distributor\n" +
       "  3 - Exclude area for distributor\n" +
       "  4 - Show distributors\n" +
       "  5 - Load sample distributor data"
end

def perform_action(action)
  case action
  when 0
    exit
  when 1
    create_distributor
  when 2
    add_inclusion
  when 3
    add_exclusion
  when 4
    show_distributors
  when 5
    load_sample_distributor_data
  else
    puts "\n\nInvalid action code"
  end
end

def show_distributors
  if @distributors.length > 0
    puts "\n\nDistributor List\n\n"
    @distributors.each(&:print_details)
  else
    puts "\n\nNo saved distributors"
  end
end

def add_inclusion
  db = choose_distributor
  return unless db
  puts "Enter area code"
  code = gets.chomp
  db.include_area(code)
  false
end

def add_exclusion
  db = choose_distributor
  return unless db
  puts "Enter area code"
  code = gets.chomp
  db.exclude_area(code)
  false
end

def choose_distributor
  if @distributors.length == 0
    puts "\n\nNo distributors saved."
  else
    puts "\n\nChoose Distributor"
    @distributors.each.with_index(1) do |db, i|
      puts "#{i} - db.name"
    end
    puts @distributors.length
    db_code = gets.chomp.to_i
    if (db_code > 0 && db_code <= @distributors.length)
       @distributors[db_code - 1]
    else
      puts "\n\nInvalid distributor code."
      nil
    end
  end
end

def create_distributor
  puts "Enter distributor name"
  name = gets
  @distributors << Distributor.new(name)
  false
end

def load_sample_distributor_data
  db1 = Distributor.new('Sample Distributor 1')
  db1.include_area("INDIA")
  db1.include_area("UNITEDSTATES")
  db1.exclude_area("CHENNAI::TAMILNADU::INDIA")

  db2 = Distributor.new('Sample Distributor 2')
  db2.extend_from(db1)
  db2.include_area("INDIA")
  db2.include_area("CHINA")
  db2.exclude_area("TAMILNADU::INDIA")

  db3 = Distributor.new('Sample Distributor 3')
  db3.extend_from(db2)
  db3.include_area("HUBLI::KARNATAKA::INDIA")

  @distributors = []
  @distributors << db1 << db2 << db3

  puts "\n\nSample distributor data loaded"
  false
end

quit = false
while quit != true
  print_menu
  quit = perform_action(gets.chomp.to_i)
  puts "\n\n"
end

# p db1.authorized_at?("CHICAGO::ILLINOIS::UNITEDSTATES")
# p db1.authorized_at?("CHENNAI::TAMILNADU::INDIA")
# p db1.authorized_at?("BANGALORE::KARNATAKA::INDIA")

# db2 = Distributor.new('worldwide')
# db2.extend_from(db1)

# db2.include_area("INDIA")
# db2.include_area("CHINA")
# db2.exclude_area("TAMILNADU::INDIA")

# # p db2.inclusions
# # p db2.authorized_at?("CHINA")
# p db2.authorized_at?("TAMILNADU::INDIA")
# p db2.authorized_at?("KARNATAKA::INDIA")

# db3 = Distributor.new('local')
# db3.extend_from(db2)

# db3.include_area("HUBLI::KARNATAKA::INDIA")

# p db1.authorized_at?("KARNATAKA::INDIA")
# p db2.authorized_at?("KARNATAKA::INDIA")
# p db3.authorized_at?("KARNATAKA::INDIA")

