Rails.application.routes.draw do
  resources :samples
  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
  get '/' => 'homes#index'
  get 'distributor' => 'homes#new_distributor'
end
