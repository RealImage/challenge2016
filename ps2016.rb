## Read both files
## LOCATION FORMAT
#  COUNTRY
    # PROVINCE
      # CITY
## DEFINE User Format
## USER :
  # INHERIT FROM
  # INCLUDE
  # EXCLUDE
class FindAuthorization

    User = Struct.new(:name, :inherited,:included, :excluded)
    def initialize(partner,permission)
      # User = Struct.new(:user, :inherit, :include, :permission)
    @@permission =  get_csv(permission,false) if !permission.empty?
    @@city, @@states, @@world = get_csv(partner,true) if !partner.empty?
    end

    # @@CITY => STATE
    # @@STATES => COUNTRY
    # @@COUNTRY => STATE => CITY
    # @@WORLD
def get_state(in_location)
  @@states.fetch(in_location)
end
def get_country(in_location)
  @@world.fetch(in_location)
end
def get_city(in_location)
  @@city.fetch(in_location)
end

  def valid_location(in_location)
    location = in_location.split("-")
    case location.size
      when 1
        raise "Invalid Country" if !@@world.has_key?(location[0])
        when 2
        raise "Invalid State/Country" if !@@states.has_key?(location[0]) or !@@world.has_key?(location[1])
      when 3
        raise "Invalid City/State/Country" if  !@@city.has_key?(location[0]) or !@@states.has_key?(location[1]) or !@@world.has_key?(location[2])
    end
    in_location
  end

  # Find
  def find_inclusive_or_not(src,dest)
    max_size = src.size
    case max_size
      when 3
        # 0 - city 1 - province 2 - country
        in_state = get_city(src[0])
        in_country = get_state(src[1])
        raise 'City-Province-Country Do not match!' if in_country != src[2] or in_state != src[1]
        if dest.size == 2
          in_state == dest[0] && in_country == dest[1]
        elsif dest.size == 1
          in_country == dest[0]
        end
      when 2
        in_country = get_state(src[0])
        raise 'Province-Country Do not match!' if in_country != src[1]
        in_country == dest[0]
    end

  end

  def check_location(from, target,excluded=false)
    # input vs Target sizes dont match"
    #  TN-INDIA != MAS-TN-INDIA , TN-INDIA => MAS-TN-INDIA , INDA != MAS-TN-INDIA  &  INDIA != MAS-TN-INDIA
   src = from.split("-")
   dest = target.split("-")
   case
   when src.size < dest.size
    return false
    when src.size == dest.size
        src == dest
    when src.size > dest.size
      #  return site included in the site or not
      find_inclusive_or_not(src, dest)
    end
  end

    def processing_logic(user, location)
      binding.break
       match_found = false
       @@excluded = 0
       @@included = 0
        raise "Invalid User" if !@@permission.has_key?(user)
        validuser = @@permission.fetch(user)
       # st = @@country["INDIA"].filter{ |item| p item if !item["TAMIL NADU"].nil? }
        to_check = valid_location(location)
      # check authorization
        if validuser[:inherited] != :inherited
          validuser[:inherited].each do |inh|
            user1 =  @@permission.fetch(inh)
            next  if user1.nil?
            # user1[:included].each do |loc|
            #   loc = valid_location(loc)
            #   @@included+=1 if  check_location(to_check,loc,false)
            # end
            user1[:excluded].each do |loc|
              loc = valid_location(loc)
              @@excluded+=1 if check_location(to_check,loc,true)
            end
          end
        end
        validuser[:included].each do |loc|
            loc = valid_location(loc)
            @@included+=1 if check_location(to_check,loc,false)
          end
          validuser[:excluded].each do |loc|
            loc = valid_location(loc)
            @@excluded+=1 if check_location(to_check,loc,true)
          end

        if @@excluded > 0 or @@included == 0
          raise "No Authorization"
        else
          raise "Authorized! Please proceed"
        end
      end

   private
  # Input file read and prepare struct format
  ##
  def get_csv(filename, header)
      require 'csv'
      user = User.new(:name, :inherited,:included, :excluded)
      included=[]
            excluded=[]
                 inherited=[]
                 permission={}
      provinces = {}
      city = {}
       state = {}
          country = {}
      world ={}
      CSV.foreach(File.open(filename), headers: header, :converters => :numeric, :header_converters => :symbol){ |row|
        next if row.empty?
        case
          when filename.include?('test_ps')
            city_name=[]
            states=[]
            row[:city_name] = row[:city_name].upcase.gsub(/[[:space:]]/, '')
            row[:province_name] = row[:province_name].upcase.gsub(/[[:space:]]/, '')
            row[:country_name] = row[:country_name].upcase.gsub(/[[:space:]]/, '')
            if !city.include?(row[:city_name])
              city.store(row[:city_name],row[:province_name])
            end

            if !state.include?(row[:province_name])
              state.store(row[:province_name], row[:country_name])
            end

            if !provinces.include?(row[:province_name])
                city_name.append(row[:city_name])
                provinces.store(row[:province_name], city_name)
            else
              city_name = provinces[row[:province_name]]
              city_name.append(row[:city_name])
              provinces[row[:province_name]] = city_name
            end

            province = provinces.select{ |p| p == row[:province_name]}

            if !country.include?(row[:country_name])
              world.store(row[:country_name],row[:country_name])
              states.append(province)
              country.store(row[:country_name],states)
            else
              states = country[row[:country_name]]
              states.append(province)
              country[row[:country_name]]  = states.uniq
            end
          when filename.include?('permission')
              if  row[0].include?('Permissions')
                  if user[:name] != :name
                    user[:included] = included if !included.empty?
                    user[:excluded] = excluded if !excluded.empty?
                    user[:inherited] = inherited if !inherited.empty?
                    permission.store(user[:name],user)
                    user = User.new(:name, :inherited,:included, :excluded)
                    included=[]
                    excluded=[]
                    inherited=[]
                  end
                  inherit_from = row[0].split(':')[1].split('<')
                  # inherit_from = user[1].split('<')
                  user[:name] = inherit_from[0].strip.upcase
                  if inherit_from.size > 1
                    inherit_from.each_with_index {|v,i| inherited.append(v.strip.upcase) if i!=0}
                  end
              elsif row[0].include?('INCLUDE')
                included.append(row[0].split(":")[1].strip.upcase)
              elsif row[0].include?('EXCLUDE')
                excluded.append(row[0].split(":")[1].strip.upcase)
              end
        end
      }
      if  filename.include?('test_ps')
        return [ city, state, world ]
      else
          if user[:name] != :name
            user[:included] = included if !included.empty?
            user[:excluded] = excluded if !excluded.empty?
            user[:inherited] = inherited if !inherited.empty?
              permission.store(user[:name],user)
          end
        permission
      end
  end
end

require 'debug'

# "start of processing"
ab = FindAuthorization.new("test_ps.txt",'permission.txt')
# "pS1" - minimum cost
run = true
puts ("--------------------------------------------------------------------------------------------------------------")
puts ("---------------------------------------- Distributor Authorization--------------------------------------------")
puts ("--------------------------------------------------------------------------------------------------------------")
while run
  puts ("Enter x/X to EXIT")
  puts ("Enter Distributor Id: ")
  user = gets.chomp.upcase.gsub(/[[:space:]]/, '')
  puts ("Location in Format( Ctry or Prov-Ctry or City-Prov-Ctry) : ")
  city = gets.chomp.upcase.gsub(/[[:space:]]/, '')
  if user == 'x' or city == 'X'
    run = false
  else
    begin
        ab.processing_logic(user, city)
    rescue => exception
      puts exception.message
    else
      next
    end
  end
end





