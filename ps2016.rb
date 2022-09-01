## Read both files
## LOCATION FORMAT
#  COUNTRY
    # PROVINCE
      # CITY
## PERMISSION FILE
## USERS and inheritance hierarchy L2R
## USER :
  # NAME
  # INHERIT FROM
  # INCLUDE
  # EXCLUDE
# -------------------------------------------------------------------------------------------------------------
  # SAMPLE FORMAT

  # Permissions: DISTRIBUTOR1
  # INCLUDE: INDIA
  # INCLUDE: UNITEDSTATES
  # EXCLUDE: KARNATAKA-INDIA
  # EXCLUDE: CHENNAI-TAMILNADU-INDIA

  # Permissions: DISTRIBUTOR2 < DISTRIBUTOR1
  # INCLUDE: INDIA
  # EXCLUDE: TAMILNADU-INDIA

# -------------------------------------------------------------------------------------------------------------
# we need three structures for this scenario
# @@CITY => STATE
# @@STATES => COUNTRY
# @@COUNTRY => STATE => CITY
# @@WORLD

  User = Struct.new(:name, :inherited,:included, :excluded)
  class FindAuthorization
    @@permission ={}
    def initialize(partner,permission)
      # User = Struct.new(:user, :inherit, :include, :permission)
      # "GET ALL LOCATIONS FIRST"
      @@city, @@states, @@world = get_csv(partner,true) if !partner.empty?
      ## get initial Users
      get_csv(permission,false) if !permission.empty?
    end

    def get_state(in_location)
      @@states.fetch(in_location)
    end
    def get_country(in_location)
      @@world.fetch(in_location)
    end
    def get_city(in_location)
      @@city.fetch(in_location)
    end

  ## check input is valid
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

  # Find  City/Province is within Province/Country
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

   def processing_logic(validuser, location)

      # binding.break
      excluded={}
      # "call checks and so forth
       @@excluded = 0
       @@included = 0
       parent={}
       to_check = valid_location(location)
       if @@permission.has_key?(validuser[:name])
          validuser = @@permission.fetch(validuser[:name])
       end

      # check authorization - Only Excluded might matter!
        if validuser[:inherited] != :inherited
          validuser[:inherited].each do |inh|
            user1 =  @@permission.fetch(inh)
            next  if user1.nil? or user1[:excluded] == :excluded
              user1[:excluded].each do |loc|
                if check_location(to_check,valid_location(loc),true)
                  @@excluded+=1
                  excluded.store(loc, loc)
                end
              end
            if user1[:included] != :included
              user1[:included].each do |loc|
                parent.store(loc, loc ) if check_location(to_check,valid_location(loc),true)
              end
          end
          end
        end
        # current user authorization
        if validuser[:included] != :included
          validuser[:included].each do |loc|
            @@included+=1 if check_location(to_check,valid_location(loc),false) and ( parent.has_key?(loc) or parent.empty? )
          end
        end
      if validuser[:excluded] != :excluded
          validuser[:excluded].each do |loc|
            @@excluded+=1 if check_location(to_check,valid_location(loc),true)
         end
        end
        if @@excluded > 0 or @@included == 0
          false
        else
          true
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
      city = {}
      state = {}
      world ={}
      CSV.foreach(File.open(filename), headers: header, :converters => :numeric, :header_converters => :symbol){ |row|
        next if row.empty?
        case
          when filename.include?('PS2016.txt')
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
            if !world.include?(row[:country_name])
              world.store(row[:country_name],row[:country_name])
            end
          when filename.include?('permission')
              if  row[0].include?('Permissions')
                  if user[:name] != :name
                    user[:included] = included if !included.empty?
                    user[:excluded] = excluded if !excluded.empty?
                    user[:inherited] = inherited if !inherited.empty?
                    user[:included].each do |city|
                      if !processing_logic(user, city)
                        puts "Parent Authorization missing#{city}"
                        next
                      else
                        @@permission.store(user[:name],user)
                      end
                    end

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
      if  filename.include?('PS2016.txt')
        return [ city, state, world ]
      else
          if user[:name] != :name
            user[:included] = included if !included.empty?
            user[:excluded] = excluded if !excluded.empty?
            user[:inherited] = inherited if !inherited.empty?
            # binding.break
            user[:included].each do |city|
              if !processing_logic(user, city)
                puts "Parent Authorization missin for #{city}"
                next
              else
                @@permission.store(user[:name],user)
              end
            end

          end
      end
  end
end

require 'debug'

# "start of processing"
ab = FindAuthorization.new("PS2016.txt",'permission.txt')
# "pS1" - minimum cost
run = true
puts ("--------------------------------------------------------------------------------------------------------------")
puts ("---------------------------------------- Distributor Authorization--------------------------------------------")
while run
  puts ("--------------------------------------------------------------------------------------------------------------")
  puts ("Enter x/X to EXIT")
  puts ("Enter Distributor Id: ")
  user = gets.chomp.upcase.gsub(/[[:space:]]/, '')
  puts ("Location in Format( INCLUDED/EXCLUDED : Ctry or Prov-Ctry or City-Prov-Ctry) : ")
  city = gets.chomp.upcase.gsub(/[[:space:]]/, '')
  if user == 'X' or city == 'X'
    run = false
  else
    begin
      obj_user = User.new(:name, :inherited, :included, :excluded)
      inherit_from = user.split('<')
      inherited=[]

      inherit_from.each_with_index do |value, index|
      case
        when index == 0
          obj_user[:name] = value
        when  index != 0
          inherited << value
        end
      end
      obj_user[:inherited] = inherited if !inherited.nil?
      ab.processing_logic(obj_user, city) ? 'Authorized' : 'Not Authorized'
    rescue => exception
      puts exception.message
    else
      next
    end
  end
end
