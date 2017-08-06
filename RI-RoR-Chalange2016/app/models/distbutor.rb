class Distbutor < ApplicationRecord
  has_many :included_countries
  has_many :included_states
  has_many :included_cities
  has_many :excluded_states
  has_many :excluded_cities
end
