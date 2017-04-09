Rails.application.routes.draw do
  resources :distributors, except: [ 'show' ]
  root 'distributors#index'
end
