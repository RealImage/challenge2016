class Loader
  attr_accessor :path, :distributors, :store

  def initialize(path)
    @path = path
    @distributors = []
    populate_store
  end

  def load_data
    CSV.foreach(path, headers: false) do |row|
      perform(*row)
    end
  end

  def populate_store
    root_dir = __dir__.match(/(.+)\/app$/)[1]
    @store = Store.new("#{root_dir}/cities.csv")
    store.load_data
  end

  def perform(command = nil, *args)
    case command
    when 'CREATE'
      create_distributor(*args)
    when 'INCLUDE'
      add_inclusion(*args)
    when 'EXCLUDE'
      add_exclusion(*args)
    when 'EXTEND'
      extend_distributor(*args)
    when 'VERIFY'
      verify_permissions(*args)
    end
  end

  def verify_permissions(db_name, area)
    return unless validate_area(area)
    db = choose_distributor(db_name)
    return unless db
    puts "Authorization for #{db.name} at #{area}: #{db.authorized_at?(area)}"
  end

  def extend_distributor(child, parent)
    db1 = choose_distributor(parent)
    return unless db1
    db2 = choose_distributor(child)
    return unless db2
    db2.extend_from db1
  end

  def add_inclusion(name, code)
    return unless validate_area(code)
    db = choose_distributor(name)
    return unless db
    db.include_area(code)
  end

  def add_exclusion(name, code)
    return unless validate_area(code)
    db = choose_distributor(name)
    return unless db
    db.exclude_area(code)
  end

  def choose_distributor(name)
    @distributors.detect { |d| d.name == name }
  end

  def create_distributor(name)
    @distributors << Distributor.new(name)
  end

  def validate_area(code)
    store.find(code)
  end
end
