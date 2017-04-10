Rails.application.routes.draw do
  resources :distributors, except: [ 'show' ] do
    collection do
      get "/check", to: "distributors#check"
    end
  end
  root 'distributors#index'
end
