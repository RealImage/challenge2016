class Country < ApplicationRecord
  has_many :states
  has_many :included_countries
end
