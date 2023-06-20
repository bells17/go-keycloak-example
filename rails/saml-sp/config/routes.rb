Rails.application.routes.draw do
  # For details on the DSL available within this file, see https://guides.rubyonrails.org/routing.html
  resources :saml, only: :index do
    collection do
      get :sso
      post :acs
      get :metadata
      get :logout
    end
  end

  root 'saml#index'
end
