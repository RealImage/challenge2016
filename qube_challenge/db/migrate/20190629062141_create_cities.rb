class CreateCities < ActiveRecord::Migration
  def change
    create_table :cities do |t|
      t.string :name
      t.string :code
      t.string :province_code

      t.timestamps null: false
    end
  end
end
