Rails.application.routes.draw do

	root :to => "distributors#new"
  
  get 'distributors/check_city'
  get 'distributors/get_provinces'
  get 'distributors/get_cities'
  get 'distributors/get_countries'

  resources :distributors


  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
end
