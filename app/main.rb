require_relative 'distributor'

@distributors = []

def print_result(result = nil)
  puts "\n\n#{result}"
end

def print_menu
  puts "Quit".ljust(30, ' ') + "- 0\n" +
       "Create new distributor".ljust(30, ' ') + "- 1\n" +
       "Add inclusion".ljust(30, ' ') + "- 2\n" +
       "Add exclusion".ljust(30, ' ') + "- 3\n" +
       "Update sub-distributors".ljust(30, ' ') + "- 4\n" +
       "Show distributors".ljust(30, ' ') + "- 5\n" +
       "Load sample distributor data".ljust(30, ' ') + "- 6\n" +
       "Verify authorization".ljust(30, ' ') + "- 7"
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
    update_subdistributor
  when 5
    show_distributors
  when 6
    load_sample_distributor_data
  when 7
    check_authorization
  else
    print_result "Invalid action code"
  end
end

def check_authorization
  db = choose_distributor
  return unless db
  print_result  "Enter area code"
  area = gets.chomp
  print_result "Result: #{db.authorized_at?(area)}"
  false
end

def update_subdistributor
  print_result "Choose parent distributor"
  db1 = choose_distributor
  return unless db1
  print_result "Choose child distributor"
  db2 = choose_distributor
  return unless db2
  db2.extend_from db1
  print_result "Updated sub distributors"
  false
end

def show_distributors
  if @distributors.length > 0
    print_result "Distributor List\n\n"
    @distributors.each(&:print_details)
  else
    print_result "No saved distributors"
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
    print_result "No saved distributors"
  else
    print_result "Choose Distributor"
    @distributors.each.with_index(1) do |db, i|
      puts "#{db.name.ljust(30, ' ')}- #{i}"
    end
    db_code = gets.chomp.to_i
    if (db_code > 0 && db_code <= @distributors.length)
       @distributors[db_code - 1]
    else
      print_result "Invalid distributor code"
      nil
    end
  end
end

def create_distributor
  print_result "Enter distributor name"
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

  print_result "Sample distributor data loaded"
  false
end

quit = false
while quit != true
  print_menu
  quit = perform_action(gets.chomp.to_i)
  print_result
end

# Quick debug #

# db1 = Distributor.new('DB1')
# db1.include_area("INDIA")
# db1.include_area("UNITEDSTATES")
# db1.exclude_area("CHENNAI::TAMILNADU::INDIA")

# p db1.authorized_at?("CHICAGO::ILLINOIS::UNITEDSTATES") # -> YES
# p db1.authorized_at?("CHENNAI::TAMILNADU::INDIA") # -> NO
# p db1.authorized_at?("BANGALORE::KARNATAKA::INDIA") # -> YES

# db2 = Distributor.new('DB2')
# db2.extend_from(db1)

# db2.include_area("INDIA")
# db2.include_area("CHINA")
# db2.exclude_area("TAMILNADU::INDIA")

# p db2.inclusions # CHINA not included
# p db2.authorized_at?("CHINA") # -> NO
# p db2.authorized_at?("TAMILNADU::INDIA")# -> NO
# p db2.authorized_at?("KARNATAKA::INDIA") # -> YES

# db3 = Distributor.new('DB3')
# db3.extend_from(db2)

# db3.include_area("HUBLI::KARNATAKA::INDIA")

# p db3.authorized_at?("HUBLI::KARNATAKA::INDIA") # -> YES
# p db1.authorized_at?("KARNATAKA::INDIA") # -> YES
# p db2.authorized_at?("KARNATAKA::INDIA") # -> YES
# p db3.authorized_at?("KARNATAKA::INDIA") # -> NO

