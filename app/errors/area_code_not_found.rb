module Cities
  class AreaCodeNotFound < StandardError
    def initialize(code)
      super("ERROR: => Could not find area code: #{code}")
    end
  end
end
