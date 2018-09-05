class Distributor
  attr_accessor :name, :inclusions, :exclusions, :master

  def initialize(name)
    @name = name
    @inclusions = {}
    @exclusions = {}
    @master = nil
  end

  def extend_from(other)
    @master = other
  end

  def include_area(code)
    return if master && master.is_not_authorized_at(code)
    inclusions[code] = code
  end

  def exclude_area(code)
    exclusions[code] = code
  end

  def authorized_at?(area)
    authorization = if master
                      (master.is_authorized_at(area) && is_authorized_at(area))
                    else
                      is_authorized_at(area)
                    end
    return_authorization(authorization)
  end

  def is_authorized_at(area)
    authorization_for(area)
  end

  def is_not_authorized_at(area)
    !is_authorized_at(area)
  end

  def print_details
    puts "Name: #{name}"
    puts "Inclusions: #{inclusions.keys.join(", ")}"
    puts "Exclusions: #{exclusions.keys.join(", ")}\n\n"
  end

  private

  def permitted_in(code)
    inclusions[code]
  end

  def excluded_from(code)
    exclusions[code]
  end

  def return_authorization(authorization)
    authorization ? 'YES' : 'NO'
  end

  def authorization_for(area)
    province, state, country = area.split("::")
    if country
      state_and_country = [state, country].join("::")
      ((permitted_in(country) || permitted_in(state_and_country) || permitted_in(area)) &&
      !(excluded_from(country) || excluded_from(state_and_country) || excluded_from(area)))
    elsif state
      ((permitted_in(state) || permitted_in(area)) && !(excluded_from(state) || excluded_from(area)))
    else
      permitted_in(area) && !excluded_from(area)
    end
  end
end
