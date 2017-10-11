require 'csv'
require 'sqlite3'
db = SQLite3::Database.new '../go/task'
# db.execute 'SELECT * FROM countries' do |row|
#   p row
# end
def get_id db,table,pair={}
  condition = "1=1"
  pair.each do |k,v|
    condition += " AND " unless condition == "1=1"
    condition += "#{k.to_s} = '#{v.to_s.gsub("'","`")}'" unless condition == "1=1"
    condition = "#{k.to_s} = '#{v.to_s.gsub("'","`")}'" if condition == "1=1"
  end
  qry = "select id from #{table} where #{condition}"
  rec = db.execute(qry)
  return rec.length > 0 ? rec[0][0] : nil
end


CSV.foreach("../cities.csv") do |row|
  country_id = get_id(db,"countries",{name: row[5], code: row[2]})
  if country_id.nil?
    db.execute "insert into countries(name,code) values ('#{row[5].gsub("'","`")}', '#{row[2]}')"
    country_id = get_id(db,"countries",{name: row[5], code: row[2]})
  end
  state_id = get_id(db,"states",{name: row[4], code: row[1], country_id: country_id})
  if state_id.nil?
    db.execute "insert into states(name,code,country_id) values ('#{row[4].gsub("'","`")}', '#{row[1]}',#{country_id})"
    state_id = get_id(db,"states",{name: row[4], code: row[1], country_id: country_id})
  end
  city_id = get_id(db,"cities",{name: row[3], code: row[0], state_id: state_id})
  if city_id.nil?
    db.execute "insert into cities(name,code,state_id) values ('#{row[3].gsub("'","`")}', '#{row[0]}',#{state_id})"
    city_id = get_id(db,"cities",{name: row[3], code: row[0], state_id: state_id})
  end
  puts "country_id: \t#{country_id}\t state_id: \t#{state_id}\t city_id: \t#{city_id}"
end
# City Code,Province Code,Country Code,City Name,Province Name,Country Name

