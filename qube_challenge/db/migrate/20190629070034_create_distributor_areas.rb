class CreateDistributorAreas < ActiveRecord::Migration
  def change
    create_table :distributor_areas do |t|
      t.references :distributor, index: true, foreign_key: true
      t.string :country_code
      t.string :province_code
      t.string :city_code
      t.boolean :is_included

      t.timestamps null: false
    end
  end
end
