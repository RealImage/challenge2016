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
    code = area.split("::")
    if code.length == 3
      ((permitted_in(code[2]) || permitted_in(code[1..2].join("::")) || permitted_in(area)) && !(excluded_from(code[2]) || excluded_from(code[1..2].join("::")) || excluded_from(area)))
    elsif code.length == 2
      ((permitted_in(code[1]) || permitted_in(area)) && !(excluded_from(code[1]) || excluded_from(area)))
    else
      permitted_in(area) && !excluded_from(area)
    end
  end
end
