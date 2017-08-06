class City < ApplicationRecord
  belongs_to :state
  has_many :included_cities
  has_many :excluded_cities
end
