class CreateProvinces < ActiveRecord::Migration
  def change
    create_table :provinces do |t|
      t.string :name
      t.string :code
      t.string :country_code

      t.timestamps null: false
    end
  end
end
