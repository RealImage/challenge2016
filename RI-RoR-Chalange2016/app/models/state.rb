class State < ApplicationRecord
  belongs_to :country
  has_many :cities
  has_many :included_states
  has_many :excluded_states
end
