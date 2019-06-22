

$distributor_array = []
$current_parent = nil
class Distributer
 attr_accessor :parent, :name, :included, :excluded, :authorized_dist
 def initialize
  @name = ""
   @included = []
   @excluded = []
   @authorized_dist = []
   @parent = ""
 end
 
def find_from_dist_array(name)
    obj = nil
  $distributor_array.each do |dist|
    if (dist.name == name)
      obj = dist
    end
  end
  return obj
end
 
 def delete_from_dist_array(name)
  to_be_deleted = nil
  $distributor_array.each_with_index do |dist,i|
    if (dist.name == name)
      to_be_deleted = i
    end
  end
  if to_be_deleted
    $distributor_array[to_be_deleted] = nil
    $distributor_array.compact
  end
 end

 def uniq_dist_array()
  $distributor_array.each_with_index do |dist,i|
  end
 end

 def sub_area_check(area1,area2)
  checked = false;
  if (area1[:country] == area2[:country])
    if (area1[:state] == area2[:state] || area1[:state] == "all")
      if (area1[:city] == area2[:city] || area1[:city] == "all") 
          checked = true;
      end
    end
  end
  return checked;
 end
 
 def include_area
  p "Enter the area to be included in city,state,country or state,country or country format:"
  area = gets
   area = area.split(' ').join()
   area_arr = area.split(",")
   if (area_arr.size < 4 && area.to_s != "")
   if (area_arr.size == 3)
    incl = {:city => area_arr[0], :state => area_arr[1] , :country => area_arr[2]}
     elsif (area_arr.size == 2)
        incl = {:city => "all", :state => area_arr[0] , :country => area_arr[1]}
     elsif (area_arr.size == 1)
        incl = {:city => "all", :state => "all" , :country => area_arr[0]}
     end
   if (@parent != "")
      if ($current_parent.name == @parent)
        parent_obj = $current_parent
      else
      p "coming to else #{@parent}"
        parent_obj = find_from_dist_array(@parent)
        p "parent found #{parent_obj}"
      end
    sub_area = false
    excluded_area = false
    if parent_obj && parent_obj.included.any?
      parent_obj.included.each do |inc|
        unless sub_area
          sub_area = sub_area_check(inc,incl)
        end
      end
    end
    p "sub area #{sub_area}"
    if sub_area
       if parent_obj && parent_obj.excluded.any?
        parent_obj.excluded.each do |exc|
          unless excluded_area
            excluded_area = sub_area_check(exc,incl)
          end
        end
      end
      p excluded_area
      if excluded_area
        p "Not authorized to this address"
      else
        @included << incl
      end
    else
      p "Not authorized to this address"
    end
   else
    @included << incl
   end
  else
       p "Area is not inserted properly"
    end
  p @included
 end
 
 def exclude_area
    p "Enter the area to be excluded in city,state,country or state,country or country format:"
    area = gets
  area = area.split(' ').join()
    area_arr = area.split(",")
  if (area_arr.size < 4 && area.to_s != "")
      if (area_arr.size == 3)
        excl = {:city => area_arr[0], :state => area_arr[1] , :country => area_arr[2]}
      elsif (area_arr.size == 2)
        excl = {:city => "all", :state => area_arr[0] , :country => area_arr[1]}
      elsif (area_arr.size == 1)
       excl = {:city => "all", :state => "all" , :country => area_arr[0]}
      end
    sub_area = false
    @included.each do |inc|
    unless sub_area
      sub_area = sub_area_check(inc,excl)
    end
    end
    if sub_area
    @excluded << excl
    else
    p "The address is not available to  the distributer"
    end
  else
      p "Area is not inserted properly"
    end
  p @excluded
 end
 
 def add_name
 
 end
 
 def add_child
  p "Enter name of your child:"
   child_name = gets
   child = nil   
   $distributor_array.each do |dist|
    if (dist.name == child_name)
    p "Distributer already present. Add other details"
    child = dist
    end
   end
   unless child
  child = Distributer.new
  child.name = child_name.chomp
   end
   child.parent = self.name
   @authorized_dist  << child_name
   child_opt = 0
   while (child_opt.to_i < 3)
    p "Choose an option to continue: \n 1) Area allocation to authorized distributers \n 2) Authorize second level distributer \n 3) Exit"
    child_opt = gets
    delete_from_dist_array(child_name)
    $distributor_array << child
    if (child_opt.to_i == 1)
      child.include_area
      child.exclude_area
    elsif (child_opt.to_i == 2)
      child.add_child
    else
       child_opt = 4
        delete_from_dist_array(child_name)
        $distributor_array << child
      p "Distributer added!"
    end
  end
 end
 
end


p "Enter Name of the distributer"
name = gets
  dist_obj = nil
  $distributor_array.each do |dist|
    if (dist.name == name.chomp)
    p "Distributer already present. Add other details"
    dist_obj = dist
    end
    end
   
   unless dist_obj
  dist_obj = Distributer.new
  dist_obj.name = name.chomp
   end
   option = 0
while (option.to_i < 3) 
  p "Choose an option to continue: \n 1) Area allocation \n 2) Authorize other distributer \n 3) Exit"
  option = gets
  if (option.to_i == 1)
    p dist_obj.parent
    dist_obj.include_area
    dist_obj.exclude_area
  elsif (option.to_i == 2)
    $current_parent = dist_obj
  dist_obj.add_child
  else 
    dist_obj.delete_from_dist_array(name)
    dist_obj.uniq_dist_array()
    $distributor_array <<  dist_obj
  option = 4
    p "Successfully you exit"
  end
end
