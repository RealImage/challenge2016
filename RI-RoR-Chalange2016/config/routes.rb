Rails.application.routes.draw do

  resources :cities do
    get :search, on: :collection
  end
  resources :states do
    get :search, on: :collection
  end
  resources :countries
  resources :distbutors, only: [:index, :show, :new, :create, :destroy] do 
    get :sub_dist, on: :member
    get :permision, on: :member
  end
  devise_for :users
  root "distbutors#index"
end

