Rails.application.routes.draw do
  resources :samples
  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
  root 'homes#index'
  get 'get_parent' => 'homes#get_parent'
  get 'get_countries' => 'homes#get_countries'
  get 'get_states' => 'homes#get_states_for_country'
  get 'get_cities' => 'homes#get_cities_for_state'
  get 'distributor' => 'homes#new_distributor'
  get 'distributor/:parent_id' => 'homes#new_sub_distributor'
  # match '/distributor/:parent_id' => 'homes#new_distributor', :via => [:get], :as => 'with_parent'
  post 'distributor' => 'homes#create_distributor'
  post 'distributor/:parent_id' => 'homes#create_sub_distributor'

  get 'add_new_location/:id' => 'homes#add_new_location'
  post 'add_new_location/:id' => 'homes#create_new_location'

  #match '/distributor/(:parent_id)' => 'homes#create_distributor', :via => [:post], :as => 'create_distributor'
  get 'show/:id' => 'homes#show_distributors'
  get 'list' => 'homes#distributors_list'
  get 'check_authorisation/:id' => 'homes#check_authorisation'
  post 'check_authorisation/:id' => 'homes#check_authorisation_for_given_location'
  
end
