require 'csv'
namespace :seed do
  desc "Seeding data"
  task locations: :environment do
    first = true
    CSV.foreach("cities.csv") do |row|
      if first
        first = false
      else
        country = Country.find_or_create_by(name: row[5], code: row[2])
        state = State.find_or_create_by(name: row[4], code: row[1], country: country)
        City.find_or_create_by(name: row[3], code: row[0], state: state)
        puts row.join("-")
      end
    end

  end

end
