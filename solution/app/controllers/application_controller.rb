class ApplicationController < ActionController::Base
  protect_from_forgery with: :exception
  # def hello
  #   dir = Rails.root.join('public', 'uploads')
  #   Dir.mkdir(dir) unless Dir.exist?(dir)
  #   File.open(dir.join('san.txt'), 'wb') do |file|
  #     file.write('uploaded_io.read')
  #   end
  #
  #   require 'csv'
  #   @csv = CSV.read(Rails.root.join('public', 'database', 'cities.csv'), :headers=>true)
  # end
end
